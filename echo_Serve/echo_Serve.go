package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func echo(conn net.Conn){
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		data,err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Read data Failed")
		}
		log.Printf("Read data: %s", data)
		writer := bufio.NewWriter(conn)
		var write_data string
		fmt.Scanln(&write_data)
		_, err = writer.WriteString(write_data)
		if err != nil {
			log.Fatalln("Write data Failed")
		}
		writer.Flush()
	}
}

func main(){
	Ear,err := net.Listen("tcp",":8045")
	if err != nil {
		log.Fatalln("The Port could be listened")
	}
	log.Println("正在监听 0.0.0.0:8045")
	for{
		conn,err := Ear.Accept()
		if err != nil {
			log.Fatalln("连接中断")
		}
		log.Println("已收到连接")
		go echo(conn)
	}
}
