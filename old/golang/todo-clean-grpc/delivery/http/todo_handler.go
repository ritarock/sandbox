package http

import (
	"net/http"
	"todo-clean-grpc/domain"

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
	e.POST("/todo", handler.Create)
	e.GET("/todo/:title", handler.Get)
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

func (t *TodoHandler) Get(c echo.Context) error {
	title := c.Param("title")
	ctx := c.Request().Context()

	todo, err := t.TUsecase.Get(ctx, title)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}
