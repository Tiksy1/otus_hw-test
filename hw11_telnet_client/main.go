package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	timeout time.Duration
	address string
	err     error
)

func flagParse() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	flag.Parse()
	if flag.NArg() > 1 {
		address = net.JoinHostPort(flag.Args()[0], flag.Args()[1])
	}
}

func main() {
	flagParse()
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err = client.Connect(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT)
	ctx, ctxCancelFunc := context.WithCancel(context.Background())
	go func() {
		defer ctxCancelFunc()
		err = client.Send()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	go func() {
		defer ctxCancelFunc()
		err = client.Receive()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	select {
	case <-signalChannel:
	case <-ctx.Done():
		close(signalChannel)
	}
}
