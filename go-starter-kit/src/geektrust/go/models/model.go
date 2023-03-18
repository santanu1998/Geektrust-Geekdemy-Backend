package models

type App struct {
	Pricing         map[string]float64 // Stores the prices of various things a student can buy
	PurchaseList    []string           // List of purchases by the student in a single transaction
	CouponList      []string           // Coupons applied, will be processed before applying
	ProDiscounts    map[string]float64 // Percentage of discounts for pro membership
	CouponDiscounts map[string]float64 // Percentage of discounts for each coupon code
	IsProMember     bool               // Defaults to false, can be set true from input
	ToPrint         bool               // If this flag is false, the solution is not printed

	// This array is for testing.
	Test_array []string
}
