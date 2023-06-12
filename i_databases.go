package ooo

import (
	"database/sql"
	sqlLib "database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

var OSVARENV map[string]string

func RepVarEnv(repo string, src string) string {
	if OSVARENV == nil {
		OSVARENV = make(map[string]string, 10)
		yfile, err := ioutil.ReadFile(repo + "/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		data := make(map[string]varenvs)
		err2 := yaml.Unmarshal(yfile, &data)
		if err2 != nil {
			log.Fatal(err2)
		}
		//	fmt.Println(data)
		for _, v := range data {
			for _, gg := range v.Keys {
				OSVARENV[gg] = OSENV(gg)
			}

		}
	}
	src0 := src
	if Rep(src0, "$", "") != src {
		for k, v := range OSVARENV {
			if len(k) > 1 {
				src0 = Rep(src0, "$"+k, v)
			}
		}
	}

	return src0
}

func SetBases(repo string) *map[string]database {
	dbs := make(map[string]database, 20)
	yfile, err := ioutil.ReadFile(repo + "/" + "config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]varenvs)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, v := range data {
		for _, gg := range v.Databases {
			alias := gg.Name
			dd := database{}
			dd.Name = gg.Name
			dd.Typ = gg.Typ
			dd.Url = RepVarEnv(repo, gg.Url)
			dbs[alias] = dd
		}
	}
	return &dbs
}

func Rsql(d database, sql string) int {
	return rsql(sql, d.Url, d.Typ)

}

func RsqlFi(d database, fileName string) int {
	sql := RFi(fileName)

	return rsql(sql, d.Url, d.Typ)
}

func rsql(sql string, databaseConnectiontring string, dbtyp string) int {
	db := openDb(dbtyp, databaseConnectiontring)
	var rows *sqlLib.Rows
	var er error
	ue := 0
	if db != nil {
		defer db.Close()
		rows, er = db.Query(sql)
		if er != nil {
			ue = -1
			fmt.Println("rSQL->", er.Error())
			fmt.Println(sql)
			ue = -1
		} else {
			defer rows.Close()

		}
	} else {
		fmt.Println("error connecting database")
		ue = -1
	}

	return ue
}

func Dsql(d *database, Sql string) (*Datset, string) {
	baseType := d.Typ
	conn := d.Url
	var DS Datset
	DS.ColsCount = 0
	stcountRows := `
	select count(*) from (
		` + Rep(Sql, ";", "") + `
	                     ) a123456789z;`
	// to avoid accidental escape
	var countRows int
	db := openDb(baseType, conn)
	er1 := db.QueryRow(stcountRows).Scan(&countRows)
	if er1 != nil {
		return nil, er1.Error()
	}
	rows, e2 := db.Query(Sql)
	if e2 != nil {
		fmt.Println(e2.Error())
		return nil, e2.Error()
	}
	defer rows.Close()
	defer db.Close()

	cols, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	DS.ColsCount = len(cols)

	var ColFi Row
	var TypFi Row
	var GTypFi Row
	var PTypFi Row

	ColFi.FI = make([]string, DS.ColsCount)
	TypFi.FI = make([]string, DS.ColsCount)
	GTypFi.FI = make([]string, DS.ColsCount)
	PTypFi.FI = make([]string, DS.ColsCount)

	DS.Cols = &ColFi
	DS.Typs = &TypFi
	DS.Gtyps = &GTypFi
	DS.Ptyps = &PTypFi

	var j int
	for j = 0; j < DS.ColsCount; j++ {
		DS.Cols.FI[j] = cols[j]
		DS.Typs.FI[j] = colTypes[j].DatabaseTypeName()
		DS.Gtyps.FI[j] = "number"
		if colTypes[j].DatabaseTypeName() == "VARCHAR" {
			DS.Gtyps.FI[j] = "string"
		}
		if colTypes[j].DatabaseTypeName() == "TEXT" {
			DS.Gtyps.FI[j] = "string"
		}
		if colTypes[j].DatabaseTypeName() == "DATE" {
			DS.Gtyps.FI[j] = "date"
		}
		if colTypes[j].DatabaseTypeName() == "DATETIME" {
			DS.Gtyps.FI[j] = "date"
		}
		if colTypes[j].DatabaseTypeName() == "BOOLEAN" {
			DS.Gtyps.FI[j] = "boolean"
		}
	}

	//tt := []Row{}
	tt := make([]Row, countRows)
	DS.RowsCount = 0
	u := 0
	columns, _ := rows.Columns()
	colNum := len(columns)
	var values = make([]interface{}, colNum)
	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}
	for rows.Next() {
		r := Row{}
		r.FI = make([]string, len(cols))
		rows.Scan(values...)
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{}))
			var raw_type = reflect.TypeOf(raw_value)
			_ = colName
			_ = raw_type
			if raw_value == nil {
				r.FI[i] = ""
			} else {
				switch raw_type.String() {
				case "int64":
					r.FI[i] = strconv.FormatInt(raw_value.(int64), 10)
				case "int32":
					r.FI[i] = strconv.FormatInt(raw_value.(int64), 10)
				case "[]uint8":
					r.FI[i] = string(raw_value.([]uint8))
				case "bool":
					r.FI[i] = strconv.FormatBool(raw_value.(bool))
				case "float32":
					r.FI[i] = strconv.FormatFloat(raw_value.(float64), 'G', -1, 32)
				case "float64":
					r.FI[i] = strconv.FormatFloat(raw_value.(float64), 'G', -1, 64)
				case "time.Time":
					r.FI[i] = raw_value.(time.Time).Format("2006-01-02 15:04:05")
				default:
					r.FI[i] = raw_value.(string)
				}
			}
		}
		tt[u] = r
		u++
		if u >= countRows {
			break
		}
	}
	DS.Rows = &tt
	DS.RowsCount = countRows
	DS.SetCols()

	rows.Close()
	return &DS, ""
}

func (dt *Datset) SetCols() {
	if dt.ColsCount > 0 {
		dt.C = make(map[string]int, dt.ColsCount)
		for u := 0; u < dt.ColsCount; u++ {
			dt.C[dt.Cols.FI[u]] = u
		}
	}
}

func openDb(baseType string, databaseConnectiontring string) *sql.DB {

	dbb, errb := sql.Open(baseType, databaseConnectiontring)
	if errb != nil {
		return nil
	}
	return dbb
}
