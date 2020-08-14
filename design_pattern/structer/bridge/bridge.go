package bridge

import (
	"fmt"
)

/*一个告警功能类，有不同的告警级别(紧急，普通，预警)，告警方式(邮件，短信，微信)
不通的告警级别对应不同的告警方式
*/

const (
	//SEVERE 正常
	SEVERE = iota
	//URGENCY 紧急
	URGENCY
	//NORMAL 普通
	NORMAL
	//TRIVIAL 轻微
	TRIVIAL
)

//Notification 告警
type Notification struct {
	emailAddresses []string
	telephones     []string
	wechatIDs      []string
}

// notify 告警函数
func (n *Notification) notify(level int, message string) {
	if level == SEVERE {
		// 发短信
	} else if level == URGENCY {
		// 发微信
	} else if level == NORMAL {
		// 发邮件
	} else if level == TRIVIAL {
		// 发邮件
	}
}

/*桥接模式
什么情况下适用桥接模式？
	当一个类存在两个或者多个变化维度，在Notification中有告警级别和告警方式两个维度
	可以把两个维度拆分开来，然后通过组合的方式使得这两个维度可以独立扩展
*/

// 下面就通过桥接模式改造一下

//MsgSender 告警消息发送接口
type MsgSender interface {
	send(message string)
}

//TelephoneMsgSender 电话
type TelephoneMsgSender struct {
	telephones []string
}

func (s *TelephoneMsgSender) send(message string) {
	fmt.Println("电话短信通知...")
}

//其他的Sender就不一一实现了，差不多类似

//INotification 告警接口
type INotification interface {
	notify(message string)
}

//SevereNotification SEVER级别的通知
type SevereNotification struct {
	msgSender MsgSender
}

func (n *SevereNotification) notify(message string) {
	n.msgSender.send(message)
}

// 其他的 Notification就不一一实现了，差不多类似
