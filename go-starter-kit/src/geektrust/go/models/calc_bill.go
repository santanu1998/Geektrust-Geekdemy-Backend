package models

import (
	"geektrust/constants"
)

func (app *App) Calc_Bill() {
	// AI detects too many "magic numbers"
	var zero float64 = zero
	var coupon_name string
	sub_total, coupon_amount, total_pro_discount, pro_membership_fee := zero, zero, zero, zero
	total_pro_discount = app.Calc_pro_discount()
	if total_pro_discount > 0 {
		pro_membership_fee = constants.Membership_fee
	}
	for _, val := range app.PurchaseList {
		sub_total += app.Pricing[val]
	}
	sub_total += pro_membership_fee
	coupon_name, coupon_amount, sub_total = app.Determine_coupon(sub_total)
	app.do_whats_necessary(sub_total, coupon_name, coupon_amount, total_pro_discount, pro_membership_fee)
}
