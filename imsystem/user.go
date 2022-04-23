package main

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	address := conn.RemoteAddr().String()
	user := &User{
		Name:   address,
		Addr:   address,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		mesage := <-this.C
		this.conn.Write([]byte(mesage))
	}
}

func (this *User) Online() {
	//用户上线
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播消息
	this.server.broderMessage(this, "当前用户上线了")
}
func (this *User) OffLine() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap[this.Name])
	this.server.mapLock.Unlock()

	this.server.broderMessage(this, "下线")
}

func (this *User) sendMessage(msg string) {

	this.conn.Write([]byte(msg))
}

//处理消息
func (this *User) Domessage(msg string) {
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMessage := "[" + user.Addr + "]" + user.Name + ":在线...\n"
			this.sendMessage(onlineMessage)
		}
		this.server.mapLock.UnLock()
	} else {
		this.server.broderMessage(this, msg)
	}
}
