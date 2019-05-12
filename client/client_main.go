package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"time"
)

const LOGGER_MODE = "GUI" // GUI OR CLI
const DEFAULT_URL = "127.0.0.1"
const DEFAULT_PORT = "9797"

var conn net.Conn

func main() {
	startGUI()
}

func startClient(dest string) {
	logger(fmt.Sprintf("Connecting to %s...\n", dest))

	var dialErr error
	conn, dialErr = net.Dial("tcp", dest)

	if dialErr != nil {
		if _, t := dialErr.(*net.OpError); t {
			logger(fmt.Sprint("Some problem connecting."))
		} else {
			logger(fmt.Sprint("Unknown error: " + dialErr.Error()))
		}
		os.Exit(1)
	}

	go readConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		logger(fmt.Sprint("> "))
		text, _ := reader.ReadString('\n')

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))

		_, err := conn.Write([]byte(text))
		fmt.Println("Writing to server")
		if err != nil {
			logger(fmt.Sprint("Error writing to stream."))
			os.Exit(2)
		}
	}
}

func writeToServer(text string) {
	fmt.Fprintf(conn, "%s\n", text)
}

func stopClient() {
	logger(fmt.Sprint("Bout to close the server!"))
	conn.Close()
}

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			command := handleCommands(text)
			if !command {
				logger(fmt.Sprintf("\b\b** %s\n> ", text))
			}

			if !ok {
				logger(fmt.Sprint("Reached EOF on server connection."))
				break
			}
		}
	}
}

func handleCommands(text string) bool {
	r, err := regexp.Compile("^%.*%$")
	if err == nil {
		if r.MatchString(text) {

			switch {
			case text == "%quit%":
				fmt.Println("\b\bServer is leaving. Hanging up.")
				os.Exit(0)
			}

			return true
		}
	}

	return false
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
