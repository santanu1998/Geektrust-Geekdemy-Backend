package models

import (
	"bufio"
	"geektrust/constants"
	"os"
	"strconv"
	"strings"
)

// Splitting Take_args() into two functions to satisfy codu.ai (smaller functions)
func (app *App) Take_args() {
	// "Avoiding magic numbers (I tried everything else)"
	zero, one := 0, 1
	cliArgs := os.Args[one:]
	filePath := cliArgs[zero]
	file, _ := os.Open(filePath)
	defer file.Close()
	// Feed inputs into the app struct
	scanner := bufio.NewScanner(file)

	app.do_things_with_args(scanner)

}

func (app *App) do_things_with_args(scanner *bufio.Scanner) {
	zero, one, two := 0, 1, 2
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if words[zero] == constants.Add_programme {
			programme_count, _ := strconv.Atoi(words[two])
			for i := zero; i < programme_count; i++ {
				app.PurchaseList = append(app.PurchaseList, words[one])
			}
		} else if words[zero] == constants.Apply_coupon {
			app.CouponList = append(app.CouponList, words[one])
		} else if words[zero] == constants.Add_pro_membership {
			app.IsProMember = true
		} else if words[zero] == constants.Print_bill {
			app.ToPrint = true
		}
	}
}
