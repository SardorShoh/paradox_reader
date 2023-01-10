package db

import (
	"database/sql"
	"fmt"

	"github.com/go-ole/go-ole"
	_ "github.com/mattn/go-adodb"
)

func connect(path string) (*sql.DB, error) {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	// ole.CoInitialize(0)
	dsn := fmt.Sprintf("Provider=Microsoft.Jet.OLEDB.4.0;Data Source=%s;Extended Properties=Paradox 5.x;", path)
	db, err := sql.Open("adodb", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
