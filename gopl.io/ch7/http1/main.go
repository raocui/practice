// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 191.

// Http1 is a rudimentary e-commerce server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	//服务
	//mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/item", http.HandlerFunc(db.item))

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/item", db.item)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) item(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	w.Write([]byte(fmt.Sprintf("%s: %s\n", item, db[item])))

}
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}

}

//!-main

/*
//!+handler
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
//!-handler
*/
