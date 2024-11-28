package app

import (
	"context"
	"fmt"
	"net"
	"sync"
)

func (a *app) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		a.server.GracefulStop()
	}()

	err = a.server.Serve(lis)
	if err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}
