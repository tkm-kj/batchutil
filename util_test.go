package batchutil

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUtilRun(t *testing.T) {
	config := &Config{
		ConcurrentNumber: 100,
		StartNumber:      1,
		EndNumber:        1450,
		BatchSize:        100,
	}
	util := Open(config)
	f := func(min, max int64) error {
		log.Printf("min: %d, max: %d", min, max)
		return nil
	}
	err := util.Run(f)
	assert.NoError(t, err)
}
