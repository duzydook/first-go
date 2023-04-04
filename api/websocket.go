package api

import (
	"demo/donut"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var connPool = struct {
	sync.RWMutex
	connections map[*websocket.Conn]struct{}
}{
	connections: make(map[*websocket.Conn]struct{}),
}

var (
	upgrader = websocket.Upgrader{}
)

func Websocket(e *echo.Echo) {
	e.GET("/ws", hello)

	go func() {
		for {
			animate := donut.Animate()
			streamToAllPool(animate)
			time.Sleep(50 * time.Millisecond)
		}
	}()
}

func hello(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	connPool.Lock()
	connPool.connections[ws] = struct{}{}

	defer func(connection *websocket.Conn) {
		connPool.Lock()
		delete(connPool.connections, connection)
		connPool.Unlock()
	}(ws)

	connPool.Unlock()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}

}

func streamToAllPool(message string) error {
	connPool.RLock()
	defer connPool.RUnlock()
	for connection := range connPool.connections {
		err := connection.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	return nil
}
