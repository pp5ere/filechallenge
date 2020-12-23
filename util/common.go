package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//GetOnlyNumbers retuns a string valid number
func GetOnlyNumbers(text string) string {
	onlyNumbers := regexp.MustCompile(`(\d)+`)
	cpfArr := onlyNumbers.FindAllString(text, len(text))
	return arrayToString(cpfArr)
}

//GetBoolean parse string to boolean
func GetBoolean(txt string) (bool, error) {
	if strings.Trim(txt, " ") != "" {
		numb, err := strconv.Atoi(txt); if err != nil{
			return false, err
		}
		switch numb {
		case 0:
			return false, nil
		case 1:
			return true, nil			
		}

	}
	return false, errors.New("Cannot to convert empty string to boolean")
}

//GetDate parse string to time.Time
func GetDate(txt string) (time.Time){
	if strings.Trim(txt, " ") != ""{
		date, err := time.Parse("2006-01-02", txt);if err != nil{			
			return time.Time{}
		}
		return date 
	}
	return time.Time{}	
}
//GetFloat returns a valid float32 from string formated like: '1,2', '1', '1.231,1', ',2' to '1.2', '1', '1231.1', '.2'
func GetFloat(txt string) (float32) {
	fNumb := regexp.MustCompile(`^(\d+(\.\d{0,3})*\,{0,1}\d*|\,{0,1}\d*)$`)
	if fNumb.MatchString(txt){
		txt = strings.ReplaceAll(txt,".","")
		txt = strings.ReplaceAll(txt,",",".")
		numb, err := strconv.ParseFloat(txt, 64); if err != nil {
			return 0
		}
		return float32(numb)
	}	
	return 0
}