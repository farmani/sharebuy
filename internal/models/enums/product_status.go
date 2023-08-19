package enums

type ProductStatus string

var ProductStatuses = struct {
	Active   ProductStatus
	InActive ProductStatus
	Deleted  ProductStatus
}{
	Active:   ProductStatusActive,
	InActive: ProductStatusInActive,
	Deleted:  ProductStatusDeleted,
}

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusInActive ProductStatus = "inactive"
	ProductStatusDeleted  ProductStatus = "deleted"
)

func (s ProductStatus) IsValid() bool {
	switch s {
	case ProductStatusActive, ProductStatusInActive, ProductStatusDeleted:
		return true
	default:
		return false
	}
}

func (s ProductStatus) String() string {
	switch s {
	case ProductStatusActive:
		return "active"
	case ProductStatusInActive:
		return "inactive"
	case ProductStatusDeleted:
		return "deleted"
	default:
		return ""
	}
}

func (s ProductStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
