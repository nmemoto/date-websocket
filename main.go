package main

import (
	"context"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.Background())
		duration := 1000 * time.Millisecond
		defer cancel()
		defer ws.Close()
		for now := range now(ctx, duration) {
			// Write
			err := websocket.Message.Send(ws, strconv.FormatInt(now.UnixNano(), 10))
			if err != nil {
				c.Logger().Error(err)
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func now(ctx context.Context, duration time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func(ch chan<- time.Time) {
		defer close(ch)
	loop:
		for now := range time.Tick(duration) {
			select {
			case <-ctx.Done():
				break loop
			default:
			}
			ch <- now
		}
	}(ch)
	return ch
}
