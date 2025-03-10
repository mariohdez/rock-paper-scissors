package main

import (
	"bufio"
	"github.com/mariohdez/rockpaperscissors/internal/input"
	"github.com/mariohdez/rockpaperscissors/internal/net/protocol"
	"log"
	"net"
	"os"
)

func main() {
	inputReader := input.NewTextReader(bufio.NewScanner(os.Stdin), os.Stdout)
	log.Print("Attempting to connect to the match server")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Unable to connect to the match server", err)
		return
	}

	log.Print("Successfully connected to the match server")

	for {
		cmd, _, err := protocol.ReceiveMessage(conn)
		if err != nil {
			log.Println("Unable to read command", err)
			conn.Close()
			return
		}

		switch cmd {
		case protocol.WhatIsYourName:
			name, err := inputReader.ReadName()
			if err != nil {
				log.Println("Unable to read name", err)
				conn.Close()
				return
			}

			err = protocol.SendMessage(conn, protocol.MyNameIs, name)
			if err != nil {
				log.Println("Unable to send command", err)
				conn.Close()
				return
			}
		}
	}
}
