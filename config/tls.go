package config

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
)

func CreateTLSConfig(caCertFile string, certFile string, keyFile string) *tls.Config {
    cfg := new(tls.Config)
    cfg.RootCAs = x509.NewCertPool()
    if ca, err := ioutil.ReadFile(caCertFile); err == nil {
        cfg.RootCAs.AppendCertsFromPEM(ca)
    }
    if cert, err := tls.LoadX509KeyPair(certFile, keyFile); err == nil {
        cfg.Certificates = append(cfg.Certificates, cert)
    }
    return cfg
}
