package traceid

import (
	"testing"
)

func TestGenerateTraceID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"valid"},
		{"empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTraceID()
			if len(got) != 32 {
				t.Errorf("NewTraceID() = %v, want length 32", got)
			}
		})
	}
}

func TestGenerateSpanID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"valid"},
		{"empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSpanID()
			if len(got) != 16 {
				t.Errorf("NewSpanID() = %v, want length 16", got)
			}
		})
	}
}

func TestGenerateRandomBytes(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid-1", 1, false},
		{"valid-16", 16, false},
		{"zeroLength", 0, true},
		{"negativeLength", -8, true},
		{"maxInt", 1<<31 - 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandomBytes(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandomBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.length {
				t.Errorf("generateRandomBytes() = %v, want length %d", got, tt.length)
			}
		})
	}
}
