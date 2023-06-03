package http

import (
	"net/http"
	"strconv"
	"todo-clean-arch/domain"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TodoHandler struct {
	Tusecase domain.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, us domain.TodoUsecase) {
	handler := &TodoHandler{
		Tusecase: us,
	}
	e.POST("/todos", handler.Create)
	e.GET("/todos/:id", handler.GetByID)
	e.PUT("/todos/:id", handler.Update)
	e.DELETE("/todos/:id", handler.Delete)
}

func (t *TodoHandler) Create(c echo.Context) (err error) {
	var todo domain.Todo
	err = c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if err = t.Tusecase.Create(ctx, &todo); err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func (t *TodoHandler) GetByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := c.Request().Context()

	todo, err := t.Tusecase.GetById(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}

func (t *TodoHandler) Update(c echo.Context) (err error) {
	var todo domain.Todo
	err = c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if err = t.Tusecase.Update(ctx, &todo); err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func (t *TodoHandler) Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()
	err = t.Tusecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
