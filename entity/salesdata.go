package entity

import "time"

//SalesData contains data after transformations
type SalesData struct{
	ID								uint   `gorm:"primaryKey; autoIncrement:true"`
	Cpf 							string
	CpfValido						bool
	Private 						bool
	Incompleto 						bool
	DataDaUltimaCompra 				time.Time
	TicketMedio 					float32
	TicketDaUltimaCompra 			float32
	LojaMaisFrequente 				string
	CnpjLojaMaisFrequenteValido 	bool
	LojaDaUltimaCompra 				string
	CnpjLojaDaUltimaCompraValido 	bool
}