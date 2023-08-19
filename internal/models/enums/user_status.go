package enums

type UserStatus string

var UserStatuses = struct {
	Active   UserStatus
	InActive UserStatus
	Banned   UserStatus
	Deleted  UserStatus
}{
	Active:   UserStatusActive,
	InActive: UserStatusInActive,
	Banned:   UserStatusBanned,
	Deleted:  UserStatusDeleted,
}

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInActive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
	UserStatusDeleted  UserStatus = "deleted"
)

func (s UserStatus) IsValid() bool {
	switch s {
	case UserStatusActive, UserStatusInActive, UserStatusBanned, UserStatusDeleted:
		return true
	default:
		return false
	}
}

func (s UserStatus) String() string {
	switch s {
	case UserStatusActive:
		return "active"
	case UserStatusInActive:
		return "inactive"
	case UserStatusBanned:
		return "banned"
	case UserStatusDeleted:
		return "deleted"
	default:
		return ""
	}
}

func (s UserStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
