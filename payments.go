package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentCompleted PaymentStatus = "Completed"
	PaymentFailed    PaymentStatus = "Failed"
)

// PaymentMethod defines available payment methods
type PaymentMethod string

const (
	Card           PaymentMethod = "Card"
	UPI            PaymentMethod = "UPI"
	CashOnDelivery PaymentMethod = "Cash on Delivery"
)

// Payment struct holds the payment details
type Payment struct {
	PaymentID     string
	UserID        int
	Amount        float64
	Method        PaymentMethod
	Status        PaymentStatus
	TransactionAt time.Time
}

// Simulate payment database
var paymentData = make(map[string]*Payment)

// Generate a random payment ID
func generatePaymentID() string {
	return fmt.Sprintf("PAY-%d", rand.Intn(1000000))
}

// ProcessPayment simulates payment processing
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

	// Simulate success/failure randomly (80% chance of success)
	if rand.Intn(100) < 80 {
		payment.Status = PaymentCompleted
		fmt.Println("✅ Payment Successful")
	} else {
		payment.Status = PaymentFailed
		fmt.Println("❌ Payment Failed")
	}

	// Store in payment data
	paymentData[payment.PaymentID] = payment

	return payment
}
