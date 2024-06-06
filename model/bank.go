package model

type Bank struct {
	BankCode string ` gorm:"primaryKey" `
	Name     string ` gorm:"column:username" `
	Address  string //	ini otomatis membacanya huruf kecil
}

func (pi *Bank) TableName() string {
	return "account"
}
