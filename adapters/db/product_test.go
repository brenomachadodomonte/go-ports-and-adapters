package db_test

import (
	"database/sql"
	"github.com/brenomachadodomonte/go-ports-and-adapters/adapters/db"
	"github.com/brenomachadodomonte/go-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "create table products(id string, name string, price float, status price);"
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func createProduct(db *sql.DB) {
	insert := "insert into products(id, name, price, status) values (?, ?, ?, ?)"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec("abc", "Product1", 0, "disabled")
	if err != nil {
		log.Fatal(err.Error())
	}

}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product.GetName(), "Product1")
	require.Equal(t, product.GetPrice(), 0.0)
	require.Equal(t, product.GetStatus(), "disabled")
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 10.0

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	err = product.Enable()
	require.Nil(t, err)

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
