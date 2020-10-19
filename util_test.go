package batchutil

import (
	"errors"
	"log"
	"testing"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestUtil_Run(t *testing.T) {
	type fields struct {
		config *Config
	}
	type args struct {
		f func(min, max int64) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errCnt  int
	}{
		{
			name: "return errors",
			fields: fields{
				config: &Config{
					ConcurrentLimit: 100,
					StartNumber:     1,
					EndNumber:       1450,
					BatchSize:       100,
				},
			},
			args: args{
				func(min, max int64) error {
					if min == 101 || min == 501 {
						return errors.New("error")
					}
					log.Printf("min: %d, max: %d", min, max)
					return nil
				},
			},
			wantErr: true,
			errCnt:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			util := &Util{
				config: tt.fields.config,
			}
			err := util.Run(tt.args.f)
			if err != nil != tt.wantErr {
				t.Errorf("Util.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			merr := err.(*multierror.Error)
			assert.Equal(t, tt.errCnt, merr.Len())
		})
	}
}
