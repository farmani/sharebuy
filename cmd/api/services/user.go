package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/internal/data"
	"github.com/farmani/sharebuy/internal/dto"
	"github.com/farmani/sharebuy/internal/validator"
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

func (s *UserService) RegisterStart(u dto.User) error {
	// do something
	// s.app.Db.QueryRow("SELECT 1")

	user := &data.User{
		Email:     u.Email,
		Activated: false,
	}

	// Use the Password.Set() method to generate and store the hashed and plaintext
	// passwords.
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	// Validate the user struct and return the error messages to the client if
	// any of the checks fail.
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Insert the user data into the database.
	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		// If we get an ErrDuplicateEmail error, use the v.AddError() method to manually add
		// a message to the validator instance, and then call our failedValidationResponse
		// helper().
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Add the "movies:read" permission for the new user.
	err = app.models.Permissions.AddForUser(user.ID, "movies:read")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// After the user record has been created in the database, generate a new activation
	// token for the user.
	token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Launch a goroutine which runs an anonymous function that sends the welcome email using
	// the background helper function.
	app.background(func() {
		// Create map to act as a 'holding structure' for the data we send to the weclome email
		// template.
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}

		// Call the Send() method on our Mailer, passing in the user's email address, name of the
		// template file, and the data map containing the activationToken and the user's ID.
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			// Importantly, if there is an error sending the email then we log the error
			// instead of raising a server error like before when we handled
			// the email send functionality without a goroutine
			app.logger.PrintError(err, nil)
		}
	})

	// Note that we also change this to send the client a 202 Accepted status code which
	// indicates that the request has been accepted for processing, but the processing has
	// not been completed.
	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
