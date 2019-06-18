package main

// go run client.go 127.0.0.1:6443 123

import (
    "os"
    "log"
    "crypto/tls"
)

func main() {
    log.SetFlags(log.Lshortfile)

    conf := &tls.Config{
         InsecureSkipVerify: true,
    }

    conn, err := tls.Dial("tcp", os.Args[1], conf)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    msg := os.Args[2] + "\n"
    n, err := conn.Write([]byte(msg))

    if err != nil {
        log.Println(n, err)
        return
    }

    buf := make([]byte, 100000)
    log.Println("Reading")
    n, err = conn.Read(buf)
    log.Println("Read")
    if err != nil {
	log.Println("Error")
        log.Println(n, err)
        return
    }
    log.Println("Got: ", n)

    println(string(buf[:n]))
}

