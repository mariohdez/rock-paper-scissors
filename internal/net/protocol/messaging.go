package protocol

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

const maxMessageSize = 1024

func SendMessage(conn net.Conn, command, data string) error {
	if !isCommandValid(command) {
		return errors.New("invalid RPS command")
	}

	message := command + ":" + data + "\n"
	b := []byte(message)
	if len(b) > maxMessageSize {
		return errors.New(fmt.Sprintf("message too large: %d bytes, with max set to%d bytes", len(b), maxMessageSize))
	}

	_, err := conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func isCommandValid(command string) bool {
	if command != WhatIsYourName && command != MyNameIs {
		return false
	}
	return true
}

func ReceiveMessage(conn net.Conn) (string, string, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	elements := strings.Split(line, ":")
	if len(elements) != 2 {
		return "", "", errors.New(fmt.Sprintf("invalid RPS message: %s", line))
	}

	return elements[0], elements[1], nil
}
