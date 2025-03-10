package main

import (
	"errors"
	"fmt"
	"github.com/mariohdez/rockpaperscissors/internal/game"
	"github.com/mariohdez/rockpaperscissors/internal/net/protocol"
	"github.com/mariohdez/rockpaperscissors/internal/user"
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

	connToName, err := getNames(player1Conn, player2Conn)
	if err != nil {
		log.Println("Could not find player1's name:", player1Conn)
		return
	}
	player1Name, _ := connToName[player1Conn]
	player2Name, _ := connToName[player2Conn]

	_ = game.NewMatch(3, &user.Player{
		Name: player1Name,
	}, &user.Player{
		Name: player2Name,
	}, nil, nil, nil)
}

func getNames(player1Conn, player2Conn net.Conn) (map[net.Conn]string, error) {
	var wg sync.WaitGroup
	connToName := make(map[net.Conn]string)
	connToNameLock := sync.Mutex{}
	errCh := make(chan error, 2)
	askAndGetName := func(wg *sync.WaitGroup, playerConn net.Conn, connToName map[net.Conn]string, connToNameLock *sync.Mutex) {
		defer wg.Done()
		err := protocol.SendMessage(playerConn, protocol.WhatIsYourName, "")
		if err != nil {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			errCh <- err
			return
		}

		cmd, data, err := protocol.ReceiveMessage(playerConn)
		if err != nil {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			errCh <- err
			return
		}
		if cmd != protocol.MyNameIs {
			log.Println("error receiving message from client:", err)
			playerConn.Close()
			errCh <- fmt.Errorf("received message from client, but it's not a name")
			return
		}

		log.Printf("[%s] Got name: %s", playerConn.RemoteAddr(), data)
		connToNameLock.Lock()
		connToName[playerConn] = data
		connToNameLock.Unlock()
	}

	wg.Add(2)
	go askAndGetName(&wg, player1Conn, connToName, &connToNameLock)
	go askAndGetName(&wg, player2Conn, connToName, &connToNameLock)

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return connToName, nil
}
