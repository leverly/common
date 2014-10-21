package common

import (
	"encoding/binary"
	"fmt"
	"net"
)

var PrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIGrAgEAAiEAsSzKcNb4KZ4Vx5nw4cwrbCa1zwLK+69weASPAIGO93kCAwEAAQIh
AK4+k7XX5OXhICBWaE1Yo1YtxNNeTnwR410j+cLhX8eBAhEAzvsMrTpcNpvD56Nj
jAuVSQIRANsiqcJ4fgLnYhupYu7mwLECEQCuSJaUBuA+20pKcjoQYnCBAhAs75O+
JDU65TXSFE8MTFdRAhAeD3CcJhSVipGdyYH+zDRX
-----END RSA PRIVATE KEY-----
`)

type EncryptContex struct {
	EncryptType                       uint8
	SessionKey, PublicKey, PrivateKey []byte
}

var (
	NilContex = EncryptContex{EncryptType: ZC_SEC_TYPE_NIL}
	AppContex = EncryptContex{EncryptType: ZC_SEC_TYPE_DES, SessionKey: []byte("a1b2c3d4")}
	DevContex = EncryptContex{EncryptType: ZC_SEC_TYPE_DES, SessionKey: []byte("d1e2f3g4")}
)

// read the dev packet and decrypt to the message //
func Receive(conn net.Conn, contex EncryptContex, message *Message, timeout int64) error {
	var packet Packet
	tempBuff := make([]byte, PAC_HEADER_LENGTH)
	err := MyRead(conn, int64(PAC_HEADER_LENGTH), tempBuff, timeout)
	if err != nil {
		fmt.Println("read Header error:", err)
		return err
	}
	packet.Header.PayloadLen = binary.BigEndian.Uint16(tempBuff[0:2])
	packet.Header.EncryptType = uint8(tempBuff[2])
	packet.Header.Reserved = uint8(tempBuff[3])
	if packet.Header.PayloadLen <= 0 {
		fmt.Println("check the payload length too small:", packet.Header.PayloadLen)
		return ErrInvalidMsg
	} else if packet.Header.PayloadLen > MAX_MSG_LENGTH {
		fmt.Println("check the payload length too large:", packet.Header.PayloadLen)
		return ErrMsgTooLarge
	}
	//fmt.Println("receive a packet", packet.Header.PayloadLen)
	packet.Payload = make([]byte, packet.Header.PayloadLen)
	err = MyRead(conn, int64(packet.Header.PayloadLen), packet.Payload, timeout)
	if err != nil {
		fmt.Println("read payload error:", err)
		return err
	}
	//fmt.Println("read packet succ:", packet.Header.EncryptType, packet.Header.PayloadLen)
	content, err := Decrypt(contex, packet.Payload)
	if err != nil {
		fmt.Println("Decrypt the packet failed:", err)
		return err
	}
	// fmt.Println("Decrypt the whole packet succ:", len(packet.Payload), len(content))
	err = message.Deserialize(content)
	if err != nil {
		fmt.Println("message Deserialize from buffer failed:", err)
	}
	return err
}

// encrypt the packet to message then send out //
func Send(conn net.Conn, contex EncryptContex, message *Message, timeout int64) error {
	data, err := message.Serialize()
	if err != nil {
		fmt.Println("message Serialize to buffer failed:", err)
		return err
	}
	payload, err := Encrypt(contex, data)
	if err != nil {
		fmt.Println("Encrypt the packet failed:", err)
		return err
	} else if len(payload) <= 0 {
		fmt.Println("check packet length too small:", len(payload))
		return ErrInvalidMsg
	} else if uint16(len(payload)) > MAX_MSG_LENGTH {
		fmt.Println("check the payload length too large:", len(payload))
		return ErrMsgTooLarge
	}
	//fmt.Println("Encrypt the whole packet succ:", len(data), len(payload))
	tempBuff := make([]byte, PAC_HEADER_LENGTH)
	binary.BigEndian.PutUint16(tempBuff[0:2], uint16(len(payload)))
	tempBuff[2] = byte(contex.EncryptType)
	tempBuff[3] = 0 // reserved
	err = MyWrite(conn, int64(PAC_HEADER_LENGTH), tempBuff, timeout)
	if err != nil {
		fmt.Println("write Header error:", err)
		return err
	}
	// fmt.Println("send a packet", len(payload))
	err = MyWrite(conn, int64(len(payload)), payload, timeout)
	if err != nil {
		fmt.Println("write body error:", err)
	}
	return err
}

func Encrypt(contex EncryptContex, orignal []byte) (data []byte, err error) {
	defer func() {
		temp := recover()
		if temp != nil {
			fmt.Println("panic err found in encrypt", temp)
			err = ErrEncryptMsg
		}
	}()
	switch {
	// using session key as decrpt key
	case contex.EncryptType == ZC_SEC_TYPE_DES:
		return DesEncrypt(contex.SessionKey, orignal)
	case contex.EncryptType == ZC_SEC_TYPE_AES:
		return AesEncrypt(contex.SessionKey, orignal)
	// using session key as decrpt key
	case contex.EncryptType == ZC_SEC_TYPE_RSA:
		return RsaBlockEncrpt(20, contex.PublicKey, orignal)
	// do nothing
	case contex.EncryptType == ZC_SEC_TYPE_NIL:
		return orignal, nil
	}
	return nil, ErrUnknown
}

func Decrypt(contex EncryptContex, crypted []byte) (data []byte, err error) {
	defer func() {
		temp := recover()
		if temp != nil {
			fmt.Println("panic err found in decrypt", temp)
			err = ErrDecryptMsg
		}
	}()
	switch {
	// using session key as decrpt key
	case contex.EncryptType == ZC_SEC_TYPE_DES:
		return DesDecrypt(contex.SessionKey, crypted)
	case contex.EncryptType == ZC_SEC_TYPE_AES:
		return AesDecrypt(contex.SessionKey, crypted)
	// using private key as decrpt key
	case contex.EncryptType == ZC_SEC_TYPE_RSA:
		return RsaBlockDecrpt(32, contex.PrivateKey, crypted)
	// do nothing
	case contex.EncryptType == ZC_SEC_TYPE_NIL:
		return crypted, nil
	}
	return nil, ErrUnknown
}
