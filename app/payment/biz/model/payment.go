package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id,omitempty"`
	OrderId       string    `json:"order_id,omitempty"`
	TransactionId string    `json:"transaction_id,omitempty"`
	Amount        float32   `json:"amount,omitempty"`
	PayAt         time.Time `json:"pay_at,omitempty"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(ctx context.Context, db *gorm.DB, payment *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}
