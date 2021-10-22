package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"todo-app/backend/prisma/db"

	"github.com/gorilla/mux"
)

type Task struct {
	Id        string `json:"id"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"update_at"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	Desc      string `json:"desc"`
}
type Comment struct {
	Id        string `json:"id"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"update_at"`
	Content   string `json:"content"`
}

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router := mux.NewRouter()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var newTask Task
		err := r.ParseForm()
		if err != nil {
			io.WriteString(w, "create error")
		}
		json.NewDecoder(r.Body).Decode(&newTask)

		created, err := client.Task.CreateOne(
			db.Task.Title.Set(newTask.Title),
			db.Task.Status.Set(false),
			db.Task.Desc.Set(newTask.Desc),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "create error")
		}
		var response struct {
			Code int    `json:"code"`
			Data []Task `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, Task{
			Id:        created.ID,
			CreatedAt: created.CreatedAt.String(),
			UpdatedAt: created.UpdatedAt.String(),
			Title:     created.Title,
			Status:    false,
			Desc: func() string {
				desc, ok := created.Desc()
				if !ok {
					desc = ""
				}
				return desc
			}(),
		})
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("POST")

	router.HandleFunc("/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]
		task, err := client.Task.FindUnique(
			db.Task.ID.Equals(taskId),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "read error")
		}
		var response struct {
			Code int    `json:"code"`
			Data []Task `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, Task{
			Id:        task.ID,
			CreatedAt: task.CreatedAt.String(),
			UpdatedAt: task.UpdatedAt.String(),
			Title:     task.Title,
			Status:    task.Status,
			Desc: func() string {
				desc, ok := task.Desc()
				if !ok {
					desc = ""
				}
				return desc
			}(),
		})
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("GET")

	router.HandleFunc("/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]
		var newTask Task
		err := r.ParseForm()
		if err != nil {
			io.WriteString(w, "update error")
		}
		json.NewDecoder(r.Body).Decode(&newTask)

		updated, err := client.Task.FindUnique(
			db.Task.ID.Equals(taskId),
		).Update(
			db.Task.Title.Set(newTask.Title),
			db.Task.Desc.Set(newTask.Desc),
			db.Task.Status.Set(newTask.Status),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "update error")
		}
		var response struct {
			Code int    `json:"code"`
			Data []Task `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, Task{
			Id:        updated.ID,
			CreatedAt: updated.CreatedAt.String(),
			UpdatedAt: updated.UpdatedAt.String(),
			Title:     updated.Title,
			Status:    updated.Status,
			Desc: func() string {
				desc, ok := updated.Desc()
				if !ok {
					desc = ""
				}
				return desc
			}(),
		})
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("POST")

	router.HandleFunc("/tasks/{task_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]

		_, err := client.Task.FindUnique(
			db.Task.ID.Equals(taskId),
		).Delete().Exec(context.Background())
		if err != nil {
			io.WriteString(w, "delete error")
		}

		tasks, err := client.Task.FindMany().Exec(context.Background())
		if err != nil {
			io.WriteString(w, "read all error")
		}
		var responseData []Task
		for _, task := range tasks {
			responseData = append(responseData, Task{
				Id:        task.ID,
				CreatedAt: task.CreatedAt.String(),
				UpdatedAt: task.UpdatedAt.String(),
				Title:     task.Title,
				Status:    false,
				Desc: func() string {
					desc, ok := task.Desc()
					if !ok {
						desc = ""
					}
					return desc
				}(),
			})
		}
		var response struct {
			Code int    `json:"code"`
			Data []Task `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, responseData...)
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("DELETE")

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := client.Task.FindMany().Exec(context.Background())
		if err != nil {
			io.WriteString(w, "read all error")
		}
		var responseData []Task
		for _, task := range tasks {
			responseData = append(responseData, Task{
				Id:        task.ID,
				CreatedAt: task.CreatedAt.String(),
				UpdatedAt: task.UpdatedAt.String(),
				Title:     task.Title,
				Status:    false,
				Desc: func() string {
					desc, ok := task.Desc()
					if !ok {
						desc = ""
					}
					return desc
				}(),
			})
		}
		var response struct {
			Code int    `json:"code"`
			Data []Task `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, responseData...)
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("GET")

	router.HandleFunc("/tasks/{task_id}/comments", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]
		var newComment Comment
		err := r.ParseForm()
		if err != nil {
			io.WriteString(w, "create error")
		}
		json.NewDecoder(r.Body).Decode(&newComment)

		created, err := client.Comment.CreateOne(
			db.Comment.Content.Set(newComment.Content),
			db.Comment.Task.Link(
				db.Task.ID.Equals(taskId),
			),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "create error")
		}
		var response struct {
			Code int       `json:"code"`
			Data []Comment `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, Comment{
			Id:        created.ID,
			CreatedAt: created.CreatedAt.String(),
			UpdatedAt: created.UpdatedAt.String(),
			Content:   created.Content,
		})
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("POST")

	router.HandleFunc("/tasks/{task_id}/comments", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]
		comments, err := client.Comment.FindMany(
			db.Comment.TaskID.Equals(taskId),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "read error")
		}
		var responseData []Comment
		for _, comment := range comments {
			responseData = append(responseData, Comment{
				Id:        comment.ID,
				CreatedAt: comment.CreatedAt.String(),
				UpdatedAt: comment.UpdatedAt.String(),
				Content:   comment.Content,
			})
		}
		var response struct {
			Code int       `json:"code"`
			Data []Comment `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, responseData...)
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("GET")

	router.HandleFunc("/tasks/{task_id}/comments/{comment_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		commentId := params["comment_id"]
		var newComment Comment
		err := r.ParseForm()
		if err != nil {
			io.WriteString(w, "update error")
		}
		json.NewDecoder(r.Body).Decode(&newComment)

		updated, err := client.Comment.FindUnique(
			db.Comment.ID.Equals(commentId),
		).Update(
			db.Comment.Content.Set(newComment.Content),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "update error")
		}
		var response struct {
			Code int       `json:"code"`
			Data []Comment `json:"comment"`
		}
		response.Code = 200
		response.Data = append(response.Data, Comment{
			Id:        updated.ID,
			CreatedAt: updated.CreatedAt.String(),
			UpdatedAt: updated.UpdatedAt.String(),
			Content:   updated.Content,
		})
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("POST")

	router.HandleFunc("/tasks/{task_id}/comments/{comment_id}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		taskId := params["task_id"]
		commentId := params["comment_id"]

		_, err := client.Comment.FindUnique(
			db.Comment.ID.Equals(commentId),
		).Delete().Exec(context.Background())
		if err != nil {
			io.WriteString(w, "delete error")
		}

		comments, err := client.Comment.FindMany(
			db.Comment.TaskID.Equals(taskId),
		).Exec(context.Background())
		if err != nil {
			io.WriteString(w, "read all error")
		}
		var responseData []Comment
		for _, comment := range comments {
			responseData = append(responseData, Comment{
				Id:        comment.ID,
				CreatedAt: comment.CreatedAt.String(),
				UpdatedAt: comment.UpdatedAt.String(),
				Content:   comment.Content,
			})
		}
		var response struct {
			Code int       `json:"code"`
			Data []Comment `json:"data"`
		}
		response.Code = 200
		response.Data = append(response.Data, responseData...)
		resp, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}).Methods("DELETE")

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
