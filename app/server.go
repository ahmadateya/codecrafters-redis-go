package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l := getListener()
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// defer conn.Close()

		buf := make([]byte, 1024)
		if _, err := conn.Read(buf); err != nil {
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("reading from client...")

		// responding to the client
		fmt.Println("responding to the client...")
		conn.Write([]byte("+PONG\r\n"))
	}
}

func getListener() net.Listener {
	// Listen on TCP port 6379 at any interface (0.0.0.0)
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	fmt.Println("Listening on port 6379...")
	return l
}
