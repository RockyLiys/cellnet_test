package proto

import (
	"chapter13/chatbycellnet/cellnet"
	_ "chapter13/chatbycellnet/cellnet/codec/json"
	"reflect"
)

// 请求发送聊天内容
type ChatREQ struct {
	Content string // 发送的内容
}

// 有人发送聊天
type ChatACK struct {
	Content string // 发送的内容
	ID      int64  // 发送者的ID
}

//init函数，外部包import这个包的时候自动调用这个函数
func init() {

	cellnet.RegisterMessageMeta("json", "proto.ChatREQ", reflect.TypeOf((*ChatREQ)(nil)).Elem(), 1)
	cellnet.RegisterMessageMeta("json", "proto.ChatACK", reflect.TypeOf((*ChatACK)(nil)).Elem(), 2)
}
