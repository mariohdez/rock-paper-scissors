package main

import (
	"errors"
	"github.com/mariohdez/rockpaperscissors/internal/net/protocol"
	"log"
	"net"
	"sync"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	log.Println("Game server listening on", listen.Addr())

	connChannel := make(chan net.Conn, 2)
	go func() {
		for i := 0; i < 2; {
			conn, err := listen.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					log.Println("Listener closed, shutting down server.")
					return
				}

				log.Println("received error accepting new connection:", err)
				continue
			}

			connChannel <- conn
			i++
		}
	}()

	player1Conn := <-connChannel
	player2Conn := <-connChannel
	listen.Close()

	log.Println("I have received two connections.")
	handleGame(player1Conn, player2Conn)
	log.Println("now will kick off game session")

}

func handleGame(player1Conn, player2Conn net.Conn) {
	var wg sync.WaitGroup
	connToName := make(map[net.Conn]string)
	connToNameLock := sync.Mutex{}
	askAndGetName := func(wg *sync.WaitGroup, playerConn net.Conn, connToName map[net.Conn]string, connToNameLock *sync.Mutex) {
		defer wg.Done()
		err := protocol.SendMessage(playerConn, protocol.WhatIsYourName, "")
		if err != nil {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			return
		}

		cmd, data, err := protocol.ReceiveMessage(playerConn)
		if err != nil {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			return
		}
		if cmd != protocol.MyNameIs {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			return
		}

		connToNameLock.Lock()
		connToName[playerConn] = data
		connToNameLock.Unlock()
	}

	wg.Add(2)
	go askAndGetName(&wg, player1Conn, connToName, &connToNameLock)
	go askAndGetName(&wg, player2Conn, connToName, &connToNameLock)

	wg.Wait()

	log.Println("Got both players names")
}
