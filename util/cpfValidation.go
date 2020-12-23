package util

import(	
	"strconv"
)

//CpfIsValid return a boolean valid cpf
func CpfIsValid(cpf string) bool {
	cpf = GetOnlyNumbers(cpf)
	if len(cpf) == 11 {
		if !allDigitsSame(cpf){
			firstNine := cpf[0:9]
			verifyingDigits := cpf[9:11]
			firstSum := sumDigitsCpf(firstNine, 10)			
			firstVD := string(verifyingDigits[0])		
			firstVDNumb, _:= strconv.Atoi(firstVD)
			rest := modulusCpf(firstSum)
			if rest == firstVDNumb {
				secondVD := string(verifyingDigits[1])
				secondVDNumb, _:= strconv.Atoi(secondVD)
				secondSum := sumDigitsCpf(firstNine + firstVD, 11)
				rest := modulusCpf(secondSum)
				if rest == secondVDNumb{
					return true
				}
			}	
		}
	}
	return false
}

func sumDigitsCpf(firstNine string, digNumb int) int {
	sum := 0
	for i := digNumb; i > 1; i-- {
		numb, _:= strconv.Atoi(string(firstNine[digNumb-i]))
		sum += numb * i
	}
	return sum
}

func modulusCpf(sum int) int {
	rest := (sum * 10) % 11; if rest > 9{
		rest = 0		
	}
	return rest
}

func allDigitsSame(cpf string) bool {
	sameDigit := 0
	for i := 0; i < len(cpf); i++ {
		if cpf[0] == cpf[i]{
			sameDigit++
		}
	}	
	return sameDigit > 10
}

func arrayToString(arr []string) string {
	var text string = ""
	if len(arr) > 1{
		for _, v := range arr {
			text += v
		}	
	}else {
		if len(arr) > 0{
			return arr[0]
		}
		return ""		
	}
	return text
}