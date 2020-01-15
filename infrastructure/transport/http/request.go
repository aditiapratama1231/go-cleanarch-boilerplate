package request

import (
	"bytes"
	"github.com/imroc/req"
)

type Request interface {
	Get(string, []byte) (*req.Resp, error)
	Post(string, []byte) (*req.Resp, error)
	Put(string, []byte) (*req.Resp, error)
	Delete(string) (*req.Resp, error)
}

type request struct {
	BaseUrl string
}

func NewRequest(baseUrl string) Request {
	return &request{
		BaseUrl: baseUrl,
	}
}

func (r *request) Send(method string, url string, body []byte) (*req.Resp, error) {
	req.Debug = true

	header := req.Header{
		"Content-Type": "application/json",
	}

	switch method {
	case "GET":
		return req.Get(r.BaseUrl+url, header, bytes.NewBuffer(body))
	case "POST":
		return req.Post(r.BaseUrl+url, header, bytes.NewBuffer(body))
	case "PUT":
		return req.Put(r.BaseUrl+url, header, bytes.NewBuffer(body))
	case "DELETE":
		return req.Delete(r.BaseUrl+url, header)
	}

	return nil, nil
}

func (r *request) Get(url string, body []byte) (*req.Resp, error) {
	return r.Send("POST", url, body)
}

func (r *request) Post(url string, body []byte) (*req.Resp, error) {
	return r.Send("POST", url, body)
}

func (r *request) Put(url string, body []byte) (*req.Resp, error) {
	return r.Send("POST", url, body)
}

func (r *request) Delete(url string) (*req.Resp, error) {
	return r.Send("POST", url, []byte(""))
}
