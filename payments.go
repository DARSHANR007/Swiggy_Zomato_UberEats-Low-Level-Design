package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentCompleted PaymentStatus = "Completed"
	PaymentFailed    PaymentStatus = "Failed"
)

type PaymentMethod string

const (
	Card           PaymentMethod = "Card"
	UPI            PaymentMethod = "UPI"
	CashOnDelivery PaymentMethod = "Cash on Delivery"
)

type Payment struct {
	PaymentID     string
	UserID        int
	Amount        float64
	Method        PaymentMethod
	Status        PaymentStatus
	TransactionAt time.Time
}

var paymentData = make(map[string]*Payment)

func generatePaymentID() string {
	return fmt.Sprintf("PAY-%d", rand.Intn(1000000))
}

func ProcessPayment(userID int, amount float64, method PaymentMethod) *Payment {
	rand.Seed(time.Now().UnixNano())

	payment := &Payment{
		PaymentID:     generatePaymentID(),
		UserID:        userID,
		Amount:        amount,
		Method:        method,
		Status:        PaymentPending,
		TransactionAt: time.Now(),
	}

	if rand.Intn(100) < 80 {
		payment.Status = PaymentCompleted
		fmt.Println("✅ Payment Successful")
	} else {
		payment.Status = PaymentFailed
		fmt.Println("❌ Payment Failed")
	}

	paymentData[payment.PaymentID] = payment

	return payment
}
