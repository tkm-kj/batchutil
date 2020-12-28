package batchutil

import "errors"

type Config struct {
	concurrentLimit int
	startNumber     int64
	endNumber       int64
	batchSize       int64
}

func NewConfig(
	concurrentLimit int,
	startNumber int64,
	endNumber int64,
	batchSize int64,
) (*Config, error) {
	if concurrentLimit == 0 {
		return nil, errors.New("batchutil: concurrentLimit is zero")
	}
	if startNumber == 0 {
		return nil, errors.New("batchutil: startNumber is zero")
	}
	if endNumber == 0 {
		return nil, errors.New("batchutil: endNumber is zero")
	}
	if batchSize == 0 {
		return nil, errors.New("batchutil: batchSize is zero")
	}

	return &Config{
		concurrentLimit: concurrentLimit,
		startNumber:     startNumber,
		endNumber:       endNumber,
		batchSize:       batchSize,
	}, nil
}
