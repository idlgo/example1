// +build idl

package pen

type EnumStatusType struct {
	Draft, Validated, Pending, Published int
}

func (EnumStatusType) New() EnumStatusType {
	return EnumStatusType{1, 2, 3, 4}
}

type GetPenRequest struct {
	EventCode string
	Status    EnumStatusType
	Name      string
	Page      int
	Limit     int
}

type GetPenResponse struct {
	Code    int
	Message string
	Data    GetPenResponseData
}

type GetPenResponseData struct {
	Total int
	Items []GetPenResponseDataItem
}

type GetPenResponseDataItem struct {
	Name       string
	IsAsync    bool
	ConfigType int32
}

type (
	Get func(GetPenRequest) GetPenResponse
)
