package test

import (
	"testing"
	"wallet/internal/repo"
	"wallet/internal/service"
)

func TestWalletService_CreateWallet(t *testing.T) {
	type fields struct {
		walletRepo repo.WalletRepo
		transSrv   service.TransactionServiceInterface
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
			s := &service.WalletService{
				WalletRepo: tt.fields.walletRepo,
				TransSrv:   tt.fields.transSrv,
			}
			if err := s.CreateWallet(tt.args.name, tt.args.xuserid); (err != nil) != tt.wantErr {
				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
