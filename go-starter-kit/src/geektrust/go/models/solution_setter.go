package models

import (
	"strconv"
)

// The things we do to prevent "magic numbers"
var sub_total_string string = "SUB_TOTAL "
var coupon_discount_string string = "COUPON_DISCOUNT "
var total_pro_discount_string string = "TOTAL_PRO_DISCOUNT "
var pro_membership_fee_string string = "PRO_MEMBERSHIP_FEE "
var enrollment_fee_string string = "ENROLLMENT_FEE "
var total_string string = "TOTAL "

func (app *App) Set_Solution(sub_total float64, coupon_name string, coupon_amount float64, total_pro_discount float64, pro_membership_fee float64, enrollment_fee float64, total float64) {
	// Remove magic numbers
	num_bits := 64
	accuracy_needed := 2
	var f byte = 'f'
	var space string = " "
	app.Test_array = append(app.Test_array, sub_total_string+strconv.FormatFloat(sub_total, f, accuracy_needed, num_bits))
	app.Test_array = append(app.Test_array, coupon_discount_string+coupon_name+space+strconv.FormatFloat(coupon_amount, f, accuracy_needed, num_bits))
	app.Test_array = append(app.Test_array, total_pro_discount_string+strconv.FormatFloat(total_pro_discount, f, accuracy_needed, num_bits))
	app.Test_array = append(app.Test_array, pro_membership_fee_string+strconv.FormatFloat(pro_membership_fee, f, accuracy_needed, num_bits))
	app.Test_array = append(app.Test_array, enrollment_fee_string+strconv.FormatFloat(enrollment_fee, f, accuracy_needed, num_bits))
	app.Test_array = append(app.Test_array, total_string+strconv.FormatFloat(total, f, accuracy_needed, num_bits))
}
