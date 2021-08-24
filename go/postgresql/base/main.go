package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

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
	// create table if not exist
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

	// Insert
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start Insert")
	o := NewObject()
	o.Name = "Object"
	o.ID = 1
	o.Value = 12.123456789
	key, value := o.InsertTo()
	insert := "INSERT INTO test_object(" + key + ") VALUES (" + value + ");"
	fmt.Println(insert)
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()

	// query
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to query")
	q := NewObject()
	query := "SELECT * FROM test_object WHERE name = 'Object';"
	row := tx.QueryRow(query)
	if err := row.Scan(q.SelectTo()...); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
	fmt.Println(q)

	// update
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to update")
	o.Value = 8.9876654321
	update := "UPDATE test_object SET " + o.UpdateTo() + "WHERE name = 'Object';"
	if _, err := tx.Exec(update); err != nil {
		fmt.Println(err)
		return
	}
	q = NewObject()
	row = tx.QueryRow(query)
	if err := row.Scan(q.SelectTo()...); err != nil {
		fmt.Println(err)
		return
	}
	tx.Commit()
	fmt.Println(q)

	// delete
	tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println("Start to delete")
	delete := "DELETE FROM test_object WHERE name = 'Object'"
	if _, err := tx.Exec(delete); err != nil {
		fmt.Println(err)
		return
	}
	// allquery
	o.Name = "obj"
	key, value = o.InsertTo()
	insert = "INSERT INTO test_object(" + key + ") VALUES (" + value + ");"
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	o.Name = "obj_d"
	o.ID = 2
	key, value = o.InsertTo()
	insert = "INSERT INTO test_object(" + key + ") VALUES (" + value + ");"
	if _, err := tx.Exec(insert); err != nil {
		fmt.Println(err)
		return
	}
	allquery := "SELECT * FROM test_object;"
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
