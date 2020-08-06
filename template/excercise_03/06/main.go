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
			if ln == "" {
				fmt.Println("This is end of request headers !!")
				break
			}
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
		conn.Close()
	}
}
