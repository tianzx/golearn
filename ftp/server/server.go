package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var Addr = flag.String("h", "localhost", "ip")
var Port = flag.Int("p", 8080,
	"Bind server with assigned port")
var password string
var p string
var buflen = 4096 * 1024

type Ftp struct {
	con net.Conn
	ip  string
}

func dealWithArgs() (ip string, port int, ok bool) {
	ip = *Addr
	port = *Port
	ipOk, _ := regexp.MatchString(
		"^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)(\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)){3}$", ip)
	if !ipOk {
		fmt.Println("Invalid IPv4 Address")
		flag.PrintDefaults()
		return "", 0, false
	}
	if 0 >= port || port > 65535 {
		fmt.Println("Invalid Port Value")
		flag.PrintDefaults()
		return "", 0, false
	}
	return ip, port, true
}
func getCurrentDirectory() (string, bool) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return "", false
	}
	return strings.Replace(dir, "\\", "/", -1), true
}
func doFTPService(net.Conn) {
	curPath, suc := getCurrentDirectory()
}
func main() {
	flag.Parse()
	// fmt.Println(*add)
	ip, port, ok := dealWithArgs()
	if !ok {
		return
	}
	fmt.Printf("Server starting at %s:%d...\n", ip, port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("Server starting with an error,break down...")
		return
	}
	fmt.Println("Please set a password for connections")
	fmt.Scanln(&password)
	fmt.Println("Please set a password for connections again")
	fmt.Scanln(&p)
	if strings.Compare(password, p) != 0 {
		fmt.Println("error")
		return
	}
	fmt.Println("Password has been set, server is running...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting", err.Error())
			continue
		}
		fmt.Println("Receive connection request from ", conn.RemoteAddr().String())
	}
	go doFTPService(conn)
}
