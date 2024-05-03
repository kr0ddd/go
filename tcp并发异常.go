
import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestSvr(t *testing.T) {
	addr := "127.0.0.1:12351"
	go svr(addr)
	<-time.After(time.Second)
	for i := 0; i < 1000; i++ {
		go c(addr)
	}
	<-time.After(time.Hour)
}

func svr(addr string) {
	lsn, err := net.Listen("tcp", addr)
	debugf("gate listen %v", addr)
	if err != nil {
		panic(err)
	}
	for {
		c, e := lsn.Accept()
		if e != nil {
			debugf("===>accept err:%v", e)
			continue
		}
		go handeConn(c)
	}
}

func handeConn(conn net.Conn) {
	conn.Write([]byte{1})
	<-time.After(time.Hour)
}

func c(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		debugf("c.dial err %v", err)
		return
	}
	buffer := make([]byte, 4)
	_, err = io.ReadFull(conn, buffer)
	if err != nil {
		// todo  MacOS 并发1000以上的时候 这里为什么会 read: connection reset by peer ？
		debugf("c.read err %v", err)
	}
}

func debugf(format string, args ...interface{}) {
	println(fmt.Sprintf(format, args...))
}
