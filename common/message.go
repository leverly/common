package common

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

const (
	ZC_HS_MSG_LEN         = 40
	ZC_HS_DEVICE_ID_LEN   = 12
	ZC_HS_SESSION_KEY_LEN = 16
)

const (
	ZC_SEC_TYPE_NIL = iota
	ZC_SEC_TYPE_RSA
	ZC_SEC_TYPE_AES
	ZC_SEC_TYPE_DES
)

const (
	ZC_CODE_SMARTCONFIG_BEGIN = iota
	ZC_CODE_SMARTCONFIG_DONE
	ZC_CODE_WIFI_CONNECT
	ZC_CODE_WIFI_DISCONNECT
	ZC_CODE_CLOUD_CONNECT
	ZC_CODE_CLOUD_DISCONNECT
	ZC_CODE_LOCAL_HANDSHAKE
	ZC_CODE_DESCRIBE
	ZC_CODE_ZDESCRIBE

	//HandShake Code
	ZC_CODE_HANDSHAKE_1
	ZC_CODE_HANDSHAKE_2
	ZC_CODE_HANDSHAKE_3
	ZC_CODE_HANDSHAKE_4

	//Respone Code
	ZC_CODE_HEARTBEAT /*no payload*/
	ZC_CODE_EMPTY     /*no payload, send by moudle when can recv another msg*/
	ZC_CODE_ACK       /*user define payload*/
	ZC_CODE_ERR       /*use ZC_ErrorMsg*/

	// OTA Code
	ZC_CODE_OTA_BEGIN
	ZC_CODE_OTA_FILE_BEGIN /*file name, len, version*/
	ZC_CODE_OTA_FILE_CHUNK
	ZC_CODE_OTA_FILE_END
	ZC_CODE_OTA_END
)

// message //
type MessageHeader struct {
	Version    uint8
	MsgId      uint8
	MsgCode    uint8
	Reserved   uint8
	PayloadLen uint16
	Checksum   [2]uint8
}

type Message struct {
	Header  MessageHeader
	Payload []byte
}

// packet //
type PacketHeader struct {
	EncryptType uint8
	Reserved    uint8
	PayloadLen  uint16
}

type Packet struct {
	Header  PacketHeader
	Payload []byte
}

const (
	MSG_HEADER_LENGTH uint16 = 8
	PAC_HEADER_LENGTH uint16 = 4
	MAX_MSG_LENGTH    uint16 = 1024
)

func (this *MessageHeader) String() string {
	ret := "ver:" + strconv.Itoa(int(this.Version))
	ret = ret + " id:" + strconv.Itoa(int(this.MsgId))
	ret = ret + " code:" + strconv.Itoa(int(this.MsgCode))
	ret = ret + " len:" + strconv.Itoa(int(this.PayloadLen))
	ret = ret + " crc:" + strconv.Itoa(int(this.Checksum[0]))
	ret = ret + " crc:" + strconv.Itoa(int(this.Checksum[1]))
	return ret
}

func (this *Message) Serialize() ([]byte, error) {
	if uint16(len(this.Payload)) != this.Header.PayloadLen {
		fmt.Println("check payload len error", this.Header.PayloadLen, len(this.Payload))
		return nil, ErrInvalidMsg
	}
	totalLen := MSG_HEADER_LENGTH + this.Header.PayloadLen
	buffer := make([]byte, totalLen)
	buffer[0] = byte(this.Header.Version)
	buffer[1] = byte(this.Header.MsgId)
	buffer[2] = byte(this.Header.MsgCode)
	buffer[3] = byte(this.Header.Reserved)
	binary.BigEndian.PutUint16(buffer[4:6], this.Header.PayloadLen)
	buffer[6] = byte(this.Header.Checksum[0])
	buffer[7] = byte(this.Header.Checksum[1])
	copy(buffer[8:], this.Payload)
	return buffer, nil
}

func (this *Message) Deserialize(buffer []byte) error {
	if len(buffer) < 8 {
		fmt.Println("check message len failed", len(buffer))
		return ErrInvalidMsg
	}
	this.Header.Version = uint8(buffer[0])
	this.Header.MsgId = uint8(buffer[1])
	this.Header.MsgCode = uint8(buffer[2])
	this.Header.Reserved = uint8(buffer[3])
	this.Header.PayloadLen = binary.BigEndian.Uint16(buffer[4:6])
	this.Header.Checksum[0] = uint8(buffer[6])
	this.Header.Checksum[1] = uint8(buffer[7])
	this.Payload = buffer[8:]
	if this.Header.PayloadLen != uint16(len(this.Payload)) {
		fmt.Println("check payload len error", this.Header.PayloadLen, len(this.Payload))
		return ErrInvalidMsg
	}
	return nil
}
