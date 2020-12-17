package pb

import (
	"cellnet"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type pbCodec struct {
}

// 编码器的名称
func (self *pbCodec) Name() string {
	return "pb"
}

// 将结构体编码为JSON的字节数组
func (self *pbCodec) Encode(msgObj interface{}) ([]byte, error) {

	return proto.Marshal(msgObj.(protoreflect.ProtoMessage))

}

// 将JSON的字节数组解码为结构体
func (self *pbCodec) Decode(data []byte, msgObj interface{}) error {

	return proto.Unmarshal(data, msgObj.(protoreflect.ProtoMessage))
}

func init() {

	// 注册编码器
	cellnet.RegisterCodec(new(pbCodec))
	//func RegisterCodec(c Codec)
}
