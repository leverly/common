package common

import (
	"fmt"
	"io"
	"net"
)

//////////////////////////////////////////////////////////////////////
// blocking mode socket read/write
//////////////////////////////////////////////////////////////////////
func MyRead(conn net.Conn, packLen int64, Buffer []byte, timeout int64) (err error) {
	CheckParam(packLen <= int64(cap(Buffer)) && timeout > 0)
	var receiveLen int64
	// conn.SetReadDeadline(time.Now().UTC().Add(time.Duration(timeout)))
	for receiveLen < packLen {
		tempLen, err := conn.Read(Buffer[receiveLen:packLen])
		if err == io.EOF {
			fmt.Println("peer exit:", conn.RemoteAddr())
			return err
		} else if err != nil {
			fmt.Println("read socket connection error:", conn.RemoteAddr(), err)
			//Assert(err != os.EAGAIN, "socket not blocking mode")
			return err
		} else {
			receiveLen += int64(tempLen)
		}
	}
	return err
}

func MyWrite(conn net.Conn, packLen int64, Buffer []byte, timeout int64) (err error) {
	CheckParam(packLen <= int64(len(Buffer)) && timeout > 0)
	//conn.SetWriteDeadline(time.Now().UTC().Add(time.Duration(timeout)))
	_, err = conn.Write(Buffer[:packLen])
	if err != nil {
		fmt.Println("write socket connection error:", conn.RemoteAddr(), err)
		//Assert(err != os.EAGAIN, "socket not blocking mode")
	}
	return err
}
