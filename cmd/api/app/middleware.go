package app

import (
	"expvar"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/farmani/sharebuy/pkg/logger"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/sessions"
	echoPrometheus "github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (app *Application) bundleMiddleware(e *echo.Echo) {
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	zapLogger := logger.NewZapLogger(app.Config.Logger.Path, app.Config.App.Env)
	e.Use(logger.ZapLogger(zapLogger))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("token"))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics") // Change "metrics" for your own path
		},
	}))
	// Enable metrics middleware
	e.Use(echoPrometheus.NewMiddleware("Sharebuy"))

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, nil)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(config))

	e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "upload")
		},
		Limit: "2M",
	}))

	/*
		e.Use(echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte(app.Config.Jwt.PrivatePem),
			Skipper: func(c echo.Context) bool {
				return strings.Contains(c.Path(), "upload")
			},
			SigningMethod: "EdDSA",
			TokenLookup:   "header:Authorization:Bearer,cookie:__Secure_token",
			// BeforeFunc middleware.BeforeFunc
			// SuccessHandler func(c echo.Context)
			// ErrorHandler func(c echo.Context, err error) error
			// ContinueOnIgnoredError bool
			// SigningKey interface{}
			// SigningKeys map[string]interface{}
			// TokenLookupFuncs []middleware.ValuesExtractor
			// ParseTokenFunc jwtToken.ExtractToken func(c echo.Context, auth string) (interface{}, error)
			NewClaimsFunc: func(c echo.Context) golangJwt.Claims {
				return new(jwt.Claims)
			},
		}))
	*/
	// e.Validator = NewValidator()

}

// enableCORS sets the Vary: Origin and Access-Control-Allow-Origin response headers in order to
// enabled CORS for trusted origins.
func (app *Application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Access-Control-Request-Method")
		origin := r.Header.Get("Origin")

		if origin != "" {
			for i := range app.Config.Cors.TrustedOrigins {
				if origin == app.Config.Cors.TrustedOrigins[i] {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
						w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
						w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

						w.Header().Set("Access-Control-Max-Age", "60")

						w.WriteHeader(http.StatusOK)
						return
					}

					break
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) metrics(next http.Handler) http.Handler {
	// Initialize the new expvar variables when middleware chain is first build.
	totalRequestsReceived := expvar.NewInt("total_requests_received")
	totalResponsesSent := expvar.NewInt("total_responses_sent")
	totalProcessingTimeMicroseconds := expvar.NewInt("total_processing_time_µs")
	totalResponsesSentbyStatus := expvar.NewMap("total_responses_sent_by_status")

	// Below runs for every request.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// use the Add method to increment the number of requests received by 1.
		totalRequestsReceived.Add(1)

		// Call the httpsnoop.CaptureMetrics function, passing in the next handler in the chain
		// along with the existing http.ResponseWriter and http.Request. This returns the metrics
		// struct.
		metrics := httpsnoop.CaptureMetrics(next, w, r)

		// On way back up middleware chain, increment the number of responses sent by 1.
		totalResponsesSent.Add(1)

		// Get the request processing time in microseconds from httpsnoop and increment the
		// cumulative processing time.
		totalProcessingTimeMicroseconds.Add(metrics.Duration.Microseconds())

		// / Use the Add method to increment the count for the given status code by 1.
		// Note, the expvar map is string-keyed, so we need to use the strconv.Itoa
		// function to convert the status (an integer) to a string.
		totalResponsesSentbyStatus.Add(strconv.Itoa(metrics.Code), 1)
	})
}

func JWTErrorChecker(err error, c echo.Context) error {
	// Redirects to the signIn form.
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
}
