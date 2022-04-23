package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

//一旦有消息，就发送给所有的user
func (this *Server) ListenMessage() {
	for {

		mess := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- mess
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) broderMessage(user *User, message string) {

	fmt.Print(1)
	this.Message <- "[" + user.Addr + "]" + user.Name + ":" + message

}

func (this *Server) Handler(conn net.Conn) {
	fmt.Print("建立连接", conn.LocalAddr().String())

	user := NewUser(conn, this)

	//上线
	user.Online()

	go func() {
		buf := make([]byte, 4896)
		for {

			n, err := conn.Read(buf)
			if n == 0 {
				//下线
				user.OffLine()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Printf("connect error", err)
				return
			}

			message := string(buf[:n-1])
			//发送消息
			user.Domessage(message)
		}
	}()

	// select {}
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Print("net.listen error", err)
	}
	defer listen.Close()

	go this.ListenMessage()

	for {
		connet, err := listen.Accept()
		if err != nil {
			fmt.Print("listen.accetpt error", err)
			continue
		}
		go this.Handler(connet)

	}

	fmt.Println("启动成功")
}
