package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/samber/lo"
	"github.com/vorprog/go-api/util"
	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

const timeout = time.Second * 10

var protocolHandlers = map[string]func(ctx context.Context, conn *websocket.Conn, rateLimiter *rate.Limiter) error{
	"echo": handleEcho,
}

func serveWebsocket(responseWriter http.ResponseWriter, request *http.Request) error {
	conn, err := websocket.Accept(responseWriter, request, &websocket.AcceptOptions{
		Subprotocols: lo.Keys(protocolHandlers),
	})

	if err != nil {
		return err
	}

	defer conn.Close(websocket.StatusInternalError, "Closing connection")

	handler := protocolHandlers[conn.Subprotocol()]

	if handler == nil {
		conn.Close(websocket.StatusPolicyViolation, "Unsupported subprotocol")
		return nil
	}

	rateLimiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)

	for {
		err = handler(request.Context(), conn, rateLimiter)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return err
		}
		if err != nil {
			util.Log("failed to echo with %v: %v", request.RemoteAddr, err)
			return err
		}
	}
}

func handleEcho(ctx context.Context, conn *websocket.Conn, rateLimiter *rate.Limiter) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := rateLimiter.Wait(ctx)
	if err != nil {
		return err
	}

	typ, r, err := conn.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := conn.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	err = w.Close()
	return err
}
