// +build idl

package pen

type DelPenRequest struct {
	Name string
}

type DelPenResponse struct {
	Code    int
	Message string
}

type (
	Del func(DelPenRequest) DelPenResponse
)
