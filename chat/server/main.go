package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//只能写入数据的通道
type client chan<- string //对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //所有连接的客户端
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {

		case msg := <-messages:
			//把所有接收到的消息广播给所有客户端
			//发送消息通道
			//只遍历k就可以
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli) //关闭通道
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)
	go clientWriter(conn,ch)

	who := conn.RemoteAddr().String()
	ch <- "welcome" + who
	messages <- who + "online"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	//注意:忽略input.Err()中可能的错误

	leaving <- ch
	messages <- who + "下线"
	_ = conn.Close()

}

func clientWriter(conn net.Conn,ch <-chan string)  {
	for msg := range ch{
		fmt.Fprintln(conn,msg) //注意:忽略网络层面的错误
	}
}
