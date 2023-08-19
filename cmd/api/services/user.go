package services

import (
	"errors"
	"github.com/farmani/sharebuy/internal/models/enums"
	"github.com/farmani/sharebuy/pkg/cookie"
	"github.com/farmani/sharebuy/pkg/encryption"
	"github.com/farmani/sharebuy/pkg/jwt"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"time"

	"github.com/farmani/sharebuy/internal/models"
	"github.com/farmani/sharebuy/internal/repository"

	"github.com/farmani/sharebuy/cmd/api/app"
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

func (us *UserService) Cast() interface{} {
	return us
}

func (us *UserService) FindById(id int64) (*models.User, error) {
	// do something
	// s.app.Db.QueryRow("SELECT 1")
	if id == 0 {
		return &models.User{}, repository.ErrRecordNotFound
	}

	us.app.Logger.Info("User service called")
	u, err := us.app.Repository.User().GetById(id)
	if err != nil {
		return &models.User{}, err
	}

	return u, nil
}

func (us *UserService) RegisterStart(u *models.User) error {
	u.Status = enums.UserStatusInActive
	err := us.app.Repository.User().Insert(u)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEmail) {
			us.app.Background(func() {
				data := map[string]interface{}{
					"userID": u.ID,
				}

				err = us.app.Mailer.Send(u.Email, "user_welcome.tmpl", data)
				if err != nil {
					us.app.Logger.Error(ErrSendingEmail.Error())
				}
			})
			return nil
		}

		return err
	}

	token, err := us.app.Repository.Token().GenerateToken(u.ID, 30*time.Minute, enums.TokenScopeActivation)
	if err != nil {
		return err
	}
	err = us.app.Repository.Token().Insert(token)
	if err != nil {
		return err
	}

	us.app.Background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          u.ID,
		}

		err = us.app.Mailer.Send(u.Email, "user_welcome.tmpl", data)
		if err != nil {
			us.app.Logger.Error(ErrSendingEmail.Error())
		}
	})
	return nil
}

func (us *UserService) ActivateUser(hash string, userId int64) (*models.User, error) {
	token, err := us.app.Repository.Token().GetByHash(hash, userId, enums.TokenScopeActivation)
	if err != nil {
		return nil, err
	}

	user, err := us.app.Repository.User().GetById(userId)
	if err != nil {
		return nil, err
	}
	user.Status = enums.UserStatusActive
	err = us.app.Repository.User().Update(user)
	if err != nil {
		return nil, err
	}

	us.app.Background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}

		err = us.app.Mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			us.app.Logger.Error(ErrSendingEmail.Error())
		}
	})

	return user, nil
}

func (us *UserService) Authorize(ectx echo.Context, user *models.User) (string, error) {
	jwtToken, err := jwt.New(us.app.Config.Jwt, us.app.Redis)
	if err != nil {
		return "", err
	}

	mapClaims := golangJwt.MapClaims{
		"email": user.Email,
		"id":    12,
	}
	token, err := jwtToken.GenerateToken(mapClaims)
	if err != nil {
		return "", err
	}

	enc := encryption.New(us.app.Config.Encryption)

	sess, _ := session.Get("session", ectx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	err = sess.Save(ectx.Request(), ectx.Response())
	if err != nil {
		return "", err
	}

	co := cookie.New(us.app.Config.Cookie)
	err = co.SetEncryptedCookies(ectx, enc, us.app.Config.Jwt.CookieTokenName, token)
	if err != nil {
		return "", err
	}

	return token, nil
}
