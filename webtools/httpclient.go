package webtools

import (
	"crypto/tls"
	"net/http"
)

type ClientOpts interface {
	Apply(*http.Client)
}

func GetHttpClient(opts ...ClientOpts) *http.Client {
	client := &http.Client{}
	for _, opt := range opts {
		opt.Apply(client)
	}

	return client
}

type WithClientCertificate tls.Certificate

func (wtc *WithClientCertificate) Apply(client *http.Client) {
	if client.Transport == nil {
		client.Transport = &http.Transport{}
	}

	if client.Transport.(*http.Transport).TLSClientConfig == nil {
		client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{}
	}

	client.Transport.(*http.Transport).TLSClientConfig.Certificates = []tls.Certificate{tls.Certificate(*wtc)}
}
