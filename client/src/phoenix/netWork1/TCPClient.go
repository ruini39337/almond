package netWork1

import (
	"bufio"
	"fmt"
	"net"
	"time"
)


type Client struct {
	tcpAddr *net.TCPAddr
	quitSemaphore chan bool
}

func (this * Client)Run()(){
	tcpAddr ,_ := net .ResolveTCPAddr("tcp","127.0.0.1:9998")

	conn,_ := net.DialTCP("tcp",nil,tcpAddr)

	defer conn.Close()

	fmt.Println("conneted!")

	go this.onMessageRecived(conn)

	b := []byte("time\n")

	conn.Write(b)
	<-this.quitSemaphore
}

func (this * Client)onMessageRecived (conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg ,err := reader.ReadString('\n')
		fmt.Println("1: ",msg,err)
		if err != nil {
			this.quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
