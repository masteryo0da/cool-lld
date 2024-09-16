package main

import (
	"fmt"
	"time"
)

type PaymentService interface {
	ProcessPayment(amount float64, userId string) (string, error)
}

type PaymentServiceImpl struct{}

func (p PaymentServiceImpl) ProcessPayment(amount float64, userId string) (string, error) {
	return fmt.Sprintf("payment-%v-%s", time.Now().Unix(), userId), nil
}
