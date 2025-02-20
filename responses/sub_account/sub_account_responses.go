package sub_account

import (
	"github.com/amir-the-h/okex/models/account"
	models "github.com/amir-the-h/okex/models/subaccount"
	"github.com/amir-the-h/okex/responses"
)

type (
	ViewList struct {
		responses.Basic
		SubAccounts []*models.SubAccount `json:"data,omitempty"`
	}
	APIKey struct {
		responses.Basic
		APIKeys []*models.APIKey `json:"data,omitempty"`
	}
	GetBalance struct {
		responses.Basic
		Balances []*account.Balance `json:"data,omitempty"`
	}
	GetAssetBalance struct {
		responses.Basic
		Balances []*account.AssetBalance `json:"data,omitempty"`
	}
	HistoryTransfer struct {
		responses.Basic
		HistoryTransfers []*models.HistoryTransfer `json:"data,omitempty"`
	}
	ManageTransfer struct {
		responses.Basic
		Transfers []*models.Transfer `json:"data,omitempty"`
	}
	GetMaxWithdrawals struct {
		responses.Basic
		MaxWithdrawals []*models.MaxWithdrawal `json:"data,omitempty"`
	}
)
