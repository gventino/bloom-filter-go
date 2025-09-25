package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"main/bloomfilter"
	"net"
	"os"
)

const SocketPath = "/tmp/bloomfilter.sock"

var (
	cmdAdd    = []byte("ADD")
	cmdExists = []byte("EXISTS")
)

func main() {
	log.Println("Initializing Bloom Filter...")
	bf := bloomfilter.NewBloomFilterWithFNVa(100000, 7)

	if err := os.RemoveAll(SocketPath); err != nil {
		log.Fatalf("Failed to remove existing socket file: %v", err)
	}

	listener, err := net.Listen("unix", SocketPath)
	if err != nil {
		log.Fatalf("Failed to listen on socket: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s. Ready to accept connections.", SocketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleConnection(conn, bf)
	}
}

func handleConnection(conn net.Conn, bf *bloomfilter.BloomFilter) {
	defer conn.Close()
	log.Printf("Client connected: %s", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from client: %v", err)
			}
			break
		}

		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := bytes.Fields(line)
		if len(parts) < 2 {
			conn.Write([]byte("ERROR: Invalid command. Use ADD <data> or EXISTS <data>\n"))
			continue
		}

		command := bytes.ToUpper(parts[0])
		data := parts[1]

		processCommand(conn, bf, command, data)
	}
	log.Printf("Client disconnected: %s", conn.RemoteAddr().String())
}

func processCommand(conn net.Conn, bf *bloomfilter.BloomFilter, command []byte, data []byte) {
	if bytes.Equal(command, cmdAdd) {
		bf.Add(data)
		conn.Write([]byte("OK\n"))
		log.Printf("Added item: %s", data)
	} else if bytes.Equal(command, cmdExists) {
		if bf.Exists(data) {
			conn.Write([]byte("true\n"))
		} else {
			conn.Write([]byte("false\n"))
		}
		log.Printf("Checked item: %s", data)
	} else {
		conn.Write([]byte("ERROR: Unknown command.\n"))
	}
}
