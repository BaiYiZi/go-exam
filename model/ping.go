package model

import "fmt"

type Ping struct {
	Msg string `json:"msg"`
}

type requestPing struct {
	Verify bool   `json:"verify"`
	Name   string `json:"name"`
}

type responsePing struct {
	Msg string `json:"msg"`
}

func (Ping) RequestPing() *requestPing {
	return &requestPing{}
}

func (Ping) ResponsePing() *responsePing {
	return &responsePing{}
}

func (rq *requestPing) VerifyParameters() (bool, error) {
	if !rq.Verify {
		return false, fmt.Errorf("parameters verify failed")
	}

	return true, nil
}
