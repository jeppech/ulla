package main

import (
	"fmt"
	"net"
	"os"
	"log"
	"strings"
)

var numConn = 1

func CheckError (err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp4","127.0.0.1:13337")
	CheckError(err)

	conn, err := net.ListenUDP("udp", addr)
	CheckError(err)

	defer conn.Close()

	fmt.Println("Listening for Ulla packets on ", addr)

	for {
		buffer := make([]byte, 1024);

		n, remote_addr, err := conn.ReadFromUDP(buffer)

		if isRemoteWhitelisted(remote_addr.String()) {
			fmt.Println(remote_addr, "is whitelisted")
		} else {
			continue
		}

		fmt.Println("Got:", string(buffer[:n]))

		if err != nil {
			log.Fatal(err)
		}

		sendAck(conn, remote_addr)
	}
}

func isRemoteWhitelisted(remote_addr string) bool {
	whitelist := []string{
		"127.0.0.1",
	}

	substr := strings.Split(remote_addr, ":")

	for _, addr := range whitelist {
		if addr == substr[0] {
			return true
		}
	}

	return false
}

func sendAck(conn *net.UDPConn, remote_addr *net.UDPAddr) {
	numConn++
	message := []byte(fmt.Sprintf("[%d]ACK", numConn))
	conn.WriteToUDP(message, remote_addr)
}
