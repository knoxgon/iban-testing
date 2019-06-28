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
	//https://bank.codes/iban/structure/guadelope/
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

	//ID: 3 - Testing experimential countries
	//Purpose: It should process below experimential countries
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

	//ID: 4 - Testing all possible iban countries
	//Purpose: It should process all of the ibans
	results4 := []bool{
		ibanpkg.ControlIban("AL35202111090000000001234567"),
		ibanpkg.ControlIban("AD1400080001001234567890"),
		ibanpkg.ControlIban("AT483200000012345864"),
		ibanpkg.ControlIban("AZ96AZEJ00000000001234567890"),
		ibanpkg.ControlIban("BH02CITI00001077181611"),
		ibanpkg.ControlIban("BY86AKBB10100000002966000000"),
		ibanpkg.ControlIban("BE71096123456769"),
		ibanpkg.ControlIban("BA393385804800211234"),
		ibanpkg.ControlIban("BR1500000000000010932840814P2"),
		ibanpkg.ControlIban("BG18RZBB91550123456789"),
		ibanpkg.ControlIban("CR23015108410026012345"),
		ibanpkg.ControlIban("HR1723600001101234565"),
		ibanpkg.ControlIban("CY21002001950000357001234567"),
		ibanpkg.ControlIban("CZ5508000000001234567899"),
		ibanpkg.ControlIban("DK9520000123456789"),
		ibanpkg.ControlIban("DO22ACAU00000000000123456789"),
		ibanpkg.ControlIban("SV43ACAT00000000000000123123"),
		ibanpkg.ControlIban("EE471000001020145685"),
		ibanpkg.ControlIban("FO9264600123456789"),
		ibanpkg.ControlIban("FI1410093000123458"),
		ibanpkg.ControlIban("FR7630006000011234567890189"),
		ibanpkg.ControlIban("GE60NB0000000123456789"),
		ibanpkg.ControlIban("DE75512108001245126199"),
		ibanpkg.ControlIban("GI04BARC000001234567890"),
		ibanpkg.ControlIban("GR9608100010000001234567890"),
		ibanpkg.ControlIban("GL8964710123456789"),
		ibanpkg.ControlIban("GT20AGRO00000000001234567890"),
		ibanpkg.ControlIban("VA59001123000012345678"),
		ibanpkg.ControlIban("HU93116000060000000012345676"),
		ibanpkg.ControlIban("IS750001121234563108962099"),
		ibanpkg.ControlIban("IQ20CBIQ861800101010500"),
		ibanpkg.ControlIban("IE64IRCE92050112345678"),
		ibanpkg.ControlIban("IL170108000000012612345"),
		ibanpkg.ControlIban("IT60X0542811101000000123456"),
		ibanpkg.ControlIban("JO71CBJO0000000000001234567890"),
		ibanpkg.ControlIban("KZ563190000012344567"),
		ibanpkg.ControlIban("XK051212012345678906"),
		ibanpkg.ControlIban("KW81CBKU0000000000001234560101"),
		ibanpkg.ControlIban("LV97HABA0012345678910"),
		ibanpkg.ControlIban("LB92000700000000123123456123"),
		ibanpkg.ControlIban("LI7408806123456789012"),
		ibanpkg.ControlIban("LT601010012345678901"),
		ibanpkg.ControlIban("LU120010001234567891"),
		ibanpkg.ControlIban("MK07200002785123453"),
		ibanpkg.ControlIban("MT31MALT01100000000000000000123"),
		ibanpkg.ControlIban("MR1300020001010000123456753"),
		ibanpkg.ControlIban("MU43BOMM0101123456789101000MUR"),
		ibanpkg.ControlIban("MD21EX000000000001234567"),
		ibanpkg.ControlIban("MC5810096180790123456789085"),
		ibanpkg.ControlIban("ME25505000012345678951"),
		ibanpkg.ControlIban("NL02ABNA0123456789"),
		ibanpkg.ControlIban("NO8330001234567"),
		ibanpkg.ControlIban("PK36SCBL0000001123456702"),
		ibanpkg.ControlIban("PS92PALS000000000400123456702"),
		ibanpkg.ControlIban("PL10105000997603123456789123"),
		ibanpkg.ControlIban("PT50002700000001234567833"),
		ibanpkg.ControlIban("QA54QNBA000000000000693123456"),
		ibanpkg.ControlIban("RO09BCYP0000001234567890"),
		ibanpkg.ControlIban("LC14BOSL123456789012345678901234"),
		ibanpkg.ControlIban("SM76P0854009812123456789123"),
		ibanpkg.ControlIban("ST23000200000289355710148"),
		ibanpkg.ControlIban("SA4420000001234567891234"),
		ibanpkg.ControlIban("RS35105008123123123173"),
		ibanpkg.ControlIban("SC52BAHL01031234567890123456USD"),
		ibanpkg.ControlIban("SK8975000000000012345671"),
		ibanpkg.ControlIban("SI56192001234567892"),
		ibanpkg.ControlIban("ES7921000813610123456789"),
		ibanpkg.ControlIban("SE7280000810340009783242"),
		ibanpkg.ControlIban("CH5604835012345678009"),
		ibanpkg.ControlIban("TL380010012345678910106"),
		ibanpkg.ControlIban("TN5904018104004942712345"),
		ibanpkg.ControlIban("TR320010009999901234567890"),
		ibanpkg.ControlIban("UA903052992990004149123456789"),
		ibanpkg.ControlIban("AE460090000000123456789"),
		ibanpkg.ControlIban("GB33BUKB20201555555555"),
		ibanpkg.ControlIban("VG21PACG0000000123456789"),
		ibanpkg.ControlIban("DZ580002100001113000000570"),
		ibanpkg.ControlIban("AO06004400006729503010102"),
		ibanpkg.ControlIban("BJ66BJ0610100100144390000769"),
		ibanpkg.ControlIban("BF42BF0840101300463574000390"),
		ibanpkg.ControlIban("BI43201011067444"),
		ibanpkg.ControlIban("CM2110002000300277976315008"),
		ibanpkg.ControlIban("CV64000500000020108215144"),
		ibanpkg.ControlIban("CF4220001000010120069700160"),
		ibanpkg.ControlIban("TD8960002000010271091600153"),
		ibanpkg.ControlIban("KM4600005000010010904400137"),
		ibanpkg.ControlIban("CG3930011000101013451300019"),
		ibanpkg.ControlIban("DJ2110002010010409943020008"),
		ibanpkg.ControlIban("EG2100037000671002392189379"),
		ibanpkg.ControlIban("GQ7050002001003715228190196"),
		ibanpkg.ControlIban("GA2140021010032001890020126"),
		ibanpkg.ControlIban("GW04GW1430010181800637601"),
		ibanpkg.ControlIban("HN54PISA00000000000000123124"),
		ibanpkg.ControlIban("IR710570029971601460641001"),
		ibanpkg.ControlIban("CI93CI0080111301134291200589"),
		ibanpkg.ControlIban("MG4600005030071289421016045"),
		ibanpkg.ControlIban("ML13ML0160120102600100668497"),
		ibanpkg.ControlIban("MA64011519000001205000534921"),
		ibanpkg.ControlIban("MZ59000301080016367102371"),
		ibanpkg.ControlIban("NI92BAMC000000000000000003123123"),
		ibanpkg.ControlIban("NE58NE0380100100130305000268"),
		ibanpkg.ControlIban("SN08SN0100152000048500003035"),
		ibanpkg.ControlIban("TG53TG0090604310346500400070"),
	}

	for p, v := range results4 {
		if v == false {
			t.Error("The Test ID 4 is false. The problem is at index:", p)
		}
	}

}
