package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/internal/dto"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	// Dependencies or state for the UserHandler
	app *app.Application
}

func NewUserService(a *app.Application) *UserService {
	return &UserService{
		app: a,
	}
}

func (s *UserService) Cast() interface{} {
	return s
}

func (s *UserService) FindById(c echo.Context) error {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	return nil
}

func (s *UserService) RegisterStart(u dto.User) data.Models.Users, err {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	// Insert the user data into the database.
	user, err := data.Models.Users.Insert(u)
	if err != nil {
		if errors.Is(err, data.ErrDuplicateEmail) {
			err.AddError("email", "a user with this email address already exists")
		}

		return nil, err
	}

	// Add the "movies:read" permission for the new user.
	err = data.Models.Permissions.AddForUser(user.ID, "movies:read")
	if err != nil {
		return err
	}

	// After the user record has been created in the database, generate a new activation
	// token for the user.
	token, err := data.Models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		return nil, err
	}

	s.app.Background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}

		err = s.app.Mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			s.app.Logger.PrintError(err, nil)
		}
	})

	return user, nil
}
