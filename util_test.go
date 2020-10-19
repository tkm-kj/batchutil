package batchutil

import (
	"log"
	"testing"

	"errors"

	"github.com/hashicorp/go-multierror"
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
		if min == 101 || min == 501 {
			return errors.New("error")
		}
		log.Printf("min: %d, max: %d", min, max)
		return nil
	}
	err := util.Run(f)
	assert.Error(t, err)
	merr := err.(*multierror.Error)
	assert.Equal(t, 2, merr.Len())
}
