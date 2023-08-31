package transaction

import "goaplication/user"

type GetTransactionsCampaignInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount 		int `json:"amount" binding:"required"`
	CampaignID 	string `json:"campaign_id" binding:"required"`
	User 		user.User	
}

type TransactionNotificationInput struct {
	TransationStatus 	string `json:"transation_status"`
	OrderID 			string `json:"order_id"`
	PaymentType 		string `json:"payment_type"`
	FraudStatus 		string `json:"fraud_status"`
}