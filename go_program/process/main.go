package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var cr cleanResources

	// deal with book resource.
	book := "book"
	if err := dealWithResource(book); err != nil {
		cr.clean()
		log.Println(err)
		return
	}
	cr.register(func() { fmt.Printf("CLEAN UP: %s\n", book) })

	// deal with pen resource.
	phone := "phone"
	if err := dealWithResource(phone); err != nil {
		cr.clean()
		log.Println(err)
		return
	}
	cr.register(func() { fmt.Printf("CLEAN UP: %s\n", phone) })

	// deal with compute.
	compute := "compute"
	if err := dealWithResource(compute); err != nil {
		cr.clean()
		log.Println(err)
		return
	}
	cr.register(func() { fmt.Printf("CLEAN UP: %s\n", compute) })

	// deal with compute.
	ipad := "ipad"
	if err := dealWithResource(ipad); err != nil {
		cr.clean()
		log.Println(err)
		return
	}
	cr.register(func() { fmt.Printf("CLEAN UP: %s\n", ipad) })

	// wait goroutines finish.
	time.Sleep(2 * time.Second)
}

func dealWithResource(r string) error {
	if r == "ipad" {
		return fmt.Errorf("deal with resource %s failed", r)
	}
	fmt.Printf("SUCCESS: deal with resource %s\n", r)
	return nil
}

type cleanResources struct {
	cleanFuncs []func()
}

func (cr *cleanResources) register(c func()) {
	if cr.cleanFuncs == nil {
		cr.cleanFuncs = make([]func(), 0, 1)
	}
	cr.cleanFuncs = append(cr.cleanFuncs, c)
}

func (cr *cleanResources) clean() {
	for _, cf := range cr.cleanFuncs {
		go cf()
	}
	cr.cleanFuncs = nil
}
