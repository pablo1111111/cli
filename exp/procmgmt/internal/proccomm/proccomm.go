package proccomm

import (
	"errors"
	"fmt"
	"net"
	"os"
)

var (
	network  = "unix"
	msgWidth = 64
)

type ProcComm struct {
	f    string
	id   string
	addr string
}

func New(name, id, addr string) *ProcComm {
	return &ProcComm{
		f:    fmt.Sprintf("/tmp/%s-%s.sock", name, id),
		id:   id,
		addr: addr,
	}
}

func (c *ProcComm) Listen(done chan struct{}) error {
	emsg := "listen: %w"

	l, err := net.Listen(network, c.f)
	if err != nil {
		return fmt.Errorf(emsg, err)
	}
	defer l.Close()

	if err = os.Chmod(c.f, 0700); err != nil {
		return fmt.Errorf(emsg, err)
	}

	conns := make(chan net.Conn)
	for {
		go func() {
			conn, err := l.Accept()
			if err != nil {
				select {
				case <-done:
					return
				default:
					fmt.Println(fmt.Errorf(emsg, err)) // wire this for return
					return
				}
			}
			conns <- conn
		}()

		select {
		case <-done:
			return errors.New("done signaled")
		case conn := <-conns:
			go func() {
				defer conn.Close()

				buf := make([]byte, msgWidth)
				n, _ := conn.Read(buf) //nolint // add error and timeout handling

				requestedID := string(buf[:n])

				if requestedID != c.id { // not really necessary - socket already ensures this
					return
				}

				_, err = conn.Write([]byte(c.addr))
				if err != nil {
					fmt.Println(fmt.Errorf(emsg, err)) // wire this for return
					return
				}
			}()
		}
	}
}

func (c *ProcComm) HTTPAddr() (string, error) {
	emsg := "check ok: %w"

	conn, err := net.Dial(network, c.f)
	if err != nil {
		return "", fmt.Errorf(emsg, err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(c.id))
	if err != nil {
		return "", fmt.Errorf(emsg, err)
	}

	buf := make([]byte, msgWidth)
	n, _ := conn.Read(buf) //nolint // add error and timeout handling

	addr := string(buf[:n])

	return addr, nil
}
