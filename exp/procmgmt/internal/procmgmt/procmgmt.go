package procmgmt

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

type ProcMgmt struct {
	f  string
	id string
}

func New( /*path, name, id string*/ ) *ProcMgmt {
	return &ProcMgmt{
		f:  "/tmp/exp-procmgmt.sock",
		id: "tester",
	}
}

func (m *ProcMgmt) Listen(done chan struct{}) error {
	emsg := "listen: %w"

	l, err := net.Listen(network, m.f)
	if err != nil {
		return fmt.Errorf(emsg, err)
	}
	defer func() {
		l.Close()
		fmt.Println("debug closing")
		// this isn't being reached, so unlink of socket doesn't seem to happen
		// if this doesn't help reuse of the sock file, a more complex scheme is needed
	}()

	if err = os.Chmod(m.f, 0700); err != nil {
		return fmt.Errorf(emsg, err)
	}

	conns := make(chan net.Conn)
	for {
		go func() {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println(fmt.Errorf(emsg, err)) // wire this for return
				return
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
				conn.Read(buf) //nolint // add error and timeout handling

				fmt.Println(string(buf))

				_, err = conn.Write(buf)
				if err != nil {
					fmt.Println(fmt.Errorf(emsg, err)) // wire this for return
					return
				}
			}()
		}
	}
}

func (m *ProcMgmt) OK() (bool, error) {
	emsg := "check ok: %w"

	conn, err := net.Dial(network, m.f)
	if err != nil {
		return false, fmt.Errorf(emsg, err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(m.id))
	if err != nil {
		return false, fmt.Errorf(emsg, err)
	}

	buf := make([]byte, msgWidth)
	conn.Read(buf) //nolint // add error and timeout handling
	fmt.Println(string(buf))

	return string(buf) == m.id, nil
}
