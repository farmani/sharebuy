package main

type envelope struct {
	Status string                 `json:"status" xml:"status"`
	Code   int                    `json:"code" xml:"code"`
	Data   map[string]interface{} `json:"data" xml:"data"`
}
