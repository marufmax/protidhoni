package websocket

// will implement socket letter

//
//import (
//	"github.com/gorilla/websocket"
//	"github.com/labstack/echo/v4"
//	"log"
//	"sync"
//)
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//type Manager struct {
//	clients main.ClientList
//	sync.RWMutex
//}
//
//func NewManager() *Manager {
//	return &Manager{
//		clients: make(main.ClientList),
//	}
//}
//
//func (m *Manager) serveWs(c echo.Context) error {
//	log.Println("Received new connection")
//	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
//
//	if err != nil {
//		log.Fatal(err)
//
//		return err
//	}
//
//	log.Printf("clients connected")
//
//	client := main.NewClient(ws, m)
//	m.addClient(client)
//
//	go client.readMessages()
//
//	//defer ws.Close()
//
//	return nil
//}
//
//func (m *Manager) addClient(client *main.Client) {
//	m.Lock()
//	defer m.Unlock()
//
//	m.clients[client] = true
//}
//
//func (m *Manager) removeClient(client *main.Client) {
//	m.Lock()
//	defer m.Unlock()
//
//	if _, ok := m.clients[client]; ok {
//		client.connection.Close()
//		delete(m.clients, client)
//	}
//}
