package service

import (
	"testing"
	"wallet/internal/repo"
)

func TestTokenService_CreateToken(t *testing.T) {
	type fields struct {
		TokenRepo  repo.TokenRepo
		WalletRepo repo.WalletRepo
	}
	type args struct {
		symbol string
		price  float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenService{
				TokenRepo:  tt.fields.TokenRepo,
				WalletRepo: tt.fields.WalletRepo,
			}
			if err := s.CreateToken(tt.args.symbol, tt.args.price); (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
