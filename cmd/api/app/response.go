package app

type Envelope struct {
	Status string      `json:"status" xml:"status"`
	Code   int         `json:"code" xml:"code"`
	Data   interface{} `json:"data" xml:"data"`
}
