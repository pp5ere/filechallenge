package repository

import (
	"desafioNeoWay/entity"
	"desafioNeoWay/util"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//TxtFile store the file path
type TxtFile struct{
	FilePath string
}

//ReadFile return a file
func (txt *TxtFile) ReadFile() (*os.File, error) {
	txtFile, err := os.Open(txt.FilePath); if err != nil{
		return nil, err
	}
	return txtFile, nil
}

//RemoveFile delete the file
func (txt *TxtFile) RemoveFile() error {	
	if err := os.Remove(txt.FilePath); err != nil{
		return err
	}
	return nil
}

//SaveFile store a new file
func SaveFile(w http.ResponseWriter, r *http.Request)(*TxtFile ,error) {
	file, fileHeader, err := r.FormFile("file"); if err != nil{
		return nil,err
	}
	dirPath := "temp"
	if !util.DirExist(dirPath){
		if err := os.Mkdir("temp", 0755); err != nil{
			return nil, err
		}
	}
	out, pathError := ioutil.TempFile(dirPath, fileHeader.Filename); if pathError != nil {
		return nil, errors.New(fmt.Sprintln("Error Creating a file for writing", pathError))
	}
	defer out.Close()

	_, copyError := io.Copy(out, file); if copyError != nil {
		return nil, errors.New(fmt.Sprintln("Error copying", copyError))
	}
	txtFile := TxtFile{FilePath: out.Name()}
	defer file.Close()
	w.Write([]byte(fmt.Sprintln("Arquivo", fileHeader.Filename , "recebido com sucesso")))

	return &txtFile, nil
}
//RemoveSpaces changes whole spaces character to ';' character
func RemoveSpaces(line string) string{
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(line, ";")
}
//ParseToStruct parse the string line to struct
func ParseToStruct(line string, txt *entity.SalesData) error {
	arr := strings.Split(line, ";")
	if len(arr) == 8 {
		cpf 		:= util.GetOnlyNumbers(arr[0])
		isValidCpf 	:= util.CpfIsValid(cpf)		
		prv, err := util.GetBoolean(arr[1]); if err != nil{
			return err
		}
		inc, err := util.GetBoolean(arr[2]); if err != nil{
			return err
		}
		date := util.GetDate(arr[3])
		tktMedio := util.GetFloat(arr[4])
		tktUlt := util.GetFloat(arr[5])
		txt.Cpf 						= cpf
		txt.CpfValido					= isValidCpf
		txt.Private 					= prv
		txt.Incompleto 					= inc
		txt.DataDaUltimaCompra 			= date
		txt.TicketMedio 				= tktMedio
		txt.TicketDaUltimaCompra		= tktUlt
		txt.LojaMaisFrequente 			= util.GetOnlyNumbers(arr[6])
		txt.CnpjLojaMaisFrequenteValido = util.CnpjIsValid(txt.LojaMaisFrequente)
		txt.LojaDaUltimaCompra 			= util.GetOnlyNumbers(arr[7])
		txt.CnpjLojaDaUltimaCompraValido= util.CnpjIsValid(txt.LojaDaUltimaCompra)
	}else{
		return errors.New(fmt.Sprintln("Cannot parse to struct because columns number should be", len(arr), arr))
	}	
	return nil
}