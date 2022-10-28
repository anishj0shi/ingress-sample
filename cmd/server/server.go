package main

import (
	"net"
)

func main() {
	listner, err := net.ListenTCP("tcp", &net.TCPAddr{
		Port: 2379,
	})
	if err != nil {
		panic(err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}

		println("Receieved Request")

		conn.Write([]byte("I heard you on TCP !!!"))

		conn.Close()

	}

}
