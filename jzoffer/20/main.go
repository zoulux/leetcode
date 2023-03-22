package main

import "fmt"

func main() {

	fmt.Println(isNumber("1.1+") == false)
	fmt.Println(isNumber(".-4") == false)
	fmt.Println(isNumber("-1E-16") == true)
	fmt.Println(isNumber("1+2") == false)
	fmt.Println(isNumber("3 .") == false)
	fmt.Println(isNumber(".") == false)
	fmt.Println(isNumber(". 1") == false)
	fmt.Println(isNumber("1 .") == false)
	fmt.Println(isNumber("1 ") == true)
	fmt.Println(isNumber("3.") == true)
	fmt.Println(isNumber("3. ") == true)
	fmt.Println(isNumber(".3") == true)
	fmt.Println(isNumber(" .3") == true)
	fmt.Println(isNumber("e9") == false)

	fmt.Println(isNumber("0") == true)
	fmt.Println(isNumber("e") == false)
	fmt.Println(isNumber(".") == false)
	fmt.Println(isNumber("   .1 ") == true)
	fmt.Println(isNumber("   ") == false)
	//
	fmt.Println(isNumber("+100") == true)
	fmt.Println(isNumber("5e2") == true)
	fmt.Println(isNumber("-123") == true)
	fmt.Println(isNumber("3.1416") == true)
	fmt.Println(isNumber("0123") == true)

	fmt.Println(isNumber("12e") == false)
	fmt.Println(isNumber("1a3.14") == false)
	fmt.Println(isNumber("1.2.3") == false)
	fmt.Println(isNumber("+-5") == false)
	fmt.Println(isNumber("12e+5.4") == false)

}

func isNumber(s string) bool {

	opflag := -1
	eflag := -1
	pflag := -1
	nflag := 1

	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			// 数字

			nflag = 2

			if opflag != -1 {
				opflag = 2
			}

			if eflag != -1 {
				eflag = 2
			}

			if pflag != -1 {
				pflag = 2
			}

		} else if v == ' ' {
			if eflag == 1 {
				return false
			}

			if opflag == 1 {
				return false
			}

			//有点不能有空格,数字，后续除了空格不能出现其他
			if nflag == 2 || pflag == 2 || pflag == 1 {
				for j := i; j < len(s); j++ {
					if s[j] != ' ' {
						return false
					}
				}
			}

		} else if v == '+' || v == '-' {
			if opflag != -1 {
				return false
			}

			if nflag == 2 {
				return false
			}
			if pflag == 1 {
				return false
			}

			opflag = 1
		} else if v == 'e' || v == 'E' {
			if eflag != -1 {
				return false
			}
			if nflag != 2 {
				return false
			}

			eflag = 1
			opflag = -1
			pflag = 2
			nflag = -1
		} else if v == '.' {
			if pflag != -1 {
				return false
			}

			//if nflag != 2 {
			//	return false
			//}

			pflag = 1
		} else {
			return false
		}
	}

	if opflag == 1 || eflag == 1 || nflag == 1 {
		return false
	}

	return true
}
