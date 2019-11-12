package main

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"log"
	"os"
	"regexp"
)

//default database connection
var (
	host     = "hh-pgsql-public.ebi.ac.uk"
	port     = "5432"
	user     = "reader"
	password = "NWDMCE5xdipIjRrp"
	dbname   = "pfmegrnargs"
	dbtype   = "postgres"
)

var (
	result  string
	connStr string
	stmt    string
)

func main() {
	if len(os.Args) > 1 {
		user = regexp.MustCompile(`(\w+):`).FindStringSubmatch(os.Args[1])[1]
		password = regexp.MustCompile(`:(.*)@`).FindStringSubmatch(os.Args[1])[1]
		host = regexp.MustCompile(`@(.*):`).FindStringSubmatch(os.Args[1])[1]
		port = regexp.MustCompile(`:(\d+)(/)`).FindStringSubmatch(os.Args[1])[1]
		dbname = regexp.MustCompile(`/(\w+)`).FindStringSubmatch(os.Args[1])[1]
	}
	if len(os.Args) == 3 {
		dbtype = os.Args[2]
	}
	switch dbtype {
	case "postgres":
		connStr = "postgres" + "://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
		stmt = "select version()"
	case "mssql":
		connStr = "sqlserver" + "://" + user + ":" + password + "@" + host + ":" + port + "?" + "database=" + dbname
		stmt = "select @@version"
	case "oracle":
		connStr = ""
		stmt = "select * from v$version"
	}
	log.Print(connStr)
	db, err := sql.Open(dbtype, connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow(stmt).Scan(&result)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success connected")
	log.Println(result)

}
