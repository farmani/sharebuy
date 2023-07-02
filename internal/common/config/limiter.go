package config

// Add a new limiter struct containing fields for the request-per-second and burst
// values, and a boolean field which we can use to enable/disable rate limiting.
type Limiter struct {
	Rps     float64
	Burst   int
	Enabled bool
}
