package models

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Data        interface{} `json:"data"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"Next_page"`
}

type ParamPag struct {
	Limit   int
	OrderBy []string
	ShowSQL bool
}

var pag = ParamPag{
	Limit:   30,
	OrderBy: []string{"id desc"},
	ShowSQL: false,
}

var db *gorm.DB //database

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error en Archivo Env")
	}

	dbDebug, _ := strconv.ParseBool(os.Getenv("db_debug"))
	pag.ShowSQL = dbDebug
	dbUser := os.Getenv(("db_user"))
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	charset := os.Getenv("charset")
	dbUri := ""
	switch dbType {
	case "mysql":
		dbUri = fmt.Sprintf("%s:%s@%s:%s/%s?charset=%s&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName, charset)
		break
	case "postgres":
		dbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
		break
	case "mssql":
		dbUri = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)
		break

	case "sqlite":
		dbUri = fmt.Sprintf(dbName)
		break
	}

	conn, err := gorm.Open(dbType, dbUri)

	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate()

}

func GetDB() *gorm.DB {
	if pag.ShowSQL {
		db.Debug()
	}
	return db
}

func Paging(result interface{}, page, items int) *Paginator {
	if items == 0 {
		dbPagination, _ := strconv.Atoi(os.Getenv("db_pagination"))
		if dbPagination > 0 {
			pag.Limit = dbPagination
		}
	} else {
		pag.Limit = items
	}
	if page < 1 {
		page = 1
	}

	if pag.ShowSQL {
		db.Debug()
	}

	if len(pag.OrderBy) > 0 {
		for _, o := range pag.OrderBy {
			db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecord(result, done, &count)

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * pag.Limit
	}
	db.Limit(pag.Limit).Offset(offset).Find(result)
	<-done
	paginator.TotalRecord = count
	paginator.Data = result
	paginator.Page = page
	paginator.Limit = pag.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(pag.Limit)))

	if page > 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}
	if page == paginator.TotalPage {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}
	return &paginator

}

func countRecord(anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
