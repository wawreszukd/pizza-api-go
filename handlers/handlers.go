package handlers

import (
	"encoding/json"
	"net/http"
	"simpledbservice/db"
	"strconv"
	"strings"
)

type Handlers struct {
	Db *db.DbHandler
}

func New(database *db.DbHandler) *Handlers {
	return &Handlers{Db: database}
}

func (h *Handlers) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	pizzas := h.Db.GetAll()
	err := json.NewEncoder(w).Encode(pizzas)
	if err != nil {
		panic(err)
	}
}
func (h *Handlers) HandlePost(w http.ResponseWriter, r *http.Request) {
	price, err := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("price")), 64)
	if err != nil {
		panic(err)
	}
	h.Db.CreatePizza(r.URL.Query().Get("name"), price, r.URL.Query().Get("topping"))
}
func (h *Handlers) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	price, err := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("price")), 64)
	if err != nil {
		panic(err)
	}
	id, err = h.Db.UpdatePizza(id, r.URL.Query().Get("name"), price, r.URL.Query().Get("topping"))
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(strconv.Itoa(id) + " updated"))
	if err != nil {
		panic(err)
	}
}
func (h *Handlers) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	id, err = h.Db.DeletePizza(id)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(strconv.Itoa(id) + " deleted"))
	if err != nil {
		panic(err)
	}
}
