package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	vt "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/truemail-rb/truemail-go"
)

var (
	// EmailRX is a regex for sanity checking the format of email addresses.
	// The regex pattern used is taken from  https://html.spec.whatwg.org/#valid-e-mail-address.
	EmailRX = `([-!#-'*+/-9=?A-Z^-~]+(\.[-!#-'*+/-9=?A-Z^-~]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?(\.[0-9A-Za-z]([0-9A-Za-z-]{0,61}[0-9A-Za-z])?)*|\[((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|IPv6:((((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){6}|::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){5}|[0-9A-Fa-f]{0,4}::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){4}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):)?(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){3}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,2}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){2}|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,3}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,4}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::)((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3})|(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3})|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,5}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3})|(((0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}):){0,6}(0|[1-9A-Fa-f][0-9A-Fa-f]{0,3}))?::)|(?!IPv6:)[0-9A-Za-z-]*[0-9A-Za-z]:[!-Z^-~]+)])`
)

// Validator struct type contains a map of validation errors.
type Validator struct {
	Errors     map[string]string
	validator  *vt.Validate
	translator ut.Translator
}

// New is a helper which creates a new Validator instance with an empty errors map.
func New(c echo.Context) *Validator {
	v := vt.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.FindTranslator(c.Request().Header.Get("Accept-Language"), "en")

	v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "'{0}' is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe vt.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "'{0}' is not a valid email address", true) // see universal-translator for details
	}, func(ut ut.Translator, fe vt.FieldError) string {
		t, _ := ut.T("email", fmt.Sprintf("%v", fe.Value()))

		return t
	})
	v.RegisterTranslation("email_dns", trans, func(ut ut.Translator) error {
		return ut.Add("email_dns", "'{0}' is not a valid email address", true) // see universal-translator for details
	}, func(ut ut.Translator, fe vt.FieldError) string {
		t, _ := ut.T("email_dns", fmt.Sprintf("%v", fe.Value()))

		return t
	})

	_ = v.RegisterValidation("email_dns", EmailDnsValidation)

	return &Validator{Errors: make(map[string]string), validator: v, translator: trans}
}

func (v *Validator) Validate(r interface{}) {
	err := v.validator.Struct(r)
	if err != nil {
		for _, err := range err.(vt.ValidationErrors) {
			v.AddError(err.Field(), err.Translate(v.translator))
		}
	}
}

// Valid returns true if the errors map doesn't contain any entries.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for the
// given key).
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of strings.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}

func EmailDnsValidation(fl vt.FieldLevel) bool {
	field := fl.Field()

	configuration, err := truemail.NewConfiguration(
		truemail.ConfigurationAttr{
			WhitelistedDomains: []string{"sharebuy.com"},
			VerifierEmail:      "ramin.farmani@gmail.com",
		},
	)

	if err != nil {
		return false
	}

	return truemail.IsValid(field.String(), configuration)
}
