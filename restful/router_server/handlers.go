package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

var todos = Todos{
	Todo{Name: "Write presentation"},
	Todo{Name: "Host meetup"},
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if todoId, err := strconv.Atoi(vars["todoId"]); err != nil {
		fmt.Fprintln(w, err)
	} else {
		todo := todos[todoId % len(todos)]
		fmt.Fprintln(w, "Todo Show:", todo)
	}
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(todos)
}
