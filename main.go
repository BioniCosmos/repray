package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/bionicosmos/repray/cert"
	"github.com/bionicosmos/repray/config"
	"github.com/bionicosmos/repray/transport"
)

func main() {
	config, err := config.FromArgs()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range config {
		c := c

		handler := &httputil.ReverseProxy{
			Rewrite: func(r *httputil.ProxyRequest) {
				r.SetURL(c.Upstream)
			},
			Transport: transport.WithH2c(),
		}

		log.Print("The proxy server is listening on ", c.Listen)
		if c.TLS != nil {
			listener, err := tls.Listen("tcp", c.Listen, &tls.Config{
				GetCertificate: cert.GetCertificateFunc(&c),
			})
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				log.Fatal(http.Serve(listener, handler))
			}()
		} else {
			go func() {
				log.Fatal(http.ListenAndServe(c.Listen, handler))
			}()
		}
	}

	select {}
}
