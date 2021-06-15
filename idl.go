// +build idl

package example1

import (
	"github.com/idlgo/example1/pen"
	"github.com/idlgo/example1/pineapple"
)

type service struct {
	pen.Get         `path:"/pen/get" method:"get"`
	pen.Del         `path:"/pen/del" method:"post"`
	pineapple.Check `path:"/pineapple/check" method:"post"`
}
