package util

import(
	"strconv"
)

//CnpjIsValid return a boolean valid cnpj
func CnpjIsValid(cnpj string) bool{
	cnpj = GetOnlyNumbers(cnpj)
	if len(cnpj) == 14 {
		firstTwelve := cnpj[0:12]
		verifyingDigits := cnpj[12:14]
		firstSum := sumDigitsCnpj(firstTwelve, 5)	
		firstVD := string(verifyingDigits[0])		
		firstVDNumb, _:= strconv.Atoi(firstVD)
		rest := modulusCnpj(firstSum)
		if rest == firstVDNumb {
			secondVD := string(verifyingDigits[1])
			secondVDNumb, _:= strconv.Atoi(secondVD)
			secondSum := sumDigitsCnpj(firstTwelve + firstVD, 6)
			rest := modulusCnpj(secondSum)
			if rest == secondVDNumb{
				return true
			}
		}
	}
	return false
}

func sumDigitsCnpj(firstTwelve string, startNumb int) int {
	sum := 0
	count := 0
	for i := startNumb; i > 1; i-- {
		numb, _:= strconv.Atoi(string(firstTwelve[count]))
		sum += numb * i
		count++
	}
	for i := 9; i > 1; i-- {
		numb, _:= strconv.Atoi(string(firstTwelve[count]))
		sum += numb * i
		count++
	}
	return sum
}

func modulusCnpj(sum int) int {
	rest := (sum % 11); if rest < 2{
		return 0		
	}
	rest = 11 - rest 
	return rest
}