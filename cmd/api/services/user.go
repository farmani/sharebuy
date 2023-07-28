package services

import (
	"github.com/farmani/sharebuy/internal/repository"
	"time"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/internal/data"
)

type UserService struct {
	// Dependencies or state for the UserHandler
	app  *app.Application
	repo repository.Repository
}

func NewUserService(a *app.Application, repo repository.Repository) *UserService {
	return &UserService{
		app:  a,
		repo: repo,
	}
}

func (s *UserService) Cast() interface{} {
	return s
}

func (s *UserService) FindById(id int64) (data.User, error) {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	if id == 0 {
		return data.User{}, data.ErrRecordNotFound
	}

	return data.User{
		ID:        id,
		Name:      "Farmani",
		Email:     "ramin.farmani@gmail.com",
		CreatedAt: time.Now(),
		Activated: true,
		Version:   1,
	}, nil
}

/*
func (s *UserService) RegisterStart(u dto.User) (data.Users, err) {
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
				app.Logger.Error(err.(string))
		}
	})

	return user, nil
}
*/
