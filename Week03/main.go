package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errG, ctx := errgroup.WithContext(context.Background())
	errG.Go(func() (err error) {
		errChan := make(chan error)
		serv := http.Server{Addr: "127.0.0.1:8080"}
		go func() {
			fmt.Println("server start")
			errChan <- serv.ListenAndServe()
		}()

		select {
		case e := <-errChan:
			err = errors.New("server err -> " + e.Error())
		case <-ctx.Done():
			serv.Close()
			fmt.Println("server stopped")
		}

		return
	})

	errG.Go(func() error {
		quit:=make(chan os.Signal)
		signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT,os.Interrupt)
		for {
			fmt.Println("waiting for quit signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-quit:
				return errors.New("receive quit signal")
			}
		}
	})

	if err := errG.Wait(); err != nil {
		fmt.Println(err)
	}
}
