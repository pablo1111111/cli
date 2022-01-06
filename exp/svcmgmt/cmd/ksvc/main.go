// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct {
	f    *os.File
	done chan struct{}
}

func (p *program) Start(s service.Service) error {
	fpath := filepath.Join(os.TempDir(), "ksvc.log")

	f, err := os.Create(fpath)
	if err != nil {
		return err
	}

	p.f = f
	p.done = make(chan struct{})

	go p.run()
	return nil
}

func (p *program) run() {
	for {
		select {
		case <-p.done:
		default:
			fmt.Fprintln(p.f, time.Now())
			select {
			case <-time.After(time.Second * 3):
			case <-p.done:
			}
		}
	}
}

func (p *program) Stop(s service.Service) error {
	close(p.done)
	return nil
}

func main() {
	cUser, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	svcConfig := &service.Config{
		Name:        "ksvc-example",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
		UserName:    cUser.Username,
		Option: service.KeyValue{
			"UserService": true,
		},
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Install(); err != nil {
		logger.Error(err) //nolint
	}

	if err = s.Run(); err != nil {
		logger.Error(err) //nolint
	}
}
