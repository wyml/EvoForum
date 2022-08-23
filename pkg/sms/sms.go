package sms

import (
	"forum/pkg/config"
	"sync"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

// SMS 发送短信操作类
type SMS struct {
	Driver Driver
}

var once sync.Once

// 内部使用的 SMS 对象
var internalSMS *SMS

func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
