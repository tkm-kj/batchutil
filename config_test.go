package batchutil

import (
	"testing"
)

func TestConfig_concurrentLimit(t *testing.T) {
	type fields struct {
		ConcurrentLimit int
		StartNumber     int64
		EndNumber       int64
		BatchSize       int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "ConcurrentLimit is zero",
			fields: fields{
				ConcurrentLimit: 0,
			},
			want: 1,
		},
		{
			name: "ConcurrentLimit isn't zero",
			fields: fields{
				ConcurrentLimit: 1000,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				ConcurrentLimit: tt.fields.ConcurrentLimit,
				StartNumber:     tt.fields.StartNumber,
				EndNumber:       tt.fields.EndNumber,
				BatchSize:       tt.fields.BatchSize,
			}
			if got := config.concurrentLimit(); got != tt.want {
				t.Errorf("Config.concurrentLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_validate(t *testing.T) {
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
			name: "StartNumber is zero",
			fields: fields{
				ConcurrentLimit: 100,
				EndNumber:       1450,
				BatchSize:       100,
			},
			wantErr: true,
		},
		{
			name: "EndNumber is zero",
			fields: fields{
				ConcurrentLimit: 100,
				StartNumber:     1,
				BatchSize:       100,
			},
			wantErr: true,
		},
		{
			name: "BatchSize is zero",
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
			config := &Config{
				ConcurrentLimit: tt.fields.ConcurrentLimit,
				StartNumber:     tt.fields.StartNumber,
				EndNumber:       tt.fields.EndNumber,
				BatchSize:       tt.fields.BatchSize,
			}
			if err := config.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
