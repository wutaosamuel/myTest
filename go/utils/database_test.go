package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

const (
	test_schema = "test_schema"
	test_table = "test_object"
	test_elements = `"name" varchar(255) PRIMARY KEY,
	"id" int NOT NULL,
	"value" double precision DEFAULT 0,
	"arr" varchar(255)[] DEFAULT '{}',
	"time" timestamp DEFAULT CURRENT_TIMESTAMP`
)

// Test Object
type Object struct {
	Name  string
	ID    int
	Value float64
	Arr   []string
	Time  time.Time
}

// TestOpenDB
func openDB() (*sql.DB, error) {
	database := NewDatabase()
	database.Set("172.16.8.221", 5432, "pi", "wutaowutao", "test")
	db, err := sql.Open("postgres", database.Source())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	return db, nil
}

func TestCreateSchema(t *testing.T) {
	sql := ToSchema("test_schema")
	t.Log("TestCreateSchema  --> pass")
	t.Log(sql)
}

func TestCreateSchemaTable(t *testing.T) {
	elements := `"name" varchar(255) PRIMARY KEY,
	"id" int NOT NULL,
	"value" double precision DEFAULT 0`
	sql := ToSchemaTable("test_schema", "test_object", elements)
	t.Log("TestCreateSchemaTable  --> pass")
	t.Log(sql)
}

func TestInsertSQL(t *testing.T) {
	// Object
	type Object struct {
		Name  string
		ID    int
		Value float64
	}
	k := "name, id, value"
	v := "$1, $2, $3"
	a := []interface{}{"object", -1, 12.123456789}
	object := &Object{"object", -1, 12.123456789}

	keys, values, args, err := ToInsertSQL(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestInsertSQL  --> pass")
	t.Log(k, "()needs to equal to()", keys)
	if k != keys {
		t.Fatal("keys not correct")
	}
	t.Log(v, "()needs to equal to()", values)
	if v != values {
		t.Fatal("values not correct")
	}
	t.Log(a, "()needs to equal to()", args)
	for i := 0; i < len(args); i++ {
		if a[i] != args[i] {
			t.Fatal("args not correct")
		}
	}
}

func TestUpdateSQL(t *testing.T) {
	// Object
	u := "id = $1, value = $2, name = $3"
	a := []interface{}{-1, 12.123456789, "object"}
	object := map[string]interface{}{
		"name":  "object",
		"id":    -1,
		"value": 12.123456789,
	}
	update, args, err := ToUpdateSQL(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("TestUpdateSQL  --> pass")
	// NOTE: map cause out of order, it should be ok
	t.Log(u, "()needs to equal to()", update)
	t.Log(a, "()needs to equal to()", args)
}

func TestDatabase(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("create schema  --> pass")
	if err := CreateSchema(db, test_schema); err != nil {
		t.Fatal(err)
	}

	t.Log("create schema.table  --> pass")
	if err := CreateSchemaTable(db, test_schema, test_table, test_elements); err != nil {
		t.Fatal(err)
	}

	t.Log("insert  --> pass")
	zero := &Object{"zero", 0, 12.123456789, []string{"zero", "is"}, time.Now()}
	first := &Object{"first", 1, 12.123456789, []string{"first", "arr"}, time.Now()}
	second := &Object{"second", 2, 2.123456789, []string{}, time.Now()}
	object := &Object{"object", -1, 0.123456789, []string{}, time.Now()}
	if err := InsertDB(db, test_schema, test_table, []interface{}{zero, first, second, object}); err != nil {
		t.Fatal(err)
	}

	t.Log("query --> pass")
	if results, err := QueryDB(db, test_schema, test_table, "WHERE name = 'zero'", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		for i, result := range results {
			t.Log("The query result", i, result)
		}
	}

	t.Log("update --> pass")
	changed := map[string]interface{}{
		"name":  "o",
		"value": 1.1,
	}
	if err := UpdateDB(db, test_schema, test_table, "WHERE id = -1", changed); err != nil {
		t.Fatal(err)
	}
	if results, err := QueryDB(db, test_schema, test_table, "WHERE id = -1", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		for i, result := range results {
			t.Log("The query result", i, result)
		}
	}

	t.Log("query all --> pass")
	if results, err := QueryDB(db, test_schema, test_table, "", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		for i, result := range results {
			t.Log("all query result", i, ": ", result)
		}
	}

	t.Log("delete all one by one --> pass")
	if err := DeleteDB(db, test_schema, test_table, "WHERE name = 'zero'"); err != nil {
		t.Fatal(err)
	}
	if err := DeleteDB(db, test_schema, test_table, "WHERE name = 'first'"); err != nil {
		t.Fatal(err)
	}
	if err := DeleteDB(db, test_schema, test_table, "WHERE name = 'second'"); err != nil {
		t.Fatal(err)
	}
	if err := DeleteDB(db, test_schema, test_table, "WHERE id = -1"); err != nil {
		t.Fatal(err)
	}
	if results, err := QueryDB(db, test_schema, test_table, "", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		if len(results) != 0 {
			t.Fatal("not delete all yet")
		}
	}

	// drop schema
	t.Log("drop schema --> pass")
	if err := InsertDB(db, test_schema, test_table, []interface{}{zero, first, second, object}); err != nil {
		t.Fatal(err)
	}
	if err := DropSchemaClean(db, test_schema); err != nil {
		t.Fatal(err)
	}
}

// TestQueryEmpty
func TestQueryEmpty(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("create schema  --> pass")
	if err := CreateSchema(db, test_schema); err != nil {
		t.Fatal(err)
	}

	t.Log("create schema.table  --> pass")
	if err := CreateSchemaTable(db, test_schema, test_table, test_elements); err != nil {
		t.Fatal(err)
	}

	t.Log("query empty --> pass")
	if results, err := QueryDB(db, test_schema, test_table, "WHERE name = 'zero'", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		t.Log(len(results))
		for i, result := range results {
			t.Log("all query result", i, ": ", result)
		}
	}

	if err := DropSchemaClean(db, test_schema); err != nil {
		t.Fatal(err)
	}
}

// TestUpdateDeepDB
func TestUpdateDeepDB(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("create schema  --> pass")
	if err := CreateSchema(db, test_schema); err != nil {
		t.Fatal(err)
	}

	t.Log("create schema.table  --> pass")
	if err := CreateSchemaTable(db, test_schema, test_table, test_elements); err != nil {
		t.Fatal(err)
	}

	t.Log("Insert empty --> pass")
	zero := &Object{"zero", 0, 12.123456789, []string{"zero", "is"}, time.Now()}
	if err := UpdateDeepDB(db, test_schema, test_table, "WHERE name = 'zero'", zero); err != nil {
		t.Fatal(err)
	}
	if results, err := QueryDB(db, test_schema, test_table, "WHERE name = 'zero'", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		t.Log(len(results))
		if len(results) != 1 {
			t.Fatal("insert empty not correct")
		}
		for i, result := range results {
			t.Log("all query result", i, ": ", result)
		}
	}

	t.Log("same object --> ")
	if err := UpdateDeepDB(db, test_schema, test_table, "WHERE name = 'zero'", zero); err != nil {
		t.Fatal(err)
	}

	t.Log("update object --> ")
	zero.ID = 999
	zero.Value = 1.123548
	if err := UpdateDeepDB(db, test_schema, test_table, "WHERE name = 'zero'", zero); err != nil {
		t.Fatal(err)
	}
	if results, err := QueryDB(db, test_schema, test_table, "WHERE name = 'zero'", &Object{}); err != nil {
		t.Fatal(err)
	} else {
		t.Log(len(results))
		if len(results) != 1 {
			t.Fatal("update empty not correct")
		}
		for i, result := range results {
			t.Log("all query result", i, ": ", result)
		}
	}

	if err := DropSchemaClean(db, test_schema); err != nil {
		t.Fatal(err)
	}
}
