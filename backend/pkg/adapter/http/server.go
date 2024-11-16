package http

import "github.com/labstack/echo/v4"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	e := echo.New()

	e.GET("/session", EstablishWebsocketConnection)

	return e.Start(":8080")
}
