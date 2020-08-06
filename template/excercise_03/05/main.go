package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		defer conn.Close()
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
	}
}
