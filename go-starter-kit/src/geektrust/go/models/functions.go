package models

import (
	"geektrust/constants"
)

func (app *App) Run() {
	app.Initialize_values()
	app.Take_args()
	app.Calc_Bill()
}

func (app *App) Initialize_values() {
	// Removing "magic numbers"
	zero := 0

	constants.Init_Constants()

	// Make new pricing hashmap
	app.Pricing = make(map[string]float64)
	app.Pricing[constants.Cert_string] = constants.Cert_cost
	app.Pricing[constants.Degree_string] = constants.Degree_cost
	app.Pricing[constants.Diploma_string] = constants.Diploma_cost

	// Make new purchase list
	app.PurchaseList = make([]string, zero)

	// Make a new coupon list (choose the best one later)
	app.CouponList = make([]string, zero)

	// Make Discount hash map
	app.ProDiscounts = make(map[string]float64)
	app.ProDiscounts[constants.Degree_string] = constants.Degree_discount
	app.ProDiscounts[constants.Cert_string] = constants.Cert_discount
	app.ProDiscounts[constants.Diploma_string] = constants.Diploma_discount
}
