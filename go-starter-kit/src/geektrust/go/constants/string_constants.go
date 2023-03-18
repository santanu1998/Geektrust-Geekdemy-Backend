package constants

// These are created to remove "magic numbers" from the code

var Add_programme string
var Apply_coupon string
var Add_pro_membership string
var Print_bill string
var None_string string
var Cert_string string
var Degree_string string
var Diploma_string string
var B4g1_string string
var Deal_G20_string string
var Deal_G5_string string

func set_strings(lines []string) {
	// Start from index 14
	Add_programme = lines[14]
	Apply_coupon = lines[15]
	Add_pro_membership = lines[16]
	Print_bill = lines[17]
	None_string = lines[18]
	set_strings2(lines)
}

func set_strings2(lines []string) {
	// Start from 19
	Cert_string = lines[19]
	Degree_string = lines[20]
	Diploma_string = lines[21]
	B4g1_string = lines[22]
	Deal_G20_string = lines[23]
	Deal_G5_string = lines[24]
}
