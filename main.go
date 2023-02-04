package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"paradox_reader/db"
	"strconv"
	"strings"
)

/*
Dasturga quyidagicha so'rov yuborildi
paradox.exe -path "\\HP\db" -query="select * from rs_nakl where Data=? and ID=?" -args="11.05.2020|s;31|i"
*/
func main() {
	path := flag.String("path", "", "Paradox bazasining joylashgan manzili")
	query := flag.String("query", "", "Paradox bazasiga yuboriladigan sql so'rov")
	args := flag.String("args", "", "Paradox bazasiga yuboriladigan so'rovning argumentlari")
	flag.Parse()

	if *path == "" {
		fmt.Println("Paradox base path not found")
		os.Exit(1)
	}

	var arguments []interface{}
	if *args != "" {
		params := strings.Split(*args, ";")
		if len(params) > 0 {
			for _, arg := range params {
				spl := strings.Split(arg, "|")
				if len(spl) == 2 {
					switch spl[1] {
					case "s":
						arguments = append(arguments, db.Encrypt(spl[0]))
					case "i":
						num, err := strconv.Atoi(spl[0])
						if err == nil {
							arguments = append(arguments, num)
						}
					case "f":
						fl, err := strconv.ParseFloat(spl[0], 64)
						if err == nil {
							arguments = append(arguments, fl)
						}
					case "b":
						bl, err := strconv.ParseBool(arg)
						if err == nil {
							arguments = append(arguments, bl)
						}
					}
				} else {
					arguments = append(arguments, db.Encrypt(spl[0]))
				}
			}
		}
	}
	data, err := db.Select(*path, *query, arguments...)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}
