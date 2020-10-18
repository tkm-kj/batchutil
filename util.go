package batchutil

import "sync"

type Util struct {
	config *Config
}

func Open(config *Config) *Util {
	if config == nil {
		config = &Config{}
	}
	return &Util{
		config: config,
	}
}

func (util *Util) Run(f func(min, max int64) error) error {
	err := util.config.validate()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	slots := make(chan struct{}, util.config.concurrentNumber())

	startNum, endNum, batchSize := util.config.StartNumber, util.config.EndNumber, util.config.BatchSize

	for i := startNum; i <= endNum; i += batchSize {
		wg.Add(1)
		slots <- struct{}{}

		go func(min, max int64) {
			defer wg.Done()
			err := f(min, max)
			if err != nil {
				// TODO: error handling
			}
			<-slots
		}(i, i+batchSize)
	}

	wg.Wait()

	return nil
}
