package main

import (
	"geektrust/constants"
	"geektrust/models"
	"testing"
)

var success string = "Success"
var failure string = "Failure, not the expected output"

func Test_main_1(t *testing.T) {
	app := models.App{}
	app.Initialize_values()
	app.PurchaseList = append(app.PurchaseList, constants.Cert_string)
	app.PurchaseList = append(app.PurchaseList, constants.Degree_string)
	app.PurchaseList = append(app.PurchaseList, constants.Diploma_string)
	app.PurchaseList = append(app.PurchaseList, constants.Diploma_string)
	app.CouponList = append(app.CouponList, constants.Deal_G20_string)
	app.Calc_Bill()

	expected_ans := []string{
		"SUB_TOTAL 13000.00",
		"COUPON_DISCOUNT B4G1 2500.00",
		"TOTAL_PRO_DISCOUNT 0.00",
		"PRO_MEMBERSHIP_FEE 0.00",
		"ENROLLMENT_FEE 0.00",
		"TOTAL 10500.00",
	}

	l := len(expected_ans)

	for i := 0; i < l; i++ {
		if expected_ans[i] != app.Test_array[i] {
			t.Errorf(failure)
			return
		}
	}
	t.Logf(success)
}

func Test_main_2(t *testing.T) {
	app := models.App{}
	app.Initialize_values()
	app.PurchaseList = append(app.PurchaseList, constants.Degree_string)
	app.PurchaseList = append(app.PurchaseList, constants.Diploma_string)
	app.PurchaseList = append(app.PurchaseList, constants.Diploma_string)
	app.CouponList = append(app.CouponList, constants.Deal_G20_string)
	app.CouponList = append(app.CouponList, constants.Deal_G5_string)
	app.Calc_Bill()

	expected_ans := []string{
		"SUB_TOTAL 10000.00",
		"COUPON_DISCOUNT DEAL_G20 2000.00",
		"TOTAL_PRO_DISCOUNT 0.00",
		"PRO_MEMBERSHIP_FEE 0.00",
		"ENROLLMENT_FEE 0.00",
		"TOTAL 8000.00",
	}

	l := len(expected_ans)

	for i := 0; i < l; i++ {
		if expected_ans[i] != app.Test_array[i] {
			t.Errorf(failure)
			return
		}
	}
	t.Logf(success)
}
