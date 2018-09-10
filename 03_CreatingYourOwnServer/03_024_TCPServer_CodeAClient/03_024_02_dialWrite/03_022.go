package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

/*
 console: go run 03_022.go
 browser: localhost:8080/somethingElse
 */
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	//we never get here
	//we have an open stream connection
	//how does the above reader know when it's done?
	fmt.Println("code got here")
}
