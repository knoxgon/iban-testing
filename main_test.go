package main

import (
	"testing"

	"github.com/knoxgon/codetest/ibanpkg"
)

/*
	Author: Volkan Güven
	Description: Test for IBAN verification
*/

//TestIBAN tests the verification of IBANs
func TestIBAN(t *testing.T) {
	//ID: 1 - Test correct IBAN group
	//Purpose: Trying to benchmark the iban with spaces and letter cases
	results1 := []bool{
		ibanpkg.ControlIban("gp1120041010050500013M02606"),
		ibanpkg.ControlIban("Gp1120041010050500013M02606"),
		ibanpkg.ControlIban("Gp1120041010050500013 M02606"),
		ibanpkg.ControlIban(" g P 1120041010050500 013 M0 2606"),
	}

	for p, v := range results1 {
		if v == false {
			t.Error("The Test ID 1 is false. The problem is at index:", p)
		}
	}

	//ID: 2 - Test different IBAN groups
	//Purpose: Trying different IBAN groups,
	//		   they should all pass except index[2] and index[5]
	results2 := []bool{
		ibanpkg.ControlIban("DE89 3704 0044 0532 0130 00"),
		ibanpkg.ControlIban("BR15 00000000000010932840814P2"),
		ibanpkg.ControlIban(""),
		ibanpkg.ControlIban("DK9520000123456789"),
		ibanpkg.ControlIban("MT31MALT 01100000 0 0 0000 000000123"),
		ibanpkg.ControlIban("MT31TALT01100000000000000000123"),
		ibanpkg.ControlIban("AE4 600900000001234567 89"),
	}

	for p, v := range results2 {
		if v == false {
			t.Error("The Test ID 2 is false. The problem is at index:", p)
		}
	}

	//ID: 3 - Test countries that are not existing in our map (getIbans())
	//		  yet have correct IBANs
	//Purpose: It should still not process correct IBAN numbers due to the lack of iso2 in the map
	results3 := []bool{
		ibanpkg.ControlIban("TG53TG0090604310346500400070"),
		ibanpkg.ControlIban("SN08SN0100152000048500003035"),
		ibanpkg.ControlIban("NE58NE0380100100130305000268"),
		ibanpkg.ControlIban("CM2110002000300277976315008"),
	}

	for p, v := range results3 {
		if v == false {
			t.Error("The Test ID 3 is false. The problem is at index:", p)
		}
	}

}
