package enums

type LotteryStatus string

var LotteryStatuses = struct {
	Active   LotteryStatus
	InActive LotteryStatus
	Deleted  LotteryStatus
}{
	Active:   LotteryStatusActive,
	InActive: LotteryStatusInActive,
	Deleted:  LotteryStatusDeleted,
}

const (
	LotteryStatusActive   LotteryStatus = "active"
	LotteryStatusInActive LotteryStatus = "inactive"
	LotteryStatusDeleted  LotteryStatus = "deleted"
)

func (s LotteryStatus) IsValid() bool {
	switch s {
	case LotteryStatusActive, LotteryStatusInActive, LotteryStatusDeleted:
		return true
	default:
		return false
	}
}

func (s LotteryStatus) String() string {
	switch s {
	case LotteryStatusActive:
		return "active"
	case LotteryStatusInActive:
		return "inactive"
	case LotteryStatusDeleted:
		return "deleted"
	default:
		return ""
	}
}

func (s LotteryStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
