package db

import (
	"strings"

	"github.com/go-ole/go-ole"
)

// Select - berilgan sql so'rov asosida bazadan ma'lumotlarni qidiradi
func Select(path, sql string, args ...interface{}) ([]map[string]interface{}, error) {
	db, err := connect(path)
	if err != nil {
		return nil, err
	}
	defer ole.CoUninitialize()
	defer db.Close()
	rows, err := db.Query(sql, args...)
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
