package handlers

import (
	"errors"
	"github.com/farmani/sharebuy/internal/repository"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/requests"
	"github.com/farmani/sharebuy/cmd/api/responses"
	"github.com/farmani/sharebuy/cmd/api/services"
	"github.com/farmani/sharebuy/pkg/validator"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	// Dependencies or state for the UserHandler
	app            *app.Application
	productService *services.ProductService
}

func NewProductHandler(a *app.Application, ps *services.ProductService) *ProductHandler {
	return &ProductHandler{
		app:            a,
		productService: ps,
	}
}

func (h *ProductHandler) RegisterRoutes(e *echo.Echo) {
	// Products handlers
	e.GET("/products/v1/", h.ProductList)
	// e.GET("/products/v1/:uuid", h.ProductView)
	// e.GET("/products/v1/:uuid", h.ProductUpdate)

	// Suggested Products handlers (Authenticated)
	// e.POST("/products/v1/suggest", h.SuggestProduct, session.MiddlewareWithConfig(session.Config{}))
	// e.GET("/products/v1/suggest", h.GetSuggestedProductList, session.MiddlewareWithConfig(session.Config{}))
	// e.GET("/products/v1/suggest/:id", h.GetSuggestedProductView, session.MiddlewareWithConfig(session.Config{}))
}

func (h *ProductHandler) ProductList(c echo.Context) error {
	pr, err := h.productService.FindById(c, 1)
	if err != nil {
		switch err {
		case repository.ErrRecordNotFound:
			notFoundResponse(c)
		default:
			serverErrorResponse(c, err)
		}
		return nil
	}

	return responses.WriteJSON(c, responses.NewProductResponse(&pr))
}

func (h *ProductHandler) ProductUpdate(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid == "" {
		notFoundResponse(c)
		return nil
	}

	// validate request
	v := validator.New(c)
	r, err := requests.NewProductUpdateRequest(c, v)
	if err != nil {
		switch {
		case errors.Is(err, requests.ErrBadRequest):
			badRequestResponse(c, err)
			return nil
		case errors.Is(err, requests.ErrFailedValidation):
			failedValidationResponse(c, v.Errors)
			return nil
		default:
			serverErrorResponse(c, err)
			return nil
		}
	}

	p, err := h.productService.Update(c, r)
	if err != nil {
		serverErrorResponse(c, err)
		return nil
	}

	return responses.WriteJSON(c, responses.NewProductResponse(&p))
}
