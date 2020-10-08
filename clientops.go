package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

func joinChannel(ktx, pubKey, signedKey, ktxCertFileName string, keyCollection *ED25519Keys) *websocket.Conn {

	// request a websocket connection
	conn := requestSocket(ktx, "1")

	// using that connection, attempt to join the channel
	joinedChannel := handShake(conn, pubKey)

	// parse channel messages
	socketMsgParser(ktx, pubKey, signedKey, joinedChannel, keyCollection)

	// return the connection
	return conn
}

func handShake(conn *websocket.Conn, pubKey string) *websocket.Conn {

	// new users should send JOIN with the pubkey
	if isFirstTime {
		joinReq := "JOIN " + pubKey
		_ = conn.WriteMessage(1, []byte(joinReq))
	}

	// returning users should send RTRN and the signed CA cert
	if !isFirstTime {
		certString := readFile(selfCertFilePath)
		rtrnReq := "RTRN " + pubKey + " " + certString
		_ = conn.WriteMessage(1, []byte(rtrnReq))
	}

	return conn
}

func socketMsgParser(ktx, pubKey, signedKey string, conn *websocket.Conn, keyCollection *ED25519Keys) {

	_, joinResponse, err := conn.ReadMessage()
	handle("There was a problem reading the socket: ", err)

	if strings.HasPrefix(string(joinResponse), "WCBK") {
		isTrusted = true
		isFirstTime = false
		fmt.Printf(brightgreen + " ✔️\nConnected!\n" + white)
		fmt.Printf("\nType `"+brightpurple+"send %s filename.json"+white+"` where filename.json is a file in %s to send a JSON object in a transaction.\n\n", ktx, configTxDir)
	}

	if strings.Contains(string(joinResponse), string(capkMsg)) {
		convertjoinResponseString := string(joinResponse)
		trimNewLinejoinResponse := strings.TrimRight(convertjoinResponseString, "\n")
		trimCmdPrefix := strings.TrimPrefix(trimNewLinejoinResponse, "CAPK ")
		ncasMsgtring := signKey(keyCollection, trimCmdPrefix[:64])
		composedNcasMsgtring := string(ncasMsg) + " " + ncasMsgtring
		_ = conn.WriteMessage(1, []byte(composedNcasMsgtring))
		_, certResponse, err := conn.ReadMessage()
		isFirstTime = false
		convertStringcertResponse := string(certResponse) // keys := generateKeys()
		trimNewLinecertResponse := strings.TrimRight(convertStringcertResponse, "\n")
		trimCmdPrefixcertResponse := strings.TrimPrefix(trimNewLinecertResponse, "CERT ")
		handle("There was an error receiving the certificate: ", err)
		ktxCertFileName := configHostsDir + "/" + ktx + ".cert"
		createFile(ktxCertFileName)
		writeFile(ktxCertFileName, trimCmdPrefixcertResponse[:192])
		fmt.Printf(brightgreen + "\nCert Name: ")
		fmt.Printf(white+"%s", ktxCertFileName)
		fmt.Printf(brightgreen + "\nCert Body: ")
		fmt.Printf(white+"%s\n", trimCmdPrefixcertResponse[:192])
	}
}

func returnMessage(conn *websocket.Conn, pubKey string) *websocket.Conn {
	if !isFirstTime {
		certString := readFile(signedKeyFilePath)
		rtrnReq := "RTRN " + pubKey[:64] + " " + certString
		_ = conn.WriteMessage(1, []byte(rtrnReq))
	}
	return conn
}

func requestSocket(ktx, protocolVersion string) *websocket.Conn {
	urlConnection := url.URL{Scheme: "ws", Host: ktx, Path: "/api/v" + protocolVersion + "/channel"}
	conn, _, err := websocket.DefaultDialer.Dial(urlConnection.String(), nil)
	handle(brightred+"There was a problem connecting to the channel: "+brightcyan, err)
	return conn
}

// Send Takes a data string and a websocket connection
func sendV1Transaction(filename, publicKey, address, certfile string, conn *websocket.Conn) {

	fmt.Printf(brightcyan+"\nSending %s to %s using cert %s"+nc, brightmagenta+filename+brightcyan, brightmagenta+address+brightcyan, brightgreen+certfile+brightcyan)
	fmt.Println("")

	// Check if the filename we have been given is good
	isTXReady := fileExists(configTxDir + "/" + filename)
	if !isTXReady {
		fmt.Printf("%s does not exist", configTxDir+"/"+filename)
		return
	}

	// Read the proposed TX as bytes
	txCandidate := readFileBytes(configTxDir + "/" + filename)
	if len(txCandidate) < 2 {
		fmt.Printf(brightred+"Error: %s\nTX is empty or non existant.\n"+nc, red+configTxDir+"/"+filename+brightred)
		return
	}
	// Read the contents of the cert as string
	channelCertContents := readFile(certfile)

	// Encode the proposed TX as a string of bytes
	fileBytesAsEncodedString := hex.EncodeToString(txCandidate)

	// Prepare the transaction SEND string
	tx := string(sendMsg) + " " + publicKey + " " + channelCertContents + " " + fileBytesAsEncodedString

	// Send the transaction over the socket
	fmt.Printf(brightcyan+"\nSending transaction: \n%s...", brightpurple+tx[:32])
	err := conn.WriteMessage(1, []byte(tx))
	handle("There was a problem sending your transaction ", err)

	for {
		// Since we may hear a message that isnt ours, we should listen
		// for a confirmation containing our public key.
		okbytes := []byte("OK " + publicKey)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf(brightred+"\n[%s] [%s] Transaction confirmation problem \n"+nc, timeStamp(), conn.RemoteAddr())
			return
		}

		// If we hear the OK message, move the transaction we just sent to the
		// sent transaction archive.
		if bytes.Equal(msg, okbytes) {
			timeNow := unixTimeStampNano()
			fmt.Printf(brightgreen + "✔️ \nTransaction success!\n" + nc)
			createFile(configTxArchiveDir + timeNow + ".tx")
			writeFile(configTxArchiveDir+timeNow+".tx", tx)
			deleteFile(configTxDir + "/" + filename)
			fmt.Printf("%s moved to %s", filename, configTxArchiveDir+timeNow+".tx\n")
			return
		}

		// If it hasnt arrived in 3 seconds, its not coming.
		delay(3)
		return
	}
}
