package bridge

import "testing"

func TestBridge(t *testing.T) {
	// 功能需求：SEVERE级别使用短信通知
	var n INotification
	var s MsgSender

	s = &TelephoneMsgSender{
		telephones: make([]string, 0),
	}

	n = &SevereNotification{
		msgSender: s,
	}

	n.notify("")
}
