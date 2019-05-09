package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"net"
	"os"
)

func relay(cf net.Conn, ct net.Conn, ch chan bool) {
	buffer := make([]byte, 10240)
	for {
		n, e := cf.Read(buffer)
		if e != nil {
			ch <- true
			break
		}
		_, e = ct.Write(buffer[:n])
		if e != nil {
			ch <- true
			break
		}
	}
}

func session(room *Room) {
	s_id, _ := uuid.NewV4()
	defer func() {
		room.c1.Close();
		room.c2.Close();
		room.state = true;
		room.c1 = nil;
		room.c2 = nil;
		log.Println(fmt.Sprint(s_id) + " Session End")
	}()
	ch := make(chan bool)
	go relay(room.c1, room.c2, ch)
	go relay(room.c2, room.c1, ch)
	log.Println(fmt.Sprint(s_id) + " Session begin")
	for {
		if <-ch {
			break
		}
	}
}

func listener(addr string, ch chan net.Conn) {
	p, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("Listen on " + addr + " Err")
	}
	for {
		c, e := p.Accept()
		if e == nil {
			log.Println(fmt.Sprintf("%s connected", c.RemoteAddr()))
			ch <- c
		}
	}
}

func inRoom(rooms map[string]*Room, conn net.Conn) *Room {
	buff := make([]byte, 10240)
	var ret *Room
	n, e := conn.Read(buff)
	if e != nil {
		log.Println(fmt.Sprintf("%s read Err", conn.RemoteAddr()))
		conn.Close()
		return ret
	}
	room := rooms[string(buff[:n])]
	if room != nil {
		if room.state {
			if room.c1 == nil {
				room.c1 = conn
			} else {
				room.c2 = conn
				room.state = false
				ret = room
			}
		}
	} else {
		r := Room{true, conn, nil}
		rooms[string(buff[:n])] = &r
	}
	return ret
}

type Room struct {
	state bool
	c1    net.Conn
	c2    net.Conn
}

var rooms map[string]*Room

func main() {

	rooms = make(map[string]*Room)
	listen := make(chan net.Conn)
	//var addrs = [2]string{"127.0.0.1:1309", "127.0.0.1:1310"}
	if len(os.Args) < 2 {
		log.Fatal("Usage: command addr1 addr2")
	}
	addrs := os.Args[1:]

	for _, l := range addrs {
		go listener(l, listen)
	}

	for {
		c := <-listen
		room := inRoom(rooms, c)
		if room != nil && !room.state {
			go session(room)
		}
	}
}
