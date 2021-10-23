package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is required for new product\n")
		return
	}

	if len(price) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price is required for new product: %q\n", item)
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	i, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "server error: %v\n", err)
		return
	}

	db[item] = dollars(i)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "item: %q, price: %s was added to the database\n", item, dollars(i))
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "item: %q, price: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is required to update product\n")
		return
	}

	if len(price) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price is required to update product: %q\n", item)
		return
	}

	i, err := strconv.Atoi(price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "server error: %v\n", err)
		return
	}

	db[item] = dollars(i)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item: %q, price: %s was successfully updated in the database\n", item, dollars(i))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is required to delete product\n")
		return
	}

	delete(db, item)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item: %q was successfully deleted from the database\n", item)
}
