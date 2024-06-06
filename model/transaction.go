package model

import "time"

type Transaction struct {
	ID              string ` gorm:"primaryKey" `
	AccountID       string ` gorm:"column:account_id" `
	BankID          string ` gorm:"column:bank_id" `
	Amount          int
	TransactionDate time.Time
}

func (pi *Transaction) TableName() string {
	return "account"
}
