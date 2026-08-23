package main

import (
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
		fmt.Println("Error listening: ", err)
		return
	}
    
    defer listener.Close()
    
    fmt.Println("Server started on :8080")
    
    for {
        conn, err := listener.Accept()
        if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    
	response := "OK\n"
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Server write error: ", err)
		return
	}
}