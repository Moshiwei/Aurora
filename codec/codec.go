package codec

import "io"

// Header 抽象数据结构Header
type Header struct {
	ServiceMethod string // format:"Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

// Codec 抽象对消息体的编解码器，抽象出接口是为了实现不同的Codec实例
// 这里居然是接口而不是结构体
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

// gob Go语言内置解码包
// 定义了两种编解码器
const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	// NewCodecFuncMap[GobType] = NewGobCodec
}
