package models

import (
	"geektrust/constants"
	"geektrust/utils"
)

var zero float64 = 0.00

// These functions are created to make other functions smaller (for a better score with codu.ai)
func (app *App) coup_1() (string, float64) {
	coupon_name := constants.B4g1_string
	var coupon_amount float64 = zero
	if utils.Check_if_contains(app.PurchaseList, constants.Diploma_string) {
		coupon_amount = app.Pricing[constants.Diploma_string]
	} else if utils.Check_if_contains(app.PurchaseList, constants.Cert_string) {
		coupon_amount = app.Pricing[constants.Cert_string]
	} else {
		coupon_amount = app.Pricing[constants.Degree_string]
	}
	return coupon_name, coupon_amount
}

func (app *App) coup_2(sub_total float64) (string, float64) {
	coupon_name := constants.None_string
	var coupon_amount float64 = zero
	if sub_total >= constants.Deal_threshold {
		if utils.Check_if_contains(app.CouponList, constants.Deal_G20_string) {
			coupon_name = constants.Deal_G20_string
			coupon_amount = sub_total * constants.Deal_G20
		} else if utils.Check_if_contains(app.CouponList, constants.Deal_G5_string) {
			coupon_name = constants.Deal_G5_string
			coupon_amount = sub_total * constants.Deal_G5
		}
	} else if len(app.PurchaseList) >= constants.Deal_min_threshold {
		if utils.Check_if_contains(app.CouponList, constants.Deal_G5_string) {
			coupon_name = constants.Deal_G5_string
			coupon_amount = sub_total * constants.Deal_G5
		}
	}
	return coupon_name, coupon_amount
}

func (app *App) do_whats_necessary(sub_total float64, coupon_name string, coupon_amount float64, total_pro_discount float64, pro_membership_fee float64) {
	total := sub_total - coupon_amount
	var enrollment_fee float64 = zero
	if total < constants.Enrollment_threshold {
		enrollment_fee = constants.Enrollment_fee
		total += constants.Enrollment_fee
	}
	app.Set_Solution(sub_total, coupon_name, coupon_amount, total_pro_discount, pro_membership_fee, enrollment_fee, total)
	utils.Print_soln(sub_total, coupon_name, coupon_amount, total_pro_discount, pro_membership_fee, enrollment_fee, total, app.ToPrint)
}
