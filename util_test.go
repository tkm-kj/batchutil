package batchutil_test

import (
	"context"
	"errors"
	"log"
	"testing"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/tkm-kj/batchutil"
)

func TestUtil_Run(t *testing.T) {
	type fields struct {
		ConcurrentLimit int
		StartNumber     int64
		EndNumber       int64
		BatchSize       int64
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
				ConcurrentLimit: 3,
				StartNumber:     1,
				EndNumber:       1450,
				BatchSize:       100,
			},
			args: args{
				func(min, max int64) error {
					if min == 101 || min == 501 {
						return errors.New("error")
					}
					return nil
				},
			},
			wantErr: true,
			errCnt:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := batchutil.NewConfig(
				tt.fields.ConcurrentLimit,
				tt.fields.StartNumber,
				tt.fields.EndNumber,
				tt.fields.BatchSize,
			)
			assert.NoError(t, err)
			util := batchutil.NewUtil(cfg)
			err = util.Run(tt.args.f)
			if err != nil != tt.wantErr {
				t.Errorf("Util.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			merr := err.(*multierror.Error)
			assert.Equal(t, tt.errCnt, merr.Len())
		})
	}
}

func TestUtil_RunWithContext(t *testing.T) {
	type fields struct {
		ConcurrentLimit int
		StartNumber     int64
		EndNumber       int64
		BatchSize       int64
	}
	type args struct {
		f func(ctx context.Context, min, max int64) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "return an error",
			fields: fields{
				ConcurrentLimit: 3,
				StartNumber:     1,
				EndNumber:       1450,
				BatchSize:       100,
			},
			args: args{
				func(ctx context.Context, min, max int64) error {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
					}

					if min == 101 {
						return errors.New("error")
					}
					log.Printf("min: %d, max: %d", min, max)
					return nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := batchutil.NewConfig(
				tt.fields.ConcurrentLimit,
				tt.fields.StartNumber,
				tt.fields.EndNumber,
				tt.fields.BatchSize,
			)
			assert.NoError(t, err)
			util := batchutil.NewUtil(cfg)
			err = util.RunWithContext(context.Background(), tt.args.f)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
