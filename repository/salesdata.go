package repository

import (
	"desafioNeoWay/entity"
	"desafioNeoWay/util"
	"fmt"
)

//SalesInterface defines the methods must be implemented by injected layer
type SalesInterface interface{
	Insert(s *entity.SalesData) error
}

//Insert a new sales data
func (c *DataBase) Insert(s *entity.SalesData) error{
	db := c.connection
	t := db.Exec(getCommand(s))
	return t.Error
	//return db.Create(&s).Error
}

func getCommand(s *entity.SalesData) string {
	command := "insert into sales_data (cpf, cpf_valido, private, incompleto, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, cnpj_loja_mais_frequente_valido, loja_da_ultima_compra, cnpj_loja_da_ultima_compra_valido) values("
	command += fmt.Sprint(
						util.StringLikeStringSQL(s.Cpf),",",
						util.BoolLikeStringSQL(s.CpfValido),",",
						util.BoolLikeStringSQL(s.Private),",",
						util.BoolLikeStringSQL(s.Incompleto),",",
						util.DateLikeStringSQL(s.DataDaUltimaCompra),",",
						util.Float32LikeStringSQL(s.TicketMedio),",",
						util.Float32LikeStringSQL(s.TicketDaUltimaCompra),",",
						util.StringLikeStringSQL(s.LojaMaisFrequente),",",
						util.BoolLikeStringSQL(s.CnpjLojaMaisFrequenteValido),",",
						util.StringLikeStringSQL(s.LojaDaUltimaCompra),",",
						util.BoolLikeStringSQL(s.CnpjLojaDaUltimaCompraValido),
						)
	command += ")"
	return 	command
}
