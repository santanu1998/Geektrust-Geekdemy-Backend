package utils

import (
	"fmt"
)

// Simple function to check if a val is present in an array
// Could not use generics due to go version being 1.17 and generics were introduced in 1.18
// Not sure if the functionality is inbuilt in go, this seemed easier and faster to implement
func Check_if_contains(list []string, s string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}

// Function to print the solution out.
// Written as a seperate function to better format the code (%.2f)
func Print_soln(sub_total float64, coupon_name string, coupon_amount float64, total_pro_discount float64, pro_membership_fee float64, enrollment_fee float64, total float64, flag bool) {
	if flag {
		fmt.Printf("SUB_TOTAL %.2f\n", sub_total)
		fmt.Printf("COUPON_DISCOUNT %v %.2f\n", coupon_name, coupon_amount)
		fmt.Printf("TOTAL_PRO_DISCOUNT %.2f\n", total_pro_discount)
		fmt.Printf("PRO_MEMBERSHIP_FEE %.2f\n", pro_membership_fee)
		fmt.Printf("ENROLLMENT_FEE %.2f\n", enrollment_fee)
		fmt.Printf("TOTAL %.2f\n", total)
	}
}
