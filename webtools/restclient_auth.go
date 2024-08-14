package webtools

import (
	"encoding/base64"
)

type Authorization interface {
	Apply(*RestRequest) *RestRequest
}

type BasicAuth struct {
	Username string
	Password string
}

func (ba *BasicAuth) Apply(req *RestRequest) *RestRequest {
	if req.Headers == nil {
		req.Headers = make(map[string]string)
	}

	req.Headers["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(ba.Username+":"+ba.Password))

	return req
}
