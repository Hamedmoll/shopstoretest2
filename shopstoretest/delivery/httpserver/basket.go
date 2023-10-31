package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/param"
)

func (s Server) addBasket(c echo.Context) error {
	req := param.AddToBasketRequest{}
	bErr := c.Bind(&req)
	if bErr != nil {

		return bErr
	}

	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if claim.ID != req.UserID {

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	res, aErr := s.productService.AddBasket(req)
	if aErr != nil {

		return aErr
	}

	return c.JSON(http.StatusOK, res)
}

func (s Server) showBaskets(c echo.Context) error {
	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	req := param.ShowBasketRequest{ID: claim.ID}

	res, sErr := s.productService.ShowBasket(req)
	if sErr != nil {

		return sErr
	}

	return c.JSON(http.StatusOK, res)
}
