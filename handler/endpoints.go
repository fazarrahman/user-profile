package handler

import (
	"fmt"
	"net/http"

	"github.com/fazarrahman/user-profile/generated"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.Response
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

// This is just a test endpoint to get you started. Please delete this endpoint.
// (POST /user-register)
func (s *Server) UserRegister(ctx echo.Context) error {
	var req generated.Users
	var resp generated.CreateResponse
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	createdUser, err := s.Svc.CreateUser(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(err.StatusCode, echo.Map{"error": err.Message})
	}
	resp.Id = createdUser.Id
	return ctx.JSON(http.StatusOK, resp)
}
