package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建用户
func NewUser(conn net.Conn) *User {
	address := conn.RemoteAddr().String()
	user := &User{
		Name: address,
		Addr: address,
		C:    make(chan string),
		conn: conn,
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
