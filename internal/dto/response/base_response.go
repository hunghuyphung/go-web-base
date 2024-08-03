package response

type BaseResponse[T any] struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
	Data T      `json:"data"`
}
