package main

import "github.com/gorilla/websocket"

// Attribution constants
const (
	appName        = "Karai Client"
	appDev         = "The TurtleCoin Developers"
	appDescription = appName + " a simple client for connecting to karai."
	appLicense     = "https://choosealicense.com/licenses/mit/"
	appRepository  = "https://github.com/karai/client"
	appURL         = "https://karai.io"
)

// File & folder constants
const (
	configDir          = "./config"
	configKeyDir       = configDir + "/keys"
	configHostsDir     = configDir + "/hosts"
	configTxDir        = configDir + "/transactions"
	configTxArchiveDir = configTxDir + "/archived/"
	pubKeyFilePath     = configKeyDir + "/" + "pub.key"
	privKeyFilePath    = configKeyDir + "/" + "priv.key"
	signedKeyFilePath  = configKeyDir + "/" + "signed.key"
	selfCertFilePath   = configKeyDir + "/" + "self.cert"
)

// Client Values
var (
	isFirstTime = true
	isTrusted   = false
)

// Coordinator values
var (
	joinMsg  []byte = []byte("JOIN")
	ncasMsg  []byte = []byte("NCAS")
	capkMsg  []byte = []byte("CAPK")
	certMsg  []byte = []byte("CERT")
	peerMsg  []byte = []byte("PEER")
	pubkMsg  []byte = []byte("PUBK")
	nsigMsg  []byte = []byte("NSIG")
	sendMsg  []byte = []byte("SEND")
	conn     *websocket.Conn
	upgrader = websocket.Upgrader{
		EnableCompression: true,
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
	}
)
