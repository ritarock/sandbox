package server

import (
	"backend/internal/model"
	"backend/internal/server/types"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetTasks()
	if err != nil {
		io.WriteString(w, "ERROR: read tasks")
	}
	var responseData []types.Task
	for _, task := range tasks {
		responseData = append(responseData, types.Task{
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
		Code int          `json:"code"`
		Data []types.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, responseData...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask types.Task
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, "ERROR: create task")
	}
	json.NewDecoder(r.Body).Decode(&newTask)
	created, err := model.CreateTask(newTask)
	if err != nil {
		io.WriteString(w, "ERROR: create task")
	}
	var response struct {
		Code int          `json:"code"`
		Data []types.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Task{
		Id:        created.ID,
		CreatedAt: created.CreatedAt.String(),
		UpdatedAt: created.UpdatedAt.String(),
		Title:     created.Title,
		Status:    created.Status,
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
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]

	task, err := model.GetTask(taskId)
	if err != nil {
		io.WriteString(w, "ERROR: read task")
	}

	var response struct {
		Code int          `json:"code"`
		Data []types.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Task{
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
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]
	var newTask types.Task
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, "ERROR: update task")
	}
	json.NewDecoder(r.Body).Decode(&newTask)

	updated, err := model.UpdateTask(taskId, newTask)
	if err != nil {
		io.WriteString(w, "ERROR: update task")
	}
	var response struct {
		Code int          `json:"code"`
		Data []types.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Task{
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
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]

	err := model.DeleteTask(taskId)
	if err != nil {
		io.WriteString(w, "ERROR: delete task")
	}
	tasks, err := model.GetTasks()
	if err != nil {
		io.WriteString(w, "ERROR: read all task")
	}
	var responseData []types.Task
	for _, task := range tasks {
		responseData = append(responseData, types.Task{
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
		Code int          `json:"code"`
		Data []types.Task `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, responseData...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]

	comments, err := model.GetComments(taskId)
	if err != nil {
		io.WriteString(w, "ERROR: read comments")
	}
	var responseData []types.Comment
	for _, comment := range comments {
		responseData = append(responseData, types.Comment{
			Id:        comment.ID,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
			Content:   comment.Content,
		})
	}
	var response struct {
		Code int             `json:"code"`
		Data []types.Comment `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, responseData...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]
	var newComment types.Comment
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, "ERROR: create comment")
	}
	json.NewDecoder(r.Body).Decode(&newComment)

	created, err := model.CreateComment(taskId, newComment)
	if err != nil {
		io.WriteString(w, "ERROR: create comment")
	}
	var response struct {
		Code int             `json:"code"`
		Data []types.Comment `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment{
		Id:        created.ID,
		CreatedAt: created.CreatedAt.String(),
		UpdatedAt: created.UpdatedAt.String(),
		Content:   created.Content,
	})
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func getComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentId := params["comment_id"]

	comment, err := model.GetComment(commentId)
	if err != nil {
		io.WriteString(w, "ERROR: read comment")
	}

	var response struct {
		Code int             `jsno:"code"`
		Data []types.Comment `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment{
		Id:        comment.ID,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
		Content:   comment.Content,
	})
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func updateComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentId := params["comment_id"]
	var newComment types.Comment
	err := r.ParseForm()
	if err != nil {
		io.WriteString(w, "ERROR: update comment")
	}
	json.NewDecoder(r.Body).Decode(&newComment)

	updated, err := model.UpdateComment(commentId, newComment)
	if err != nil {
		io.WriteString(w, "update error")
	}
	var response struct {
		Code int             `json:"code"`
		Data []types.Comment `json:"comment"`
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment{
		Id:        updated.ID,
		CreatedAt: updated.CreatedAt.String(),
		UpdatedAt: updated.UpdatedAt.String(),
		Content:   updated.Content,
	})
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId := params["task_id"]
	commentId := params["comment_id"]

	err := model.DeleteComment(commentId)
	if err != nil {
		io.WriteString(w, "ERROR: delete comment")
	}

	comments, err := model.GetComments(taskId)
	if err != nil {
		io.WriteString(w, "ERROR: read comments")
	}
	var responseData []types.Comment
	for _, comment := range comments {
		responseData = append(responseData, types.Comment{
			Id:        comment.ID,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
			Content:   comment.Content,
		})
	}
	var response struct {
		Code int             `json:"code"`
		Data []types.Comment `json:"data"`
	}
	response.Code = 200
	response.Data = append(response.Data, responseData...)
	resp, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
