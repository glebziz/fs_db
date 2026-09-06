package app

import (
	"context"
	"fmt"
	"net"
	"sync"
)

func (a *app) Run(ctx context.Context) error {
	var lc net.ListenConfig
	lis, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
	}()

	wg.Go(func() {
		<-ctx.Done()
		a.server.GracefulStop()
	})

	err = a.server.Serve(lis)
	if err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}
