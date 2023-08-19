package services

import (
	"errors"
	"github.com/farmani/sharebuy/cmd/api/requests"
	"strconv"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/repository"
	"github.com/labstack/echo/v4"
)

type ProductService struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewProductService(a *app.Application) *ProductService {
	return &ProductService{
		app: a,
	}
}

func (ps *ProductService) FindByUuid(c echo.Context, uuid string) (models.Product, error) {
	if uuid == "" {
		return models.Product{}, repository.ErrRecordNotFound
	}

	return ps.app.Repository.Product().GetByUuid(uuid)
}

func (ps *ProductService) FindById(c echo.Context, id int64) (models.Product, error) {
	if id < 1 {
		return models.Product{}, repository.ErrRecordNotFound
	}

	return ps.app.Repository.Product().GetById(id)
}

func (ps *ProductService) Update(c echo.Context, pr requests.ProductUpdateRequest) (models.Product, error) {

	product, err := ps.app.Repository.Product().GetByUuid(pr.UUID)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			return models.Product{}, repository.ErrRecordNotFound
		default:
			return models.Product{}, err
		}
	}

	if c.Request().Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(product.Version), 10) != c.Request().Header.Get("X-Expected-Version") {
			return models.Product{}, repository.ErrEditConflict
		}
	}

	// Pass the updated product record to the Update() method.
	err = ps.app.Repository.Product().Update(&product)
	if err != nil {
		return models.Product{}, repository.ErrSavingDataFailed
	}

	return product, nil
}
