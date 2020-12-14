package main
import (
	"fmt"
)

type socketPeer struct {
	id int
}
type socksession struct {

}
type SessionManager interface {

}
type Session interface {

}
type Peer interface {

}
type socketConnector struct {
	socketPeer //匿名结构体

	SessionManager//匿名接口
	ses Session //接口变量，用于接收转换的结构体指针类型
}
func (s *socketConnector)GetSesson() Session{
	return s.ses
}
func Newc() Peer{
	s := &socketConnector{}
	s.ses = &socksession{}
	return s
}
func main() {

	s := Newc()
	fmt.Printf("%T\n", s)

	//接口之间转换 Peer -> Session
	k := s.(interface {
		GetSesson() Session
	}).GetSesson()

	fmt.Printf("%T\n", k)

	//接口转换成结构体指针类型 Peer -> *socketConnector
	k2 := s.(*socketConnector).GetSesson()
	fmt.Printf("%T\n", k2)
}
/*
[gdut17@localhost cellnet_test]$ go run test.go 
*main.socketConnector
*main.socksession
*main.socksession

*/