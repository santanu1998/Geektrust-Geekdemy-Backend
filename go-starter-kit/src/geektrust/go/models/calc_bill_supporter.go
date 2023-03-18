package models

import (
	"geektrust/constants"
)

func (app *App) Calc_pro_discount() float64 {
	var total_pro_discount float64 = zero
	if app.IsProMember {
		newPricing := make(map[string]float64)
		for k, v := range app.Pricing {
			newPricing[k] = v * (1 - app.ProDiscounts[k])
		}

		for _, val := range app.PurchaseList {
			total_pro_discount = total_pro_discount + (app.Pricing[val] - newPricing[val])
		}
		app.Pricing = newPricing
	}
	return total_pro_discount
}

func (app *App) Determine_coupon(sub_total float64) (string, float64, float64) {
	var coupon_name string = constants.None_string
	var coupon_amount float64 = zero
	if len(app.PurchaseList) >= constants.B4G1_threshold {
		coupon_name, coupon_amount = app.coup_1()
	} else {
		coupon_name, coupon_amount = app.coup_2(sub_total)
	}
	return coupon_name, coupon_amount, sub_total
}
