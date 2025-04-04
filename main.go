package main

import (
	"fmt"
)

func main() {
	userID := 101
	amount := 499.99
	method := UPI

	fmt.Println("🛒 Order placed successfully")
	fmt.Printf("💰 Initiating payment of ₹%.2f using %s...\n", amount, method)

	payment := ProcessPayment(userID, amount, method)

	fmt.Println("----------------------------")
	fmt.Printf("📄 Payment ID: %s\n", payment.PaymentID)
	fmt.Printf("👤 User ID: %d\n", payment.UserID)
	fmt.Printf("💳 Payment Method: %s\n", payment.Method)
	fmt.Printf("💸 Amount: ₹%.2f\n", payment.Amount)
	fmt.Printf("📆 Date: %s\n", payment.TransactionAt.Format("02-Jan-2006 15:04:05"))
	fmt.Printf("✅ Status: %s\n", payment.Status)
	fmt.Println("----------------------------")
}
