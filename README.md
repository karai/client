![karaiclient](https://user-images.githubusercontent.com/34389545/95399749-78ecc480-08ce-11eb-885e-eae789dc2ecd.png)

[![Discord](https://img.shields.io/discord/388915017187328002?label=Join%20Discord)](http://chat.turtlecoin.lol) [![GitHub issues](https://img.shields.io/github/issues/karai/client?label=Issues)](https://github.com/karai/client/issues) ![GitHub stars](https://img.shields.io/github/stars/karai/client?label=Github%20Stars) ![Build](https://github.com/karai/client/workflows/Build/badge.svg) ![GitHub](https://img.shields.io/github/license/karai/client) ![GitHub issues by-label](https://img.shields.io/github/issues/karai/client/Todo) [![Go Report Card](https://goreportcard.com/badge/github.com/karai/client)](https://goreportcard.com/report/github.com/karai/client)

**Website:** [📝 karai.io](https://karai.io) **Browse:** [💻 Karai Pointer Explorer](https://karai.io/explore/) **Read:** [🔗 Official Karai Blog](https://karai.io/dev/)

![karai-demo](https://user-images.githubusercontent.com/34389545/95405569-cc1a4380-08dd-11eb-97ef-ca9c8368e52e.gif)

## Usage

> Note: Karai aims to always compile and run on **Linux** targetting the **AMD64** CPU architecture. Other operating systems and architectures may compile and run this software but should do so expecting some inconsistencies.

**Launch Karai Client**

```bash
go build

./client
```

**Connect to a Karai Channel**

```bash
./client
# Launch the client

connect 167.71.104.172:4200
# connect <ip.ip.ip.ip:port>
```

**Send A Transaction**

```bash
./client
# Launch the client

connect 167.71.104.172:4200
# connect <ip.ip.ip.ip:port>
# Connect to a channel.

send 167.71.104.172:4200 new.json
# send <ip.ip.ip.ip:port> <file.json>
# ./config/transaction/file.json
```

**A note on transactions**

> Note: A transaction can be a JSON object or arbitrary data. If the message contents do not pass JSON validation, they get hex encoded and sent as a string of bytes.

> Type `menu` to view a list of functions. Functions that are darkened are disabled.

## Dependencies

-   Golang 1.14+ https://golang.org

## Operating System

Karai supports Linux on AMD64 architecture, but may compile in other settings. Differences between Linux and non-Linux installs should be expected.

**Optional:** Compile with all errors displayed, then run binary. Avoids "too many errors" from hiding error info.

`go build -gcflags="-e" && ./client`

## Contributing

This repo only receives stable version release updates, development happens in a private repo. Please make an issue before writing code for a PR.

-   MIT License
-   `gofmt`
-   go modules
-   stdlib > \*

## Thanks to:

[![turtlecoin](https://user-images.githubusercontent.com/34389545/80266529-fb0b6880-8661-11ea-9a75-4cb066834775.png)](https://turtlecoin.lol)
[![IPFS](https://user-images.githubusercontent.com/34389545/80266356-0c07aa00-8661-11ea-8308-84639318213a.png)](https://ipfs.io)
[![LibP2P](https://user-images.githubusercontent.com/34389545/80266502-e4651180-8661-11ea-8367-54bf59e26470.png)](https://libp2p.io)
[![GOLANG](https://user-images.githubusercontent.com/34389545/80266422-6b65ba00-8661-11ea-836a-d1904ec15b94.png)](https://golang.org)
