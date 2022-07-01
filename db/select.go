package db

import (
	"strings"

	"github.com/go-ole/go-ole"
)

//Select - berilgan sql so'rov asosida bazadan ma'lumotlarni qidiradi
func Select(sql string, args ...interface{}) ([]map[string]interface{}, error) {
	defer ole.CoUninitialize()
	defer DB.Close()
	rows, err := DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	colLength := len(colNames)
	cols := make([]interface{}, colLength)
	colPtrs := make([]interface{}, colLength)
	for i := 0; i < colLength; i++ {
		colPtrs[i] = &cols[i]
	}
	retMap := []map[string]interface{}{}
	for rows.Next() {
		retM := make(map[string]interface{})
		err = rows.Scan(colPtrs...)
		if err != nil {
			return nil, err
		}
		for i, col := range cols {
			if colTypes[i].DatabaseTypeName() == "ADVARWCHAR" {
				if col == nil {
					retM[strings.ToLower(colNames[i])] = col
				} else {
					retM[strings.ToLower(colNames[i])] = decrypt(col.(string))
				}
			} else {
				retM[strings.ToLower(colNames[i])] = col
			}
		}
		retMap = append(retMap, retM)
	}
	return retMap, nil
}

//Exec - berilgan sql asosida bazaga ma'lumotlarni yozadi
func Exec(sql string, args ...interface{}) error {
	defer ole.CoUninitialize()
	defer DB.Close()
	_, err := DB.Exec(sql, args...)
	return err
}
