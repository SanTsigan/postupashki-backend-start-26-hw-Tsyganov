package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()
    
	message, _ := bufio.NewReader(conn).ReadString('\n')
	
    if message == "OK\n" {
        fmt.Println("Correct answer from server")
        os.Exit(0)
    } else {
        fmt.Println("Wrong answer from server")
        os.Exit(1)
    }
}