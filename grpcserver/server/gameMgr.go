package server

import (
	"google.golang.org/grpc"
	"reflect"
)

type BaseServer interface {
	Start()
	Stop()
	Restart()
}
type GameServer struct {
	server  *grpc.Server
	Handler *Handler
}

func Init() {

	handlerclass := reflect.TypeOf((*Handler)(nil)).Elem()

	t := reflect.ValueOf(handlerclass)
	handler := t.Method(0)
	data := []byte{1, 2}
	values := reflect.ValueOf(data)
	argValues := []reflect.Value{}
	for v := 0; v < values.Len(); v++ {
		argValues[v] = values.Index(v)
	}
	handler.Call(argValues)

}

func (server *GameServer) Start() {

}

func (server *GameServer) Stop() {

}

func (server *GameServer) Restart() {

}

type Handler interface {
	Handler(data []byte)
}

type TcpHandler struct {
}

func (tcp *TcpHandler) Handler(data []byte) {

}
