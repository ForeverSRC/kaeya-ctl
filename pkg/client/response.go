package client

const (
	CodeSuccess       = 0
	CodeInternalError = 5000
	CodeBadRequest    = 5001
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type KVDataResponse struct {
	Response
	Data KVData `json:"data"`
}

type KVData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
