package db

import (
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

type OptionFunc func(opts *options)

type options struct {
	certPath           string
	serverNameOverride string
}

func WithCertPath(certPath string) OptionFunc {
	return func(opts *options) {
		opts.certPath = certPath
	}
}

func WithServerNameOverride(serverNameOverride string) OptionFunc {
	return func(opts *options) {
		opts.serverNameOverride = serverNameOverride
	}
}

func (opts options) credentials() (credentials.TransportCredentials, error) {
	if opts.certPath == "" {
		return credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		}), nil
	}

	return credentials.NewClientTLSFromFile(opts.certPath, opts.serverNameOverride)
}
