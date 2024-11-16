package http

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func EstablishWebsocketConnection(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return nil
	}

	blockingHandleWebsocket(ws)
	return nil
}

func blockingHandleWebsocket(conn *websocket.Conn) {
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}
	}()
}
