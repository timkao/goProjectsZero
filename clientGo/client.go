package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Dialing error", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? type Q to exit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			fmt.Println("Exit the application!")
			return
		}
		conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
