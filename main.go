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
	"sync"
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
			w.Write([]byte(fmt.Sprintln(time.Now().Format(time.RFC3339),"Iniciando processamento do Arquivo...")))
			Start(txtFile)
			w.Write([]byte(fmt.Sprintln(time.Now().Format(time.RFC3339),"Processamento do Arquivo finalizado")))
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
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
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

	defer file.Close()

	txtFile.ProcessFile(file)
	
	if err := InsertIntoDb(txtFile); err != nil{
		log.Println(err.Error())
	}

	if err := txtFile.RemoveFile(); err != nil{
		log.Println(err.Error())
	}

	log.Println("stop to proccess file")
}

//InsertIntoDb insert txtfile into database
func InsertIntoDb(txtFile *repository.TxtFile) error {	
	linesByGoRotine := linesToProcessForEachGoRotine(txtFile.Lines)
	wg := &sync.WaitGroup{}
	wg.Add(len(linesByGoRotine))
	for i := 0; i < len(linesByGoRotine); i++ {
		go Persist(wg, txtFile, linesByGoRotine[i].Start, linesByGoRotine[i].End)
	}
	wg.Wait()
	
	return nil
}

func Persist(wg *sync.WaitGroup, txtFile *repository.TxtFile, pStart int64, pEnd int64)  {
	repo, err := repository.New(os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE")); if err != nil {
		log.Println(err.Error())
	}
	ctrl := controller.New(repo)
	salesData := entity.SalesData{}
	for i := pStart; i <= pEnd; i++ {
		line := txtFile.Lines[i]
		if err := txtFile.ParseToStruct(line, &salesData); err != nil{
			log.Println(err.Error())
		}
		if err := ctrl.Sales.Insert(&salesData); err != nil{
			log.Println(err.Error())
		}
	}
	wg.Done()
}

//LineInterval defines start/end line for each csvfile part to be processed by goroutines
type LineInterval struct {
	Start int64
	End   int64
}

func linesToProcessForEachGoRotine(txtlines [] string) [] LineInterval{
	var line LineInterval
	var lines []LineInterval
	qtdGoRotine := int64(25)
	arrayTotalLines := int64(len(txtlines))
	linesByGoRotine := int64(arrayTotalLines / qtdGoRotine)
	line.Start = 0
	if arrayTotalLines > arrayTotalLines / linesByGoRotine {
		ToProcess := int64(arrayTotalLines / linesByGoRotine)
		for i := int64(0); i <= ToProcess; i++ {
			line.End = line.Start + linesByGoRotine - 1
			if line.End > arrayTotalLines {
				line.End = arrayTotalLines - line.Start - 1
				line.End = line.End + line.Start
			}
			lines = append(lines, line)
			line.Start = line.Start + linesByGoRotine
		}
	} else {
		line.End = arrayTotalLines - 1
		lines = append(lines, line)
	}
	return lines
}