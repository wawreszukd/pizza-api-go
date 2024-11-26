package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"simpledbservice/models"
)

type DbHandler struct {
	Db *sql.DB
}

func New() *DbHandler {
	return &DbHandler{}
}

func (db *DbHandler) New() {
	// Connect to the database

	database, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	db.Db = database
	err = db.Db.Ping()
	if err != nil {
		panic(err)
	}
	_, err = database.Exec("	CREATE TABLE IF NOT EXISTS pizza (id INTEGER PRIMARY KEY, name TEXT, price REAL, topping TEXT)")
	if err != nil {
		return
	}
}
func (db *DbHandler) Close() {
	err := db.Db.Close()
	if err != nil {
		panic(err)
	}
}
func (db *DbHandler) GetOne(id int) (models.Pizza, error) {
	var pizza models.Pizza
	rows, err := db.Db.Query("SELECT * FROM pizza WHERE id=$1", id)
	if err != nil {
		return pizza, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Topping)
		if err != nil {
			return pizza, err
		}
	}
	if pizza.ID == 0 {
		return models.Pizza{}, errors.New("No pizza found")
	}
	return pizza, nil
}
func (db *DbHandler) GetAll() []models.Pizza {
	rows, err := db.Db.Query("SELECT * FROM pizza")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	var pizzas []models.Pizza
	for rows.Next() {
		var pizza models.Pizza
		err := rows.Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Topping)
		if err != nil {
			panic(err)
		}
		pizzas = append(pizzas, pizza)
	}
	return pizzas
}
func (db *DbHandler) CreatePizza(name string, price float64, topping string) error {
	_, err := db.Db.Exec("INSERT INTO pizza (name, price, topping) VALUES ($1, $2, $3)", name, price, topping)
	if err != nil {
		return err
	}
	return nil
}

func (db *DbHandler) UpdatePizza(id int, name string, price float64, topping string) (int, error) {
	_, err := db.Db.Exec("UPDATE pizza SET name=$1, price=$2, topping=$3 WHERE id=$4", name, price, topping, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (db *DbHandler) DeletePizza(id int) (int, error) {
	_, err := db.Db.Exec("DELETE FROM pizza WHERE id=$1", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
