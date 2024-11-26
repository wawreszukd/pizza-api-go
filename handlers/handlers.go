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
func (h *Handlers) HandleGetOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id == 0 {
		w.Write([]byte("Invalid id"))
		return
	}
	pizza, err := h.Db.GetOne(id)
	if err != nil {
		w.Write([]byte("Error: invalid id"))
		return
	}
	err = json.NewEncoder(w).Encode(pizza)
	if err != nil {
		panic(err)
	}
}
func (h *Handlers) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	pizzas := h.Db.GetAll()
	err := json.NewEncoder(w).Encode(pizzas)
	if err != nil {
		panic(err)
	}
}
func (h *Handlers) HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("Invalid request method"))
		return
	}
	price, err := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("price")), 64)
	if err != nil {
		w.Write([]byte("Invalid price"))
		return
	}
	err = h.Db.CreatePizza(r.URL.Query().Get("name"), price, r.URL.Query().Get("topping"))
	if err != nil {
		w.Write([]byte("Error creating pizza"))
	}
}
func (h *Handlers) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Write([]byte("Invalid request method"))
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id == 0 {
		w.Write([]byte("Invalid id"))
		return
	}
	price, err := strconv.ParseFloat(strings.TrimSpace(r.URL.Query().Get("price")), 64)
	if err != nil {
		w.Write([]byte("Invalid price"))
	}
	id, err = h.Db.UpdatePizza(id, r.URL.Query().Get("name"), price, r.URL.Query().Get("topping"))
	if err != nil {
		w.Write([]byte("Error updating pizza"))
	}
	_, err = w.Write([]byte(strconv.Itoa(id) + " updated"))
	if err != nil {
		panic(err)
	}
}
func (h *Handlers) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Write([]byte("Invalid request method"))
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id == 0 {
		w.Write([]byte("Invalid id"))
		return
	}
	id, err = h.Db.DeletePizza(id)
	if err != nil {
		w.Write([]byte("Error deleting pizza"))
	}
	_, err = w.Write([]byte(strconv.Itoa(id) + " deleted"))
	if err != nil {
		panic(err)
	}
}
