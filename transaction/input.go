package transaction

import "goaplication/user"

type GetTransactionsCampaignInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}