package batchutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkm-kj/batchutil"
)

func TestConfig_NewConfig(t *testing.T) {
	type fields struct {
		ConcurrentLimit int
		StartNumber     int64
		EndNumber       int64
		BatchSize       int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "four fields fill",
			fields: fields{
				ConcurrentLimit: 100,
				StartNumber:     1,
				EndNumber:       1450,
				BatchSize:       100,
			},
			wantErr: false,
		},
		{
			name: "concurrentLimit is zero",
			fields: fields{
				StartNumber: 1,
				EndNumber:   1450,
				BatchSize:   100,
			},
			wantErr: true,
		},
		{
			name: "startNumber is zero",
			fields: fields{
				ConcurrentLimit: 100,
				EndNumber:       1450,
				BatchSize:       100,
			},
			wantErr: true,
		},
		{
			name: "endNumber is zero",
			fields: fields{
				ConcurrentLimit: 100,
				StartNumber:     1,
				BatchSize:       100,
			},
			wantErr: true,
		},
		{
			name: "batchSize is zero",
			fields: fields{
				ConcurrentLimit: 100,
				StartNumber:     1,
				EndNumber:       1450,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := batchutil.NewConfig(
				tt.fields.ConcurrentLimit,
				tt.fields.StartNumber,
				tt.fields.EndNumber,
				tt.fields.BatchSize,
			)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
