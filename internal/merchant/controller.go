package merchant

import (
	"net/http"

	"github.com/etnk125/go-webserver-merchant-management/model"
	"github.com/labstack/echo/v4"
)

type merchantService interface {
	RegisterMerchant(req *model.RegisterMerchantRequest) (*model.Merchant, error)
	GetMerchantInfo(id string) (*model.Merchant, error)
	UpdateMerchantInfo(id string, req *model.UpdateMerchantRequest) error
}

type Controller struct {
	merchantService merchantService
}

func NewController(service merchantService) *Controller {
	return &Controller{
		merchantService: service,
	}
}

func (c *Controller) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func (c *Controller) RegisterMerchant(ctx echo.Context) error {
	req := new(model.RegisterMerchantRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}
	merchant, err := c.merchantService.RegisterMerchant(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, merchant)
}

func (c *Controller) GetMerchantInfo(ctx echo.Context) error {
	id := ctx.Param("id")
	merchant, err := c.merchantService.GetMerchantInfo(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, merchant)
}

func (c *Controller) UpdateMerchantInfo(ctx echo.Context) error {
	merchantID := ctx.Param("id")
	req := new(model.UpdateMerchantRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// update merchant info
	c.merchantService.UpdateMerchantInfo(merchantID, req)

	return ctx.JSON(http.StatusOK, "JSON")
}

func (c *Controller) AddProduct(ctx echo.Context) error {

	// add product
	return ctx.JSON(http.StatusCreated, "Created")
}

func (c *Controller) GetProducts(ctx echo.Context) error {
	// get products
	return ctx.JSON(http.StatusOK, "JSON")
}

func (c *Controller) BuyProduct(ctx echo.Context) error {
	// buy product
	return ctx.JSON(http.StatusOK, "JSON")
}

func (c *Controller) GetSellReport(ctx echo.Context) error {
	// get sell report
	return ctx.JSON(http.StatusOK, "JSON")
}
