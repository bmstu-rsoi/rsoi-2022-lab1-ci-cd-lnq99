package controller

import (
	"fmt"
	"net/http"
	"rsoi-1/internal/model"
	errors "rsoi-1/internal/model/error"
	"rsoi-1/internal/service"

	"github.com/labstack/echo/v4"
)

// EchoController converts echo contexts to parameters.
type EchoController struct {
	services *service.Services
}

func NewEchoController(services *service.Services) *EchoController {
	return &EchoController{services}
}

func EchoResponseInternalServerError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{err.Error()})
}

func (c *EchoController) ListPersons(ctx echo.Context) error {
	r, err := c.services.Person.ListPersons()
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, r)
}

func (c *EchoController) GetPerson(ctx echo.Context) error {
	var id int32
	err := echo.PathParamsBinder(ctx).Int32("id", &id).BindError()
	if err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}

	r, err := c.services.Person.GetPerson(id)
	if err == errors.NotFound {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{"Not found Person for ID"})
	}
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, r)
}

func (c *EchoController) CreatePerson(ctx echo.Context) error {
	var person model.PersonRequest
	err := ctx.Bind(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}
	err = ctx.Validate(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}

	id, err := c.services.Person.CreatePerson(&person)
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}

	location := fmt.Sprintf("/api/v1/persons/%d", id)
	ctx.Response().Header().Set("Location", location)
	return ctx.NoContent(http.StatusCreated)
}

func (c *EchoController) EditPerson(ctx echo.Context) error {
	var id int32
	err := echo.PathParamsBinder(ctx).Int32("id", &id).BindError()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}

	var person model.PersonRequest
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}

	_, err = c.services.Person.GetPerson(id)
	if err == errors.NotFound {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{"Not found Person for ID"})
	}
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}

	err = ctx.Bind(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}
	err = ctx.Validate(&person)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ValidationErrorResponse{Message: err.Error()})
	}

	r, err := c.services.Person.EditPerson(id, &person)
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, r)
}

func (c *EchoController) DeletePerson(ctx echo.Context) error {
	var id int32
	err := echo.PathParamsBinder(ctx).Int32("id", &id).BindError()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{"Not found Person for ID"})
	}

	err = c.services.Person.DeletePerson(id)
	if err == errors.NoAffected {
		return ctx.JSON(http.StatusNotFound, model.ErrorResponse{"Not found Person for ID"})
	}
	if err != nil {
		return EchoResponseInternalServerError(ctx, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
