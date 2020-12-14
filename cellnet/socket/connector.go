package socket

import (
	"chapter13/chatbycellnet/cellnet"
	"chapter13/chatbycellnet/cellnet/internal"
	"net"
)

type socketConnector struct {
	socketPeer //内嵌匿名结构体
	internal.SessionManager  //对于连接对象，这个总管理类无用，因为不需要管理所有连接对象

	ses cellnet.Session //type socketSession struct
}

func (c *socketConnector) Start(address string) cellnet.Peer {

	c.address = address

	go c.connect(address)

	return c
}

func (c *socketConnector) Session1() cellnet.Session {
	return c.ses
}

func (c *socketConnector) Stop() {

}

// 连接器，传入连接地址和发送封包次数
func (c *socketConnector) connect(address string) {

	// 尝试用Socket连接地址
	conn, err := net.Dial("tcp", address)

	ses := newSession(conn, &c.socketPeer)
	c.ses = ses  //结构体指针赋值给接口类型, 为何不直接定义结构体变量

	// 发生错误时退出
	if err != nil {
		c.fireEvent(ConnectErrorEvent{ses, err})
		return
	}

	log.Infof("#connected(%s) %s", c.Name(), c.address)

	ses.start()//开启三个协程 recv-loop, send-loop, wait

}

func NewConnector(f cellnet.EventFunc, q cellnet.EventQueue) cellnet.Peer {

	p := &socketConnector{
		SessionManager: internal.NewSessionManager(),
	}
	p.queue = q
	p.eventFunc = f
	p.peerInterface = p

	return p
}
