package demo

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

var count = uint32(0) // 俩大爷已经预见了多少次

var total = uint32(10) //总共要预见多少次

var z0 = "吃了没，您呐？"
var z3 = "嗨，吃饱了溜溜弯。"
var z5 = "回头去给老太太请安！"

var l1 = "刚吃。"
var l2 = "您这，嘛去？"
var l4 = "有空家里坐坐啊。"

var liWriteLock sync.Mutex    //李大爷的写锁
var zhangWriteLock sync.Mutex //张大爷的写锁

type RequestResponse struct {
	Serial  uint32 //序号
	Payload string //内容
}

/*
序列化 RequestResponse ,并发送
序列化后的结构如下
长度  4字节
serial 4字节
Payload 变长
*/
func writeTo(r *RequestResponse, conn *net.TCPConn, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	payloadBytes := []byte(r.Payload)
	serialBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(serialBytes, r.Serial)

	length := uint32(len(payloadBytes) + len(serialBytes))
	lengthByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthByte, length)

	conn.Write(lengthByte)
	conn.Write(serialBytes)
	conn.Write(payloadBytes)
	//fmt.Println("发送：" + r.Payload)
}

/*
 接收消息，反序列化成RequestResponse
*/
func readFrom(conn *net.TCPConn) (*RequestResponse, error) {
	ret := &RequestResponse{}

	buff := make([]byte, 4)

	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, fmt.Errorf("读长度故障：%s", err.Error())
	}

	length := binary.BigEndian.Uint32(buff)

	if _, err := io.ReadFull(conn, buff); err != nil {
		return nil, fmt.Errorf("读Serial故障：%s", err.Error())
	}
	ret.Serial = binary.BigEndian.Uint32(buff)
	payloadBytes := make([]byte, length-4)

	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, fmt.Errorf("读Payload故障：%s", err.Error())
	}
	ret.Payload = string(payloadBytes)
	return ret, nil

}

/*
 张大爷的耳朵
*/

func zhangDaYeListen(conn *net.TCPConn) {
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
		}
		switch r.Payload {
		case l2:
			go writeTo(&RequestResponse{r.Serial, z3}, conn, &zhangWriteLock)

		case l4:
			go writeTo(&RequestResponse{r.Serial, z5}, conn, &zhangWriteLock)
		case l1:
			//如果收到刚吃，不用回复
		default:
			fmt.Println("张大爷听不懂：" + r.Payload)
		}
	}
}

/*
 张大爷的嘴
*/

func zhangDaYeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, z0}, conn, &zhangWriteLock)
		nextSerial++
	}
}

/*
李大爷的耳朵
*/

func liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if r.Payload == z0 {
			writeTo(&RequestResponse{r.Serial, l1}, conn, &liWriteLock)
		} else if r.Payload == z3 {

		} else if r.Payload == z5 {
			count++
		} else {
			fmt.Println("李大爷听不懂", r.Payload)
			break
		}
	}
}

/*
李大爷的嘴
*/
func liDaYeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, l2}, conn, &liWriteLock)
		nextSerial++
		writeTo(&RequestResponse{nextSerial, l4}, conn, &liWriteLock)
		nextSerial++
	}
}

func startServer() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListenner, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListenner.Close()
	fmt.Println("张大爷在胡同口等着...")
	for {
		conn, err := tcpListenner.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("碰见第一个李大爷：", conn.RemoteAddr().String())
		go zhangDaYeListen(conn)
		go zhangDaYeSay(conn)
	}
}

func startClient() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go liDaYeListen(conn, &wg)
	go liDaYeSay(conn)
	wg.Wait()
}

func TestServer(t *testing.T) {
	go startServer()
	time.Sleep(time.Second * 1)
	t1 := time.Now()
	startClient()
	elapsed := time.Since(t1)
	fmt.Println("耗时：", elapsed)
}
