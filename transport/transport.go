package transport

import (
	"crypto/tls"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

// from: https://github.com/traefik/traefik/blob/4f6c15cc14cdedc34484c697994134959fdff493/pkg/server/service/smart_roundtripper.go#L14
func WithH2c() *http.Transport {
	transport := new(http.Transport)
	transport.RegisterProtocol("h2c", &h2cTransportWrapper{
		Transport: &http2.Transport{
			DialTLS: func(network string, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
			AllowHTTP: true,
		}})
	return transport
}

type h2cTransportWrapper struct {
	*http2.Transport
}

func (t *h2cTransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	return t.Transport.RoundTrip(req)
}
