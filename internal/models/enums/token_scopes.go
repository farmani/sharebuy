package enums

type TokenScope string

var TokenScopes = struct {
	Activation         TokenScope
	Authentication     TokenScope
	ForgetPassword     TokenScope
	ChangePayoutMethod TokenScope
}{
	Activation:         TokenScopeActivation,
	Authentication:     TokenScopeAuthentication,
	ForgetPassword:     TokenScopeForgetPassword,
	ChangePayoutMethod: TokenScopeChangePayoutMethod,
}

const (
	TokenScopeActivation         TokenScope = "activation"
	TokenScopeAuthentication     TokenScope = "authentication"
	TokenScopeForgetPassword     TokenScope = "forget_password"
	TokenScopeChangePayoutMethod TokenScope = "change_payout_method"
)

func (t TokenScope) IsValid() bool {
	switch t {
	case TokenScopeActivation, TokenScopeAuthentication, TokenScopeForgetPassword, TokenScopeChangePayoutMethod:
		return true
	default:
		return false
	}
}

func (t TokenScope) String() string {
	switch t {
	case TokenScopeActivation:
		return "activation"
	case TokenScopeAuthentication:
		return "authentication"
	case TokenScopeForgetPassword:
		return "forget_password"
	case TokenScopeChangePayoutMethod:
		return "change_payout_method"
	default:
		return ""
	}
}

func (t TokenScope) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}
