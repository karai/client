package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// inputHandler This is a basic input loop that listens for
// a few words that correspond to functions in the app. When
// a command isn't understood, it displays the help menu and
// returns to listening to input.
func inputHandler(keyCollection *ED25519Keys) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n\n%v%v%v\n", white+"Type '", brightgreen+"menu", white+"' to view a list of commands")
	for {
		fmt.Print(brightcyan + "$> " + nc)

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.Compare("help", text) == 0 {
			menu()
		} else if strings.Compare("?", text) == 0 {
			menu()
		} else if strings.Compare("pubkey", text) == 0 {
			fmt.Printf(brightcyan + "Pubkey: ")
			fmt.Printf(cyan+"%s\n", keyCollection.publicKey)
		} else if strings.Compare("menu", text) == 0 {
			menu()
		} else if strings.Compare("version", text) == 0 {
			menuVersion()
		} else if strings.Compare("license", text) == 0 {
			printLicense()
		} else if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		} else if strings.HasPrefix(text, "connect") {
			menuConnectChannel(text, keyCollection)
		} else if strings.HasPrefix(text, "send") {
			sendBody := strings.TrimPrefix(text, "send ")
			cmdString := strings.Split(sendBody, " ")
			// conn := requestSocket(cmdString[0], "1")
			handShake(conn, keyCollection.publicKey)
			certFile := configHostsDir + "/" + cmdString[0] + ".cert"
			sendV1Transaction(cmdString[1], keyCollection.publicKey, cmdString[0], certFile, conn)
		}
	}
}

func menuConnectChannel(text string, keyCollection *ED25519Keys) {
	address := strings.TrimPrefix(text, "connect ")

	if !strings.Contains(address, ":") {
		fmt.Printf("\nDid you forget to include the port?\n")
	}

	if strings.Contains(address, ":") {
		var domain = strings.Split(address, ":")
		var ktxCert = configHostsDir + "/" + domain[0] + ".cert"

		if !fileExists(ktxCert) {
			fmt.Printf(brightcyan+"\nConnecting:"+white+" %s", address)
			conn = joinChannel(address, keyCollection.publicKey, keyCollection.signedKey, "", keyCollection)
		} else if fileExists(ktxCert) {
			isFirstTime = false
			conn = joinChannel(address, keyCollection.publicKey, keyCollection.signedKey, ktxCert, keyCollection)
		}
	}
}
func menu() {
	menuOptions := []string{"GENERAL_OPTIONS"}
	menuData := map[string][][]string{
		"GENERAL_OPTIONS": {
			{
				white + "For more help ->" + brightcyan + " https://github.com/karai" + brightwhite,
				"",
				"connect <ip.ip.ip.ip:port> \t\t Connects to Karai Coordinator",
				"send <ip.ip.ip.ip:port> <file.json> \t Send tx with contents of <file.json> to channel",
				"",
				"version \t\t\t\t Displays version",
				"license \t\t\t\t Displays license",
				"exit \t\t\t\t\t Quit immediately",
				"",
			},
		},
	}

	for _, opt := range menuOptions {
		fmt.Printf(brightgreen + "\n" + opt)
		for menuOptionColor, options := range menuData[opt] {
			switch menuOptionColor {
			case 0:
				fmt.Printf(brightwhite)
			case 1:
				fmt.Printf(brightblack)
			}
			for _, message := range options {
				fmt.Printf("\n" + message)
			}
		}
	}

	fmt.Printf("\n")
}
