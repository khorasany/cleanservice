package echo

import (
	"net/http"

	"github.com/kianooshaz/cleanservice/internal/pkg/logger"
	"github.com/kianooshaz/cleanservice/internal/service"
	"github.com/kianooshaz/cleanservice/internal/transport/http/request"
	"github.com/kianooshaz/cleanservice/internal/transport/http/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	logger  logger.Logger
	account service.Account
}

func (h *handler) login(c echo.Context) error {
	// start span with context

	req := new(request.Login)

	if err := c.Bind(req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	token, err := h.account.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	}

	return c.JSON(http.StatusOK, response.Login{Token: token, Status: http.StatusOK})
}

func (h *handler) register(c echo.Context) error {
	return nil
}
