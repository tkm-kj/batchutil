package batchutil

import (
	"context"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

type Util struct {
	config *Config
}

func NewUtil(config *Config) *Util {
	return &Util{
		config: config,
	}
}

func (util *Util) Run(f func(min, max int64) error) error {
	var wg sync.WaitGroup
	slots := make(chan struct{}, util.config.concurrentLimit)
	defer close(slots)

	var result error

	startNum, endNum, batchSize := util.config.startNumber, util.config.endNumber, util.config.batchSize

	for i := startNum; i <= endNum; i += batchSize {
		wg.Add(1)
		slots <- struct{}{}

		go func(min, max int64) {
			defer wg.Done()
			err := f(min, max)
			if err != nil {
				result = multierror.Append(result, err)
			}
			<-slots
		}(i, i+batchSize)
	}

	wg.Wait()

	return result
}

func (util *Util) RunWithContext(ctx context.Context, f func(ctx context.Context, min, max int64) error) error {
	eg, newCtx := errgroup.WithContext(ctx)

	slots := make(chan struct{}, util.config.concurrentLimit)
	defer close(slots)

	startNum, endNum, batchSize := util.config.startNumber, util.config.endNumber, util.config.batchSize
	for i := startNum; i <= endNum; i += batchSize {
		slots <- struct{}{}

		i := i
		eg.Go(func() error {
			err := f(newCtx, i, i+batchSize)
			<-slots
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
