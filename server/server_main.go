package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const LOGGER_MODE = "GUI" // GUI OR CLI
const DEFAULT_URL = "127.0.0.1"
const DEFAULT_PORT = "9797"

var ln net.Listener

func main() {

	// addr := flag.String("addr", "127.0.0.1", "Enter the custom address")
	// port := flag.Int("port", 8000, "Enter the custom port")
	// flag.Parse()

	startGUI()
}

func startServer(fullAddr string) {
	// fullAddr := *addr + ":" + strconv.Itoa(*port)
	var lnErr error

	ln, lnErr = net.Listen("tcp", fullAddr)
	if lnErr != nil {
		logger(fmt.Sprintf("%s\n", lnErr))
		return
	}
	logger(fmt.Sprintf("Server started at %s\nWaiting for connections...\n", fullAddr))

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func stopServer() {
	logger(fmt.Sprintf("Stopping the server on %s\n", ln.Addr()))
	ln.Close()
	cleanLogger()
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	logger(fmt.Sprintf("Got a connection from %s!\n", remoteAddr))

	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()

		if !ok {
			break
		}
		handleMessage(scanner.Text(), conn)
	}

	logger(fmt.Sprintf("%s disconnected!\n", remoteAddr))
}

func handleMessage(msg string, conn net.Conn) {
	logger(fmt.Sprintf("> %s\n", msg))
	conn.Write([]byte(fmt.Sprintf("> %s\n", msg)))

	if len(msg) > 0 && msg[0] == '/' {
		switch msg {
		case "/time":
			timeStamp := time.Now()
			responseMsg := fmt.Sprintf("< %s\n", timeStamp)
			conn.Write([]byte(responseMsg))

		default:
			conn.Write([]byte("Unknown Command!\n"))
		}
	}
}

func logger(msg string) {
	switch LOGGER_MODE {
	case "GUI":
		appendToLogger(msg)
	case "CLI":
		fmt.Println(msg)
	default:
		fmt.Println("No logging interface specified!")
	}
}
