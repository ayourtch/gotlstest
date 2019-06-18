package main

import (
    "strings"
    "bytes"
    "log"
    "crypto/tls"
    "net"
    "bufio"
    "strconv"
)

func main() {
    log.SetFlags(log.Lshortfile)

    cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Println(err)
        return
    }

    config := &tls.Config{Certificates: []tls.Certificate{cer}}
    ln, err := tls.Listen("tcp", ":6443", config) 
    if err != nil {
        log.Println(err)
        return
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    for {
        msg, err := r.ReadString('\n')
        if err != nil {
            log.Println(err)
            return
        }
	msg = strings.TrimRight(msg, "\n")
	println("msg: '", msg, "'")
	reply_len, err := strconv.Atoi(msg)

	if err != nil {
		log.Println(err)
		return
	}

        println(msg)
	println("reply len", reply_len)

	// var reply []byte
	var b bytes.Buffer
	b.Grow(reply_len)
	b.Write([]byte("z"))

	for i := 1; i < reply_len; i++ {
	        b.Write([]byte("X"))
	}

	bb := b.Bytes()
	log.Println(b.Len())

        n, err := conn.Write(bb)
        if err != nil {
            log.Println(n, err)
            return
        }
	// conn.Close()
	log.Println("Finished writing")
    }
}

