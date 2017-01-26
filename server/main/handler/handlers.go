package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"../model"
	"../repo"
	"github.com/gorilla/mux"
)

//TodoIndex route
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	setContentType(w)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repo.GetTodo()); err != nil {
		panic(err)
	}
}

//TodoShow route
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}

//TodoCreate route
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) //This is a good way to protect against malicious attacks on your server. Imagine if someone wanted to send you 500GBs of json!
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		setContentType(w)
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := repo.RepoCreateTodo(todo)
	setContentType(w)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
