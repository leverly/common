package common

import (
	"bytes"
	"testing"
)

func TestSerialize(t *testing.T) {
	var msg, temp Message
	msg.Header.MsgId = 100
	msg.Header.MsgCode = 101
	msg.Header.Version = 102
	// empty body
	msg.Header.PayloadLen = 0
	buffer, err := msg.Serialize()
	if err != nil {
		t.Error("empty msg1 serialize failed:", err)
	}
	err = temp.Deserialize(buffer)
	if err != nil {
		t.Error("msg1 deserialize failed:", err)
	}
	if temp.Header != msg.Header {
		t.Error("check msg1 deserialize header failed")
	}
	if len(temp.Payload) != len(msg.Payload) {
		t.Error("check msg2 deserialize payload len failed")
	}
	if !bytes.Equal(temp.Payload, msg.Payload) {
		t.Error("check msg1 deserialize header failed")
	}

	// fill some body
	msg.Header.MsgId = 255
	msg.Header.MsgCode = 8
	msg.Header.Version = 13
	msg.Header.PayloadLen = 10240
	msg.Payload = make([]byte, 10240)
	copy(msg.Payload, GenerateRandomKey(10240))
	buffer2, err := msg.Serialize()
	if err != nil {
		t.Error("msg2 serialize failed:", err)
	}
	err = temp.Deserialize(buffer2)
	if err != nil {
		t.Error("msg2 deserialize failed:", err)
	}
	if temp.Header != msg.Header {
		t.Error("check msg2 deserialize header failed")
	}
	if len(temp.Payload) != len(msg.Payload) {
		t.Error("check msg2 deserialize payload len failed")
	}
	if !bytes.Equal(temp.Payload, msg.Payload) {
		t.Error("check msg2 deserialize header failed")
	}

	// old buffer exist
	err = temp.Deserialize(buffer)
	if err != nil {
		t.Error("msg1 deserialize failed:", err)
	}
	if temp.Header == msg.Header {
		t.Error("check msg2 deserialize header failed")
	}
}
