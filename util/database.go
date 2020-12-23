package util

import (
	"fmt"
	"strings"
	"time"
)

//DateLikeStringSQL gets Date and return SQL string
func DateLikeStringSQL(date time.Time) string{
	if !date.IsZero(){		
		return fmt.Sprint("TO_DATE('", date.Format(time.RFC3339), "','YYYY-MM-DD')")
	}
	return "NULL"
}

//StringLikeStringSQL gets string and return SQL string
func StringLikeStringSQL(txt string) string {
	if strings.Trim(txt, " ") != ""{
		return fmt.Sprint("'",txt,"'")
	}
	return "NULL"
}

//Float32LikeStringSQL gets float32 and return SQL string
func Float32LikeStringSQL(numb float32) string {
	if numb > 0 {
		return fmt.Sprint("'",numb,"'")
	}
	return "NULL"
}

//BoolLikeStringSQL gets bool and return SQL string
func BoolLikeStringSQL(value bool) string {
	if value {
		return fmt.Sprint("'","true","'")
	}
	return fmt.Sprint("'","false","'")
}