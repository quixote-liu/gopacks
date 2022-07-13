package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// }

// type CreateResource func(interface{}) (gin.H, error)

// type DeleteResource func(id string) error

// type CreateWrap struct {
// 	cleanFuncs []func
// }

// func (cw *CreateWrap) Wrap(cr CreateResource, createArg interface{}, dr DeleteResource) (gin.H, error) {
// 	if cr == nil || dr == nil {
// 		return nil, fmt.Errorf("the argument of create or delete is nil")
// 	}

// 	r, err := cr(createArg)
// 	if err != nil {
// 		cw.clean()
// 		return r, err
// 	}

// 	rid, ok := r["id"]
// 	if !ok {
// 		log.Println("the resource not include id")
// 		return r, nil
// 	}
// 	id, ok := rid.(string)
// 	if !ok {
// 		log.Println("the id of resource is not string type")
// 		return r, nil
// 	}

// 	cw.appendCleanF(func(){

// 	})

// 	return r, nil
// }

// func (cw *CreateWrap) appendCleanF(f func()) {
// 	if cw.cleanFuncs == nil {
// 		cw.cleanFuncs = make([]func(), 0)
// 	}
// 	cw.cleanFuncs = append(cw.cleanFuncs, f)
// }

// func (cw *CreateWrap) clean() {
// 	for _, cf := range cw.cleanFuncs {
// 		go cf()
// 	}
// 	cw.cleanFuncs = nil
// }
