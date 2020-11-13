package common

type ResultPo struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func CreateResultPo() *ResultPo {
	return &ResultPo{
		Code:    200,
		Message: "option successful",
	}
}

func (po *ResultPo) PushData(key string, data interface{}) {
	po.Data[key] = data
}

