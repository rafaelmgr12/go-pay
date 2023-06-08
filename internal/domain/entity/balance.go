package entity

type Balance struct {
	ID     string
	UserID string
	Amount float64
}

func NewBalance(id string, userID string, amount float64) *Balance {
	return &Balance{
		ID:     id,
		UserID: userID,
		Amount: amount,
	}
}
