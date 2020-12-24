package main

import (
	"bufio"
	"filechallenge/controller"
	"filechallenge/entity"
	"filechallenge/repository"
	"filechallenge/util"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var scanner *bufio.Scanner

func receiveFile(w http.ResponseWriter, r *http.Request) { 
	r.ParseMultipartForm(32 << 20)   
	switch r.Method {
		case "POST":
			txtFile, err := repository.SaveFile(w, r);if err != nil{
				log.Fatal(err)
			}
			w.Write([]byte(fmt.Sprintln(time.Now(),"Iniciando processamento do Arquivo...")))
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
	setLogOutPut()	
	log.Println("start application")
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/upload", receiveFile)
	http.ListenAndServe(":8080", nil)
}

func setLogOutPut(){
	dirPath := "log"
	if err := util.CreateDir(dirPath); err != nil {		
		log.Println(err.Error())
	}
	f, err := os.OpenFile(dirPath + "/logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); if err != nil {
        log.Println("error opening file:", err)
    }
    wrt := io.MultiWriter(os.Stdout, f)

	log.SetOutput(wrt)
}

//Start and proccess a received file
func Start(txtFile *repository.TxtFile) {
	log.Println("start to proccess file")
	file, err := txtFile.ReadFile();if  err != nil{
		log.Println(err.Error())
	}
	if err := ProcessFile(file); err != nil{
		log.Println(err.Error())
	}
	if err := txtFile.RemoveFile(); err != nil{
		log.Println(err.Error())
	}
	log.Println("stop to proccess file")
}

//ProcessFile scan whole file transform and insert into database
func ProcessFile(file *os.File) error {	
	repo, err := repository.New(os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE")); if err != nil {
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