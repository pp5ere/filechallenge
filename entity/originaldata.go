package entity

//OriginalData contains the original datas import from file
type OriginalData struct{
	ID						uint   `gorm:"primaryKey; autoIncrement:true"`
	Cpf 					string
	Private 				string
	Incompleto 				string
	DataDaUltimaCompra 		string
	TicketMedio 			string
	TicketDaUltimaCompra 	string
	LojaMaisFrequente 		string
	LojaDaUltimaCompra 		string
}