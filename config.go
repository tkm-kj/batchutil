package batchutil

import "errors"

type Config struct {
	ConcurrentLimit int
	StartNumber     int64
	EndNumber       int64
	BatchSize       int64
}

func (config *Config) concurrentLimit() int {
	if config.ConcurrentLimit == 0 {
		return 1
	}
	return config.ConcurrentLimit
}

func (config *Config) validate() error {
	if config.StartNumber == 0 {
		return errors.New("batchutil: StartNumber is zero")
	}
	if config.EndNumber == 0 {
		return errors.New("batchutil: EndNumber is zero")
	}
	if config.BatchSize == 0 {
		return errors.New("batchutil: BatchSize is zero")
	}

	return nil
}
