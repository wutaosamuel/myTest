package main

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	pq "github.com/lib/pq"
)

// TODO: (wu tao) don't use *sql.DB, sql.Tx to increase performance

// Database
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewDatabase
func NewDatabase() *Database {
	return &Database{
		Host:     "",
		Port:     5432,
		User:     "",
		Password: "",
		DBName:   "",
	}
}

// Set
func (db *Database) Set(host string, port int, user, password, db_name string) {
	db.Host = host
	db.Port = port
	db.User = user
	db.Password = password
	db.DBName = db_name
}

// Open
func (db *Database) Open(name string) (*sql.DB, error) {
	if name == "postgres" {
		db, err := sql.Open("postgres", db.Source())
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	return nil, errors.New("not support " + name)
}

// Source
// postgre
func (db *Database) Source() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.DBName)
}

// ToSchema
// postgre
func ToSchema(schema string) string {
	return "CREATE SCHEMA IF NOT EXISTS " + schema + ";"
}

// ToSchemaTable
// postgre
func ToSchemaTable(schema, table, elements string) string {
	t := schema + "." + table
	if table == "" {
		return ""
	}
	if schema == "" {
		t = table
	}
	sql := "CREATE TABLE IF NOT EXISTS " + t + `(
	` + elements + `
	);`
	return sql
}

// CreateSchema
func CreateSchema(DB *sql.DB, schema string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	schemaSQL := ToSchema(schema)
	if _, err := tx.Exec(schemaSQL); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DropSchema
func DropSchemaClean(DB *sql.DB, schema string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	dropSQL := fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE;", schema)
	if _, err := tx.Exec(dropSQL); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// CreateSchemaTable
func CreateSchemaTable(DB *sql.DB, schema, table, elements string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	tableSQL := ToSchemaTable(schema, table, elements)
	if _, err := tx.Exec(tableSQL); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DropTableDB
func DropTableDB(DB *sql.DB, schema, table string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	drop := fmt.Sprintf("DROP TABLE %s.%s;", schema, table)
	if _, err := tx.Exec(drop); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Insert
/**
 *	Insert INSERT INTO table(column1, column2 ...) VALUE ($1, $2, ...);
 *		, 'value1', value2, ...
 *  @Params
 *    - DB  			database
 *		- schema		schema, "" will ignore it
 *		- table     schema.table
 *		- new 			object{field:value}
 **/
func InsertDB(DB *sql.DB, schema, table string, new_s []interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	if err := InsertTx(tx, schema, table, new_s); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// InsertDBs
/**
 *	Insert INSERT INTO table(column1, column2 ...) VALUE ($1, $2, ...);
 *		, 'value1', value2, ...
 *  @Params
 *    - DB  			database
 *		- schema		schema, "" will ignore it
 *		- values    [Table][]ValueInTable
 **/
func InsertDBs(DB *sql.DB, schema string, values map[string][]interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	for table, new_s := range values {
		if err := InsertTx(tx, schema, table, new_s); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// insert
func InsertTx(tx *sql.Tx, schema, table string, new_s []interface{}) error {
	if schema != "" {
		table = fmt.Sprintf("%s.%s", schema, table)
	}

	for _, new := range new_s {
		keys, values, args, err := ToInsertSQL(new)
		if err != nil {
			return err
		}

		insert := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, keys, values)
		if _, err := tx.Exec(insert, args...); err != nil {
			return err
		}
	}
	return nil
}

// Query
/**
 *	Query SELECT * FROM table condition;
 *	NOTE: select * can fit to model field
 *  @Params
 *    - DB  				database
 *		- schema			schema, "" will ignore it
 *		- table     	schema.table
 *		- condition		e.g. WHERE name = 'name' AND value >= 10; "" will get all values in table
 *		- new 				object{field:value}
 *	@Returns
 *		- []object{field:value}
 *		- error
 **/
func QueryDB(DB *sql.DB, schema, table, condition string, model interface{}) ([]interface{}, error) {
	results := make([]interface{}, 0)
	tx, err := DB.Begin()
	if err != nil {
		return results, err
	}

	if results, err = QueryTx(tx, schema, table, condition, model); err != nil {
		tx.Rollback()
		return results, err
	}
	tx.Commit()
	return results, nil
}

// QueryTx
func QueryTx(tx *sql.Tx, schema, table, condition string, model interface{}) ([]interface{}, error) {
	results := make([]interface{}, 0)
	if schema != "" {
		table = fmt.Sprintf("%s.%s", schema, table)
	}

	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return results, errors.New("query model type is not struct")
	}

	query := fmt.Sprintf("SELECT * FROM %s %s;", table, condition)
	if condition == "" {
		query = fmt.Sprintf("SELECT * FROM %s;", table)
	}
	rows, err := tx.Query(query)
	if err != nil {
		return results, err
	}
	for rows.Next() {
		object := reflect.New(modelType).Elem()
		numField := object.NumField()
		column := make([]interface{}, numField)
		for i := 0; i < numField; i++ {
			field := object.Field(i)
			if arr := ToSQLArray(field); arr != nil {
				column[i] = arr
				continue
			}
			column[i] = field.Addr().Interface()
		}
		if err := rows.Scan(column...); err != nil {
			return results, err
		}
		results = append(results, object.Addr().Interface())
	}
	if err := rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}

// Update
/**
 * Update UPDATE table SET column1 = 'value1', column2 = value2 ... condition;
 * @Params
 *   - DB  					database
 *	 - schema				schema, "" will ignore it
 *	 - table     		schema.table
 *	 - condition		e.g. WHERE name = 'name' AND value >= 10
 *	 - new 					object{field:value}
 **/
func UpdateDB(DB *sql.DB, schema, table, condition string, new map[string]interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if schema != "" {
		table = fmt.Sprintf("%s.%s", schema, table)
	}

	updateColumns, args, err := ToUpdateSQL(new)
	if err != nil {
		tx.Rollback()
		return err
	}
	update := fmt.Sprintf("UPDATE %s SET %s %s;", table, updateColumns, condition)
	if _, err := tx.Exec(update, args...); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateDeepDB check exist first. If not, insert it. If yes, delete and than insert
func UpdateDeepDB(DB *sql.DB, schema, table, condition string, object interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	// query
	results, err := QueryTx(tx, schema, table, condition, object)
	if err != nil {
		tx.Rollback()
		return err
	}

	// if not exist, insert it
	if len(results) == 0 {
		if err := InsertTx(tx, schema, table, []interface{}{object}); err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}

	// if too much result
	if len(results) != 1 {
		tx.Rollback()
		return errors.New("the query contain multiple result")
	}

	// if exist and equal value, do nothing
	result := results[0]
	if structEqual(result, object) {
		tx.Rollback()
		return nil
	}

	// if exit but not equal, delete and insert
	if err := DeleteTx(tx, schema, table, condition); err != nil {
		tx.Rollback()
		return err
	}
	if err := InsertTx(tx, schema, table, []interface{}{object}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Delete
/**
 * Delete DELETE FROM table condition;
 * @Params
 *   - DB  					database
 *	 - schema				schema, "" will ignore it
 *	 - table     		schema.table
 *	 - condition		e.g. WHERE name = 'name' AND value >= 10
 **/
func DeleteDB(DB *sql.DB, schema, table, condition string) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if err := DeleteTx(tx, schema, table, condition); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// DeleteTx
func DeleteTx(tx *sql.Tx, schema, table, condition string) error {
	if schema != "" {
		table = fmt.Sprintf("%s.%s", schema, table)
	}

	delete := fmt.Sprintf("DELETE FROM %s %s;", table, condition)
	if _, err := tx.Exec(delete); err != nil {
		return err
	}
	return nil
}

// ToInsertSQL
func ToInsertSQL(new interface{}) (string, string, []interface{}, error) {
	keys := ""
	values := ""
	args := make([]interface{}, 0)
	t := reflect.TypeOf(new)
	v := reflect.ValueOf(new)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return keys, values, args, errors.New("object not struct in update sql")
	}
	numField := t.NumField()
	for i := 0; i < numField-1; i++ {
		field := t.Field(i)
		value := v.Field(i)
		keys += fmt.Sprintf("%s, ", strings.ToLower(field.Name))
		values += fmt.Sprintf("$%d, ", i+1)
		// add args, but careful of arg is array
		if arr := ToSQLArray(value); arr != nil {
			args = append(args, arr)
			continue
		}
		args = append(args, value.Interface())
	}
	//last one
	field := t.Field(numField - 1)
	value := v.Field(numField - 1)
	keys += strings.ToLower(field.Name)
	values += fmt.Sprintf("$%d", numField)
	// add args, but careful of arg is array
	if arr := ToSQLArray(value); arr != nil {
		args = append(args, arr)
		return keys, values, args, nil
	}
	args = append(args, value.Interface())
	return keys, values, args, nil
}

// ToUpdateSQL return column1 = 'value', column2 = 12.123
func ToUpdateSQL(new map[string]interface{}) (string, []interface{}, error) {
	sql := ""
	args := make([]interface{}, 0)
	i := 0
	for k, v := range new {
		// last one
		if i == len(new)-1 {
			sql += fmt.Sprintf("%s = $%d", strings.ToLower(k), i+1)
			value := reflect.ValueOf(v)
			if arr := ToSQLArray(value); arr != nil {
				args = append(args, arr)
				return sql, args, nil
			}
			args = append(args, v)
			return sql, args, nil
		}

		i += 1
		sql += fmt.Sprintf("%s = $%d, ", strings.ToLower(k), i)
		value := reflect.ValueOf(v)
		if arr := ToSQLArray(value); arr != nil {
			args = append(args, arr)
			continue
		}
		args = append(args, v)
	}
	return sql, args, nil
}

// ToSQLArray
func ToSQLArray(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for _, driver := range sql.Drivers() {
			if driver == "postgres" {
				return pq.Array(value.Addr().Interface())
			}
		}
	}
	return nil
}

// structEqual
func structEqual(object1, object2 interface{}) bool {
	object1Value := reflect.ValueOf(object1)
	object2Value := reflect.ValueOf(object2)
	if object1Value.Kind() == reflect.Ptr {
		object1Value = object1Value.Elem()
	}
	if object2Value.Kind() == reflect.Ptr {
		object2Value = object2Value.Elem()
	}
	return reflect.DeepEqual(object1Value.Interface(), object2Value.Interface())
}
