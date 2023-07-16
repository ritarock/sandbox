package http

import (
	"net/http"
	"strconv"
	"todo-clean2/domain"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TodoHandler struct {
	TUsecase domain.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, tu domain.TodoUsecase) {
	handler := &TodoHandler{
		TUsecase: tu,
	}
	e.GET("/todo/:id", handler.Get)
	e.GET("/todos", handler.GetList)
	e.POST("/todo", handler.Create)
}

func (t *TodoHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()

	todo, err := t.TUsecase.Get(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (t *TodoHandler) GetList(c echo.Context) error {
	ctx := c.Request().Context()

	todos, err := t.TUsecase.GetList(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) Create(c echo.Context) error {
	var todo domain.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if err := t.TUsecase.Create(ctx, &todo); err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, todo)
}
