// The cryptools package contains some commonly used cryptographic functions
package cryptools

import "crypto/tls"

func X509FromFiles(certFile, keyFile string) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func X509FromBytes(cert, key []byte) (*tls.Certificate, error) {
	outCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}
	return &outCert, nil
}
