package handler

import "wallet/internal/service"

type WalletHandler struct {
	WalletService service.WalletService
}

func NewWalletHandler(WalletService service.WalletService) WalletHandler {
	return WalletHandler{
		WalletService: WalletService,
	}
}
