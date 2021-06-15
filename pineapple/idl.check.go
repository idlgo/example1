// +build idl

package pineapple

type Request struct {
	Params map[string]Param
}

type Param struct {
	Value string
	Valid bool
}

type Response struct {
	Code    int
}

type (
	Check func(Request) Response
)
