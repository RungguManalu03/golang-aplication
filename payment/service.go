package payment

import (
	"goaplication/user"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
    midclient.ServerKey = "SB-Mid-server-C5cMNUgrGd9b0zAJtRZYgfib"
    midclient.ClientKey = "SB-Mid-client-GPeFCtOjagoEWHdy"
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: transaction.ID,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}