package repository

import (
	"desafioNeoWay/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DataBase defines the connection with database
type DataBase struct{
	host string
	port string
	user string
	pass string
	dbname string
	connection *gorm.DB
}

//New creates a new Repository layer sharing the mysql connection
func New(host string, port string, user string, pass string, dbname string) (*DataBase, error) {
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=disable" + " TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil, err
	}	
	db.AutoMigrate(&entity.SalesData{})
	return &DataBase{
		host: host,
		port: port,
		user: user,
		pass: pass,
		dbname: dbname,
		connection: db,
	}, err
}