package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/go-ole/go-ole"
	_ "github.com/mattn/go-adodb"
	"gopkg.in/yaml.v3"
)

var DB *sql.DB

func init() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	conf := map[string]string{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	dsn := fmt.Sprintf("Provider=Microsoft.Jet.OLEDB.4.0;Data Source=%s;Extended Properties=Paradox 5.x;", conf["db_path"])
	db, err := sql.Open("adodb", dsn)
	if err != nil {
		panic(err)
	}
	DB = db
}
