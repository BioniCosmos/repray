package cert

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"slices"

	"github.com/bionicosmos/repray/config"
	"github.com/caddyserver/certmagic"
)

func GetCertificateFunc(config *config.Config) func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if ok, err := hasDomainName(config.TLS.CertFile, info.ServerName); !ok {
			return nil, err
		}
		cert, err := tls.LoadX509KeyPair(config.TLS.CertFile, config.TLS.KeyFile)
		if err != nil {
			return nil, err
		}
		return &cert, nil
	}
}

type ErrDecodeCert string

func (err ErrDecodeCert) Error() string {
	return fmt.Sprintf("Failed to decode the certificate file: %v", string(err))
}

func hasDomainName(certFile string, domainName string) (bool, error) {
	certData, err := os.ReadFile(certFile)
	if err != nil {
		return false, err
	}
	pemData, _ := pem.Decode(certData)
	if err != nil {
		return false, ErrDecodeCert(certFile)
	}
	cert, err := x509.ParseCertificate(pemData.Bytes)
	if err != nil {
		return false, err
	}
	if !slices.ContainsFunc(cert.DNSNames, func(san string) bool {
		return certmagic.MatchWildcard(domainName, san)
	}) {
		return false, errNoCertificates
	}
	return true, nil
}
