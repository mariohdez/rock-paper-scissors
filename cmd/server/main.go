package main

import (
	"errors"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/net/protocol"
	"log"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Game server listening on", listener.Addr())
	connChannel := make(chan net.Conn, 2)

	go listenToConnections(listener, connChannel)
	go kickOffGameSession(connChannel)

}

func listenToConnections(listener net.Listener, connChannel chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Println("Listener closed, shutting down server.")
				return
			}

			log.Println("received error accepting new connection:", err)
			continue
		}

		connChannel <- conn
	}
}

func kickOffGameSession(connChannel <-chan net.Conn) {
	for {
		log.Println("Waiting for two players...")
		player1Conn := <-connChannel
		player2Conn := <-connChannel
		log.Println("players received starting new match")

		go orchestrateNewGame(player1Conn, player2Conn)
	}
}

func orchestrateNewGame(player1Conn, player2Conn net.Conn) {
	defer player1Conn.Close()
	defer player2Conn.Close()

	getNames(player1Conn, player2Conn)

	game := game.NewMatch(3, nil, nil, nil, nil, nil)
	game.Start()
}

func getNames(player1Conn, player2Conn net.Conn) {
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
