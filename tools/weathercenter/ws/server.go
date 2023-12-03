package ws

import (
	"github.com/gorilla/websocket"
	"github.com/peter-mount/go-kernel/v2/log"
	"net/http"
	"sync"
)

// Server implements a simple WebSocket server
type Server struct {
	mutex    sync.Mutex
	clients  []*websocket.Conn  // current active clients
	ch       chan []byte        // channel used when sending data to clients
	upgrader websocket.Upgrader // use default options
}

func NewServer() *Server {
	s := &Server{
		ch: make(chan []byte),
	}

	// Required to allow websockets to work over a proxy
	s.upgrader.CheckOrigin = func(r *http.Request) bool {
		// TODO add a check here to allow specific domains to work rather than everyone
		return true
	}

	return s
}

// Run performs the main loop for the WebSocket server.
// This is normally run within its own goroutine.
func (s *Server) Run() {
	for {
		msg := <-s.ch
		s.send(msg)
	}
}

func (s *Server) addClient(c *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients = append(s.clients, c)
}

func (s *Server) removeClient(c *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.removeClientImpl(c)
}

func (s *Server) removeClientImpl(c *websocket.Conn) {
	defer c.Close()

	var na []*websocket.Conn
	for _, ec := range s.clients {
		if ec != c {
			na = append(na, ec)
		}
	}
	s.clients = na
}

// Send sends a message to all connected WebSocket clients
func (s *Server) Send(msg []byte) {
	if s != nil && s.ch != nil {
		s.ch <- msg
	}
}

func (s *Server) send(msg []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, c := range s.clients {
		err := c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			s.removeClientImpl(c)
		}
	}
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer s.removeClient(c)
	s.addClient(c)

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Printf("read %d %s", mt, err)
			break
		}

		switch mt {
		case websocket.TextMessage:
			// Ignore for now

		case websocket.PingMessage:
			err = c.WriteMessage(websocket.PongMessage, msg)
			if err != nil {
				log.Printf("pong %d %s", mt, err)
				break
			}

		case websocket.PongMessage:
			// Ignore?

		default:
			// TODO implement correctly
			log.Printf("Close %d", mt)
			break
		}

	}
}
