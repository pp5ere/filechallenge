package main

import (
	"bufio"
	"desafioNeoWay/controller"
	"desafioNeoWay/entity"
	"desafioNeoWay/repository"
	"desafioNeoWay/util"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var scanner *bufio.Scanner

func main() {	
	log.Println("start app")
	txtFile := TxtFile{FilePath: "base_teste.txt"}
	file, err := txtFile.ReadFile();if  err != nil{
		log.Fatal(err.Error())
	}
	if err := ProccessFile(file); err != nil{
		log.Fatal(err.Error())
	}
	log.Println("stop app")
}

func ProccessFile(file *os.File) error {
	repo, err := repository.New("127.0.0.1", "5432", "postgres", "1234", "store"); if err != nil {
		return err
	}
	ctrl := controller.New(repo)
	salesData := entity.SalesData{}
	scanner = bufio.NewScanner(file)
	defer file.Close()
	scanner.Scan()
    for scanner.Scan() {
		line := RemoveSpaces(scanner.Text())
		if err := ParseToStruct(line, &salesData); err != nil{
			return err
		}
		if err := ctrl.Sales.Insert(&salesData); err != nil{
			return err
		}
    }
	return nil
}

type TxtFile struct{
	FilePath string
}

func (txt *TxtFile) ReadFile() (*os.File, error) {
	txtFile, err := os.Open(txt.FilePath); if err != nil{
		return nil, err
	}
	return txtFile, nil
}

func RemoveSpaces(line string) string{
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(line, ";")
}

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
		return errors.New(fmt.Sprintln("Cannot parse to struct because columns number is wrong", arr))
	}	
	return nil
}