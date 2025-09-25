package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const SocketPath = "/tmp/bloomfilter.sock"

func main() {
	conn, err := net.Dial("unix", SocketPath)
	if err != nil {
		log.Fatalf("Failed to connect to socket: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to Bloom Filter server.")

	runAutomatedTests(conn)

	runInteractiveMode(conn)
}

func runAutomatedTests(conn net.Conn) {
	fmt.Println("--- Running Automated Tests ---")
	sendCommand(conn, "ADD golang")
	sendCommand(conn, "ADD bloomfilter")
	sendCommand(conn, "EXISTS golang")
	sendCommand(conn, "EXISTS bloomfilter")
	sendCommand(conn, "EXISTS non-existent-item")
	sendCommand(conn, "INVALID COMMAND")
	fmt.Println("--- Automated Tests Finished ---")
	fmt.Println("\nEnter commands in the format: ADD <data> or EXISTS <data>. Type 'exit' to quit.")
}

func runInteractiveMode(conn net.Conn) {
	userInputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := userInputReader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading user input: %v", err)
			break
		}

		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting.")
			break
		}

		if input == "" {
			continue
		}

		sendCommand(conn, input)
	}
}

func sendCommand(conn net.Conn, command string) {
	_, err := conn.Write([]byte(command + "\n"))
	if err != nil {
		log.Printf("Failed to send command '%s': %v", command, err)
		return
	}

	serverReader := bufio.NewReader(conn)
	response, err := serverReader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read server response for command '%s': %v", command, err)
		return
	}

	fmt.Printf("CMD: [%s] -> RSP: [%s]\n", command, strings.TrimSpace(response))
}
