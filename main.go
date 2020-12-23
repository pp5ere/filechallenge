package main

import (
	"bufio"
	"desafioNeoWay/controller"
	"desafioNeoWay/entity"
	"desafioNeoWay/repository"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var scanner *bufio.Scanner

func receiveFile(w http.ResponseWriter, r *http.Request) { 
	r.ParseMultipartForm(32 << 20)   
	switch r.Method {
		case "POST":
			txtFile, err := repository.SaveFile(w, r);if err != nil{
				log.Fatal(err)
			}
			w.Write([]byte(fmt.Sprintln("Iniciando processamento do Arquivo...")))
			go Start(txtFile)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
   	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("template/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
   

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/upload", receiveFile)
    http.ListenAndServe(":8080", nil)
}

//Start and proccess a received file
func Start(txtFile *repository.TxtFile) {
	log.Println("start app")
	file, err := txtFile.ReadFile();if  err != nil{
		log.Fatal(err.Error())
	}
	if err := ProccessFile(file); err != nil{
		log.Fatal(err.Error())
	}
	if err := txtFile.RemoveFile(); err != nil{
		log.Println(err.Error())
	}
	log.Println("stop app")
}

//ProccessFile scan whole file transform and insert into database
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
		line := repository.RemoveSpaces(scanner.Text())
		if err := repository.ParseToStruct(line, &salesData); err != nil{
			return err
		}
		if err := ctrl.Sales.Insert(&salesData); err != nil{
			return err
		}
    }
	return nil
}