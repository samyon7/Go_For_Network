package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1201"
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		// Run as goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	// Close connection exit
	defer conn.Close()

	var buf [512]byte

	for {
		// Read up
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		fmt.Println(string(buf[0:]))
		// Menulis beberapa byte (sebanyak n)
		_, err2 := conn.Write(buf[0:n])

		if err2 := nil {
			return
		}
	}
}


func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s",err.Error())
		os.Exit(1)
	}
}


