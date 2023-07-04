package app

type Service interface {
	Cast() interface{}
}

func (app *Application) GetService(name string) Service {
	return app.Services[name]
}
