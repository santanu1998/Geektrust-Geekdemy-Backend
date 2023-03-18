package utils

import (
	"testing"
)

func Test_if_contains(t *testing.T) {
	l1 := []string{"apple", "orange", "grapes", "guava", "mango", "pear"}
	t1 := "apple"
	r1 := Check_if_contains(l1, t1)
	l2 := []string{"cats", "dogs", "parrots", "horses"}
	t2 := "giraffes"
	r2 := Check_if_contains(l2, t2)

	error_msg := "The function \"Check_if_contains\" fails\n"
	success_msg := "Test Successful"

	if r1 == false {
		t.Errorf(error_msg)
	} else {
		t.Logf(success_msg)
	}

	if r2 == true {
		t.Errorf(error_msg)
	} else {
		t.Logf(success_msg)
	}
}
