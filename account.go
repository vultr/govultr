package govultr

import (
	"context"
	"net/http"
)

type AccountService interface {
	GetInfo(ctx context.Context) (*Account, error)
}

type AccountServiceHandler struct {
	client *Client
}

type Account struct {
	Balance           string `json:"balance"`
	PendingCharges    string `json:"pending_charges"`
	LastPaymentDate   string `json:"last_payment_date"`
	LastPaymentAmount string `json:"last_payment_amount"`
}

func (a *AccountServiceHandler) GetInfo(ctx context.Context) (*Account, error) {

	uri := "/v1/account/info"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	account := new(Account)
	err = a.client.DoWithContext(ctx, req, account)

	if err != nil {
		return nil, err
	}

	return account, nil
}
