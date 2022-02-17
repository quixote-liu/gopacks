package main

import (
	"fmt"
)

type reauthFuture struct {
	done chan struct{}
	err  error
}

func newReauthFuture() *reauthFuture {
	return &reauthFuture{
		make(chan struct{}),
		nil,
	}
}

func (f *reauthFuture) Set(err error) {
	f.err = err
	close(f.done)
}

func (f *reauthFuture) Get() error {
	<-f.done
	return f.err
}

func main() {
	rf := newReauthFuture()

	// rf.Set(errors.New("heihei"))

	err := rf.Get()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("end...")
}
