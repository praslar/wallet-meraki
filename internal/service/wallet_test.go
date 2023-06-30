package service

import (
	"testing"
	"wallet/internal/repo"
)

func TestWalletService_CreateWallet(t *testing.T) {
	type fields struct {
		walletRepo repo.WalletRepo
		transSrv   TransactionServiceInterface
	}
	type args struct {
		name    string
		xuserid string
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
			s := &WalletService{
				walletRepo: tt.fields.walletRepo,
				transSrv:   tt.fields.transSrv,
			}
			if err := s.CreateWallet(tt.args.name, tt.args.xuserid); (err != nil) != tt.wantErr {
				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
