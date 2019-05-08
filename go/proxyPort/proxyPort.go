package main

import (
	"log"
	"net"
	"os"
)

var proxyAddr = [2]string{"127.0.0.1:1309", "127.0.0.1:1310"}

func proxyRequest(f net.Conn, t net.Conn) {
	defer f.Close()
	defer t.Close()
	var buffer = make([]byte, 40960)
	for {
		n, ef := f.Read(buffer)
		if ef != nil {
			break
		}
		_, et := t.Write(buffer[:n])
		if et != nil {
			break
		}
	}
	log.Println("Proxy Done")
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: proxyPort [from_addr] [to_addr]")
	}

	proxyAddr[0] = os.Args[1]
	proxyAddr[1] = os.Args[2]

	proxyListenr, err := net.Listen("tcp", string(proxyAddr[0]))
	if err != nil {
		log.Fatal("Listen error")
	}
	for {
		log.Println("waiting to proxy...")
		c, e := proxyListenr.Accept()
		if e == nil {
			log.Printf("Connect From: %s\n", c.RemoteAddr())
			targetConn, e := net.Dial("tcp", string(proxyAddr[1]))
			if e == nil {
				go proxyRequest(c, targetConn)
				go proxyRequest(targetConn, c)
				log.Printf("Proxy With: %s\n", c.RemoteAddr())
			} else {
				log.Println("target Connect error")
			}
		}
	}
}
