package batchutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigGetConcurrentNumber(t *testing.T) {
	config1 := Config{}
	assert.Equal(t, 1, config1.concurrentLimit())

	config2 := Config{
		ConcurrentLimit: 1000,
	}
	assert.Equal(t, 1000, config2.concurrentLimit())
}

func TestConfigValidate(t *testing.T) {
	config1 := &Config{
		ConcurrentLimit: 100,
		StartNumber:     1,
		EndNumber:       1450,
		BatchSize:       100,
	}
	err := config1.validate()
	assert.NoError(t, err)

	config2 := &Config{
		ConcurrentLimit: 100,
		EndNumber:       1450,
		BatchSize:       100,
	}
	err = config2.validate()
	assert.Error(t, err)

	config3 := &Config{
		ConcurrentLimit: 100,
		StartNumber:     1,
		BatchSize:       100,
	}
	err = config3.validate()
	assert.Error(t, err)

	config4 := &Config{
		ConcurrentLimit: 100,
		StartNumber:     1,
		EndNumber:       1450,
	}
	err = config4.validate()
	assert.Error(t, err)
}
