package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	
	conn, err := l.accept()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err = io.EOF {
			break
		}
		fmt.Println("error reading from client: ", err.Error())
        os.Exit(1)


		fmt.Println(string(buf[:n]))
		conn.Write([]byte("+OK\r\n"))
	}
}