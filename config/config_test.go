package config

import (
	"testing"
	"wallet/deploy"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deploy.LoadEnv()
		})
	}
}
