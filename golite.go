package golite

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// AppConfig stores the applicaiton configuration
var App *RootApp = &RootApp{}

type RootApp struct {
	CTL map[string]Controller
	DBM map[string]DbHandle
	TPL map[string]Template
}
type Data map[string]any

type Template struct {
	Name   string
	File   string
	Action func()
}
type Query struct {
	Sql  string
	Args []any
}
type Controller struct {
	Name     string
	Path     string
	Template string
	Data     map[string]any
	Methods  string
	writer   http.ResponseWriter
	request  *http.Request
}

func (c *Controller) Get() {
	// Route()
}

func (c *Controller) Post() {
	// Route()
}

//func (c *Controller) AddData(d map[string]any) {

// }
func (c *Controller) AddData() {
	c.Data["key"] = "that"
}
func (c *Controller) AddQuery(q Query, args ...[]any) {

}

func MakeRoute(c Controller) Controller {

	h := func(w http.ResponseWriter, req *http.Request) {
		c.writer = w
		c.request = req
		// If No Http Method Was Defined Default at GET
		if c.Methods == "" {
			c.Methods = "GET"
		}

		if ok := c.methodAllowed(req.Method); ok {
			// Check to see of the template is a file
			var t *template.Template
			if _, err := os.Stat(c.Template); err == nil {
				t, _ = template.ParseFiles(c.Template)
			} else {
				t, _ = template.New(c.Name).Parse(c.Template)
			}
			t.Execute(w, c.Data)
			_, filename, line, _ := runtime.Caller(1)
			fmt.Printf("File:%s Line:%v\n", filename, line)
		} else {
			http.Error(w, "Method Not Allowed:", http.StatusMethodNotAllowed)
		}

	}
	http.HandleFunc(c.Path, h)
	return c
}

func InitData() *map[string]any {
	m := make(map[string]any)
	return &m
}

//func (c Controller) InitData() map[string]any {
//m := make(map[string]any)
//c.Data = m
//return c.Data
//}

func (c Controller) methodAllowed(m string) bool {
	c.Methods = strings.ReplaceAll(c.Methods, " ", "|")
	fmt.Println(c.Methods)
	fmt.Println(m)
	match, _ := regexp.MatchString(c.Methods, m)
	return match
}

type DbHandle struct {
	DB *sql.DB
}

type RowsHandle struct {
	Rows *sql.Rows
}

func DbConnect(dataSourceName string) {
	//TODO Add checks for safty
	//dss := app.Config["datasources"]
	//fmt.Println("DSS:", dss)

	ds_type := "sqlite3"
	ds_dsn := "/tmp/test.db"

	//fmt.Println(dataSource["type"]
	dsh := new(DbHandle)
	switch "type" {
	case "sqlite3", "mysql", "postgres":
		dsh.DB = Connect(ds_type, ds_dsn)
	default:
		log.Fatal("Not Implemented Yet\n", "type")
	}
}

func Connect(dstype string, dsn string) *sql.DB {
	db_connection, err := sql.Open(dstype, dsn)

	dbh := db_connection

	if err != nil {
		log.Fatal("Can not connect to Database")
	}

	return dbh
}

func (h DbHandle) Query(sql string, args ...interface{}) *RowsHandle {

	rows, err := h.DB.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return &RowsHandle{Rows: rows}
}

func (rh RowsHandle) Next() bool {
	return rh.Rows.Next()
}

func (rh RowsHandle) FetchRowMap() map[string]interface{} {

	rc, _ := rh.Rows.Columns()
	hm := make(map[string]interface{}, len(rc))
	vl := make([]interface{}, len(rc))

	for i := range rc {
		vl[i] = new(interface{})
	}

	err := rh.Rows.Scan(vl...)
	if err != nil {
		log.Fatal(err)
	}

	for i, c := range rc {
		hm[c] = *vl[i].(*interface{})
	}

	return hm
}
