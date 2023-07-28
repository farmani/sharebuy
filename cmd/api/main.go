package api

import (
	"github.com/farmani/sharebuy/internal/repository"
	"os"

	"github.com/farmani/sharebuy/cmd/api/app"
	"github.com/farmani/sharebuy/cmd/api/handlers"
	"github.com/farmani/sharebuy/cmd/api/services"
	cfg "github.com/farmani/sharebuy/internal/config"
	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var version = "1.0.0"

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(cfg.Load(true))
	}

	return &cobra.Command{
		Use:   "api-v1",
		Short: `Start the API server v1`,
		Long:  "Start the API server v1",
		Run:   run,
	}
}

func (cmd Server) run(cfg *cfg.Config) {
	zapLogger := logger.NewZap(cfg.Logger)
	field := zap.String("version", version)
	zapLogger.Info("Starting API server", field)

	apiApp := app.NewApiApplication(cfg)
	repo := repository.New(zapLogger, cfg.Repository, apiApp.Db)

	// Add the handlers to the Application
	apiApp.Handlers = []app.Handler{
		handlers.NewSiteHandler(apiApp, services.NewSiteService(apiApp, repo)),
		handlers.NewUserHandler(apiApp, services.NewUserService(apiApp, repo)),
		// Add other handlers here
	}

	if err := apiApp.Start(); err != nil {
		panic(err)
	}
}

func main() {
	Server{}.run(cfg.Load(true))
}
