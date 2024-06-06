package model

type Account struct {
	AccountID string //` gorm: "primaryKey" `
	Username  string //` gorm: "coloumn:username" `
	Password  string // ini otomatis membacanya huruf kecil
	Name      string
}

func (a *Account) TableName() string {
	return "account"
}
