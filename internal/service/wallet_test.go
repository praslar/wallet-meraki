package service

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"wallet/internal/model"
	"wallet/internal/repo"
	mock_repo "wallet/internal/repo/mock"
)

//func TestWalletService_CreateWallet(t *testing.T) {
//	type fields struct {
//		walletRepo repo.WalletRepoInterface
//	}
//	crtl := gomock.NewController(t)
//	walletRepoMock := mock_repo.NewMockWalletRepoInterface(crtl)
//
//	type args struct {
//		name    string
//		xuserid string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "case 1 success",
//			fields: fields{
//				walletRepo: walletRepoMock,
//			},
//			args: args{
//				name:    "",
//				xuserid: "",
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &WalletService{
//				walletRepo: tt.fields.walletRepo,
//			}
//			if err := s.CreateWallet(tt.args.name, tt.args.xuserid); (err != nil) != tt.wantErr {
//				t.Errorf("CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestWalletService_GetOneWallet(t *testing.T) {
	type fields struct {
		walletRepo repo.WalletRepoInterface
	}
	var wallet []model.Wallet

	userID := "5c303f1d-361e-4e54-aa5d-d7df25ba5013"
	walletName := "Wallet 1"
	crtl := gomock.NewController(t)
	walletRepoMock := mock_repo.NewMockWalletRepoInterface(crtl)

	type args struct {
		userID     string
		name       string
		expectFunc func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Wallet
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1 ",
			fields: fields{
				walletRepo: walletRepoMock,
			},
			args: args{
				userID: userID,
				name:   walletName,
				expectFunc: func() {
					walletRepoMock.EXPECT().CheckWalletExist(walletName).Return(nil)
					walletRepoMock.EXPECT().GetOneWallet(walletName, userID).Return(wallet, nil)
				},
			},
			want:    wallet,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WalletService{
				walletRepo: tt.fields.walletRepo,
			}
			tt.args.expectFunc()
			got, err := s.GetOneWallet(tt.args.userID, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOneWallet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
