package cryptools

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"

	"github.com/scheiblingco/gofn/errtools"
	"software.sslmate.com/src/go-pkcs12"
)

func X509FromPkcs12(data []byte, password string) (*tls.Certificate, error) {
	privkey, cert, err := pkcs12.Decode(data, password)
	if err != nil {
		return nil, err
	}

	priv, ok := privkey.(*rsa.PrivateKey)
	if !ok {
		return nil, errtools.InvalidFieldError("private key - expected RSA private key type")
	}

	ppriv, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	pemPriv := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: ppriv,
	})

	outCert, err := tls.X509KeyPair(pemCert, pemPriv)
	if err != nil {
		return nil, err
	}

	outCert.Leaf, err = x509.ParseCertificate(cert.Raw)
	if err != nil {
		return nil, err
	}

	return &outCert, nil
}
