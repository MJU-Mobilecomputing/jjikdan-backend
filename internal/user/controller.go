package user

import (
	"errors"
	"net/http"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService interfaces.IUserService
}

func InitUserController() *UserController {
	return &UserController{}
}

func (u *UserController) SignupUserController(ctx echo.Context) error {
	if param, ok := ctx.Get("userParam").(repository.CreateUserParams); ok {
		user, err := u.UserService.Create(param)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, user)
	} else {
		return ctx.JSON(http.StatusInternalServerError, errors.New("param not provided"))
	}
}

func (u UserController) WithUserService(service interfaces.IUserService) UserController {
	u.UserService = service
	return u
}
