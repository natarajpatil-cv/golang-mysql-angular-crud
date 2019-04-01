package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Conf struct {
	Database Database `json:"database"`
}

type Database struct {
	Drive  string   `json:"drive"`
	User   string   `json:"user"`
	Pass   string   `json:"pass"`
	Host   string   `json:"host"`
	Port   string   `json:"port"`
	Scheme []string `json:"scheme"`
}

var db *sql.DB
var err error

func connectDatabse() {
	/* db, err = sql.Open("mysql", "root:root@/test")
	fmt.Println("Database connected.") */

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}

	dbconf := conf.Database
	db, err = sql.Open(dbconf.Drive, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbconf.User, dbconf.Pass, dbconf.Host, dbconf.Port, dbconf.Scheme[0]))
	if err != nil {
		log.Panic("open db error", err.Error())
	}

	log.Println("successfully connected to mysql")
}
