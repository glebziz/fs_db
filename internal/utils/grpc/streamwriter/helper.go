//go:build test

package streamwriter

//go:generate mockgen -source helper.go -destination mocks/mocks.go -typed true

type TestRequest struct {
	P []byte
}

type StreamTest interface {
	Send(req *TestRequest) error
	CloseAndRecv() (*TestRequest, error)
}
