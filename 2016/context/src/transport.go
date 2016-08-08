package http

import (
	"context"
	"net"
)

// START OMIT
type Transport struct {
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error) // HL

	// *Deprecated*: Use DialContext instead, which allows the transport...
	Dial func(network, addr string) (net.Conn, error)
	// ...
}

// END OMIT
