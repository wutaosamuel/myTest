package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"io/ioutil"

	_ "github.com/lib/pq"
)

// PQTest
type PQDb struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewPQTest
func NewPQTest() *PQDb {
	return &PQDb{
		Host:     "172.16.8.221",
		Port:     5432,
		User:     "pi",
		Password: "wutaowutao",
		DBName:   "test",
	}
}

// Source
func (pq *PQDb) Source() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pq.Host, pq.Port, pq.User, pq.Password, pq.DBName)
}

// Insert
func Insert(key, value string) string {
	return ""
}

func main() {
	pqInfo := NewPQTest()
	pq_info := pqInfo.Source()
	db, err := sql.Open("postgres", pq_info)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	defer db.Close()
	// db.Ping()
	// create scheme/table if not exist
	createST(db)

	// Insert
	o := NewObject()
	o.Name = "Object"
	o.ID = 1
	o.Value = 12.123456789
	insert(db, o)

	// query
	query(db)

	// update
	o.Value = 8.9876654321
	update(db, o)

	// delete
	delete(db)
	// allquery
	fmt.Println("all query now")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	o.Name = "obj"
	key, value := o.InsertTo()
	insert := "INSERT INTO test_schema.test_object(" + key + ") VALUES (" + value + ");"
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	o.Name = "obj_d"
	o.ID = 2
	key, value = o.InsertTo()
	insert = "INSERT INTO test_schema.test_object(" + key + ") VALUES (" + value + ");"
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
	//allquery(db)
	res := OQuery(db, &Object{})
	for _, r := range res {
		fmt.Println(r)
	}
}

// create scheme/table if not exist
func createST(db *sql.DB) {
	// create scheme/table if not exist
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("if create not exists, create table")
	f, err := ioutil.ReadFile("./test_object_if_exists.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	table := string(f)
	if _, err := tx.Exec(table); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
}

// insert
func insert(db *sql.DB, obj *Object) {
	// Insert
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start Insert")
	key, value := obj.InsertTo()
	insert := "INSERT INTO test_schema.test_object(" + key + ") VALUES (" + value + ");"
	fmt.Println(insert)
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
}

// query
func query(db *sql.DB) {
	// query
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to query")
	q := NewObject()
	query := "SELECT * FROM test_schema.test_object WHERE name = 'Object';"
	row := tx.QueryRow(query)
	if err := row.Scan(q.SelectTo()...); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
	fmt.Println(q)
}

// update
func update(db *sql.DB, obj *Object) {
	// update
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to update")
	update := "UPDATE test_schema.test_object SET " + obj.UpdateTo() + "WHERE name = 'Object';"
	if _, err := tx.Exec(update); err != nil {
		fmt.Println(err)
		return
	}
	// read whether exists
	q := NewObject()
	query := "SELECT * FROM test_schema.test_object WHERE name = 'Object';"
	row := tx.QueryRow(query)
	if err := row.Scan(q.SelectTo()...); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
	fmt.Println(q)
}

// delete
func delete(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to delete")
	delete := "DELETE FROM test_schema.test_object WHERE name = 'Object'"
	if _, err := tx.Exec(delete); err != nil {
		fmt.Println(err)
		return
	}
}

// allquery
func Allquery(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	allquery := "SELECT * FROM test_schema.test_object;"
	rows, err := tx.Query(allquery)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		a := NewObject()
		if err := rows.Scan(a.SelectTo()...); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(a)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
}

// OQuery
func OQuery(db *sql.DB, model interface{}) []interface{} {
	result := make([]interface{}, 0)
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return result
	}
	m := reflect.TypeOf(model)
	if m.Kind() == reflect.Ptr {
		m = m.Elem()
	}
	if m.Kind() != reflect.Struct {
		return nil
	}

	allquery := "SELECT * FROM test_schema.test_object;"
	rows, err := tx.Query(allquery)
	if err != nil {
		fmt.Println(err)
		return result
	}
	for rows.Next() {
		s := reflect.New(m).Elem()
		num := s.NumField()
		column := make([]interface{}, num)
		for i := 0; i < num; i++ {
			field := s.Field(i)
			column[i] = field.Addr().Interface()
		}
		if err := rows.Scan(column...); err != nil {
			fmt.Println(err)
			return result
		}
		result = append(result, s.Addr().Interface())
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return []interface{}{}
	}
	tx.Commit()
	return result
}