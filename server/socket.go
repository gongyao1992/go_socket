package server

import (
	"fmt"
	"os"
	"net"
)

func checkError(err error)  {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvConnMsg(conn net.Conn)  {
	buf := make([]byte, 50)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn cloesed")
			return
		}

		fmt.Println("recv msg: ", string(buf[0:n]))
	}
}

func SocketRun()  {
	listen_sock, err := net.Listen("tcp", "127.0.0.1:10000")

	checkError(err)

	defer listen_sock.Close()

	for {
		new_conn, err := listen_sock.Accept()

		if err != nil {
			continue
		}

		go recvConnMsg(new_conn)
	}
}