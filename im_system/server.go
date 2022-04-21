package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(Ip string, Port int) *Server {
	server := &Server{
		Ip,
		Port,
	}
	return server
}

func Handler(accept net.Conn) {
	fmt.Print("链接成功")
}

func (this *Server) start() {

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Print("net.listen error", err)
	}
	defer listen.Close()

	for {
		connet, err := listen.Accept()
		if err != nil {
			fmt.Print("listten.accept error", err)
		}

		go Handler(connet)

	}

}
