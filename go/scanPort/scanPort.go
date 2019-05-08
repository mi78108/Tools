package main

import (
  "fmt"
  "log"
  "net"
  "os"
  "strconv"
  "strings"
)

func connect(addr string, ch chan string) {
  s, e := net.Dial("tcp", addr)
  if e == nil {
    ch <- fmt.Sprintf("The %c[1;32m%-21s%c[0m is   Connect able --%c[1;32m%s%c[0m", 0x1B,addr,0x1B,0x1B,"True",0x1B)
    s.Close()
  } else {
    ch <- fmt.Sprintf("The %c[1;31m%-21s%c[0m is UnConnect able --%c[1;31m%s%c[0m", 0x1B,addr,0x1B,0x1B,"False",0x1B)
  }
}

func main() {
  var ports []uint16
  ch := make(chan string)
  all := false

  if len(os.Args) < 3 || (len(os.Args) < 3 && strings.Contains(string(os.Args[1]), "-a")) {
    log.Fatalln("usage: scanPort [-a] ip_addr port_range")
  }

  port_v := string(os.Args[2])
  host := string(os.Args[1])

  if strings.Contains(string(os.Args[1]), "-a") {
    port_v = string(os.Args[3])
    host = string(os.Args[2])
    all = true
  }

  if v,e:=strconv.Atoi(port_v);e==nil{
    ports=append(ports,uint16(v))
  }

  if n := strings.Index(port_v, "-"); n > 0 {
    b, _ := strconv.Atoi(port_v[:n])
    e, _ := strconv.Atoi(port_v[n+1:])
    for i := b; i <= e; i++ {
      ports = append(ports, uint16(i))
    }
  }

  if strings.Contains(string(port_v), ",") {
    ps := strings.Split(port_v, ",")
    for _, p := range ps {
      i, _ := strconv.Atoi(p)
      ports = append(ports, uint16(i))
    }
  }

  for _, v := range ports {
    go connect(host+":"+strconv.Itoa(int(v)), ch)
  }

  for i := 0; i < len(ports); i++ {
    r := string(<-ch)
    if !all && strings.Contains(r, "False") {
      continue
    }
    println(r)
  }
}
