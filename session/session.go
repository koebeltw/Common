package session

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type Session interface {
	RemoteAddr() string
	LocalAddr() string
	Start()
	Close()
	Send(msg EventMsg)
	SendMsg(msgNo byte, subNo byte, buffer []byte)
	PushEvent(event func())

	GetID() uint16
	SetID(value uint16)

	GetBuffer() []byte

	GetConn() net.Conn

	SetData(value interface{})
	GetData() (interface{})
}

// Session 代表一个连接会话
type session struct {
	id       uint16
	index    uint16
	conn     net.Conn
	sendChan chan EventMsg
	buffer   []byte

	//terminated     bool
	//terminatedLock sync.RWMutex
	// userHandler    UserEventHandler
	// eventHandler   map[byte]map[byte]func([]byte)
	wg           *sync.WaitGroup
	eventChan    chan func()
	eventHandler EventHandler
	userHandler  UserHandler

	coder     Coder
	data      interface{}
	isReConn  bool
	closeOnce sync.Once
	readDeadline time.Time
	writeDeadline time.Time
	//pool         *sync.Pool
}

func (s *session) GetConn() net.Conn {
	return s.conn
}

func (s *session) GetBuffer() []byte {
	return s.buffer
}

func (s *session) GetID() uint16 {
	return s.id
}

func (s *session) SetID(value uint16) {
	s.id = value
}

func (s *session) SetData(value interface{}) {
	s.data = value
}

func (s *session) GetData() (interface{}) {
	return s.data
}

func (s *session) SetReadDeadline(value time.Time) {
	s.readDeadline = value
}

func (s *session) GetReadDeadline() (time.Time) {
	return s.readDeadline
}

func (s *session) SetWriteDeadline(value time.Time) {
	s.writeDeadline = value
}

func (s *session) GetWriteDeadline() (time.Time) {
	return s.writeDeadline
}

// EventMsg blabla
type EventMsg struct {
	MsgNo  byte
	SubNo  byte
	Buffer []byte
}

// RemoteAddr 返回客户端的地址和端口
func (s *session) RemoteAddr() string {
	return s.conn.RemoteAddr().String()
}

// LocalAddr 返回本机地址和端口
func (s *session) LocalAddr() string {
	return s.conn.LocalAddr().String()
}

// Start 開始
func (s *session) Start() {
	func() {
		s.wg.Add(1)
		go s.receiveThread()
		s.wg.Add(1)
		go s.sendThread()
		s.wg.Add(1)
		go s.eventThread()
	}()
}

// Close 关闭连接
func (s *session) Close() {
	s.closeOnce.Do(s.close)
}

func (s *session) close() {
	s.conn.Close()
}

func (s *session) receiveThread() {
	defer s.wg.Done()

	if s.isReConn == false {
		s.userHandler.OnUserConnect(s)
	} else {
		//s.userHandler.OnUserReConnect(s)
	}

	for {
		if err := s.conn.SetReadDeadline(time.Now().Add(time.Second * 1 * 60 * 60)); err != nil {
			log.Printf("SetReadDeadline TimeOut:%v\n", err)
			break
		}

		if msg, err := s.coder.Decode(s); err != nil {
			// log.Println(S.terminated)
			// S.terminatedLock.RLock()
			// if S.terminated {
			// 	S.terminatedLock.RUnlock()
			// 	// 直接退出
			// 	break
			// }
			// S.terminatedLock.RUnlock()

			if err != io.EOF {
				// log.Println("err != io.EOF")
				break
			}

			// log.Println("receiveThread err:", err)
			break
		} else {
			//log.Println("msg:", msg)

			if s.eventHandler[msg.MsgNo][msg.SubNo] == nil {
				log.Printf("eventHandler[%d][%d] nil \n", msg.MsgNo, msg.SubNo)
				continue
			}

			s.PushEvent(func() { s.eventHandler[msg.MsgNo][msg.SubNo](s, msg.Buffer) })
		}

		//if err := s.conn.SetReadDeadline(time.Time{}); err != nil {
		//	log.Println("SetReadDeadline Error")
		//	break
		//}
	}

	s.userHandler.OnUserDisconnect(s)
	close(s.eventChan)
	// log.Printf("Session %s receiveThread Exit", S.RemoteAddr())
}

// eventThread blabla
func (s *session) eventThread() {
	defer s.wg.Done()

	for event := range s.eventChan {
		event()
	}

	close(s.sendChan)
	// log.Printf("Session %s eventThread Exit", S.RemoteAddr())
}

func (s *session) sendThread() {
	defer s.wg.Done()

	for msg := range s.sendChan {
		if err := s.conn.SetWriteDeadline(time.Now().Add(time.Second * 60)); err != nil {
			log.Println("SetWriteDeadline TimeOut")
			break
		}

		//log.Println("sendThread:", msg)

		if buffer, err := s.coder.Encode(s, msg.MsgNo, msg.SubNo, msg.Buffer); err != nil {
			break
		} else {
			if _, err := s.conn.Write(buffer); err != nil {
				break
			}
		}

		//if err := s.conn.SetWriteDeadline(time.Time{}); err != nil {
		//	log.Println("SetWriteDeadline Error")
		//	break
		//}
	}

	// log.Printf("Session %s sendThread Exit", S.RemoteAddr())

	s.conn.Close()
}

// Send 发送数据
func (s *session) Send(msg EventMsg) {
	if s.sendChan != nil {
		s.sendChan <- msg
	}
}

// SendMsg 发送数据
func (s *session) SendMsg(msgNo byte, subNo byte, buffer []byte) {
	if s.sendChan != nil {
		s.sendChan <- EventMsg{MsgNo: msgNo, SubNo: subNo, Buffer: buffer}
	}
}

// NewSession 生成一个新的Session
func NewSession(conn net.Conn, userHandler UserHandler, coder Coder, wg *sync.WaitGroup, isReConn bool, eventHandler *EventHandler) (result Session) {
	result = &session{
		conn:     conn,
		sendChan: make(chan EventMsg, 10),
		buffer:   make([]byte, 1024*1024*4),
		//terminated: false,
		// userHandler: userHandler,
		// eventHandler: make(map[byte]map[byte]func([]byte)),
		wg:           wg,
		eventChan:    make(chan func(), 100),
		userHandler:  userHandler,
		coder:        coder,
		eventHandler: *eventHandler,
		isReConn:     isReConn,
		//pool:         pool,
	}

	if err := conn.SetReadDeadline(time.Time{}); err != nil {
		log.Println("SetReadDeadline Error")
	}

	return result
}

//PushEvent 使用者事件
func (s *session) PushEvent(event func()) {
	if s.eventChan != nil {
		s.eventChan <- event
	}
}
