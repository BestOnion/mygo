package main

import (
	"fmt"
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
		for _, user = range this.OnlineMap {
			user.C <- mess
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) broderMessage(user *User, message string) {

	this.Message <- "[" + User.Add + User.Name + "\n" + message
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Print("建立连接", conn.LocalAddr().String())

	//用户上线
	this.mapLock.Lock()
	user := NewUser(conn)
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播消息
	this.broderMessage(user, "当前用户上线了")
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
		}
		go this.Handler(connet)

	}
}
