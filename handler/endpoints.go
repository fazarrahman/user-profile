package handler

import (
	"net/http"
	"strings"

	"github.com/fazarrahman/user-profile/generated"
	"github.com/labstack/echo/v4"
)

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

// (POST /login)
func (s *Server) Login(ctx echo.Context) error {
	var req generated.LoginInput
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	loginResp, err := s.Svc.Login(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.JSON(err.StatusCode, echo.Map{"error": err.Message})
	}
	return ctx.JSON(http.StatusOK, loginResp)
}

// (GET /user)
func (s *Server) GetUser(ctx echo.Context) error {
	accessToken := extractToken(ctx)
	userResp, err := s.Svc.GetUserByAccessToken(ctx.Request().Context(), accessToken)
	if err != nil {
		return ctx.JSON(err.StatusCode, echo.Map{"error": err.Message})
	}
	return ctx.JSON(http.StatusOK, generated.UserResponse{
		FullName:    userResp.FullName,
		PhoneNumber: userResp.PhoneNumber,
	})
}

// (PUT /user)
func (s *Server) UserUpdate(ctx echo.Context) error {
	var req generated.UpdateUsers
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	accessToken := extractToken(ctx)
	err := s.Svc.UpdateUser(ctx.Request().Context(), req, accessToken)
	if err != nil {
		return ctx.JSON(err.StatusCode, echo.Map{"error": err.Message})
	}
	return ctx.JSON(http.StatusOK, generated.Users{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	})
}

func extractToken(c echo.Context) string {
	bearerToken := c.Request().Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
