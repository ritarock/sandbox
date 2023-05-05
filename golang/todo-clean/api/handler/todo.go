package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-clean/api/presenter"
	"todo-clean/entity"
	"todo-clean/usecase/todo"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeTodoHandlers(r *mux.Router, n negroni.Negroni, service todo.UseCase) {
	r.Handle("/v1/todo", n.With(
		negroni.Wrap(listTodos(service)),
	)).Methods("GET", "OPTIONS").Name("listTodos")

	r.Handle("/v1/todo/{id}", n.With(
		negroni.Wrap(getTodo(service)),
	)).Methods("GET", "OPTIONS").Name("getTodo")

	r.Handle("/v1/todo", n.With(
		negroni.Wrap(createTodo(service)),
	)).Methods("POST", "OPTIONS").Name("createTodo")
}

func getTodo(service todo.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error reading todo"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

		data, err := service.GetTodo(id)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Todo{
			ID:     data.ID,
			Title:  data.Title,
			Status: data.Status,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func listTodos(service todo.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error reading todos"

		data, err := service.GetTodoList()
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Todo
		for _, d := range data {
			toJ = append(toJ, &presenter.Todo{
				ID:     d.ID,
				Title:  d.Title,
				Status: d.Status,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createTodo(service todo.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error adding todo"
		var input struct {
			Title  string `json:"title"`
			Status bool   `json:"status"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateTodo(input.Title, input.Status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		toJ := &presenter.Todo{
			ID:     id,
			Title:  input.Title,
			Status: input.Status,
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}
