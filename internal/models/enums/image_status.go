package enums

type ImageStatus string

var ImageStatuses = struct {
	Active    ImageStatus
	Processed ImageStatus
	Uploaded  ImageStatus
	Ready     ImageStatus
	Deleted   ImageStatus
}{
	Active:    ImageStatusPending,
	Processed: ImageStatusProcessed,
	Uploaded:  ImageStatusUploaded,
	Ready:     ImageStatusReady,
	Deleted:   ImageStatusDeleted,
}

const (
	ImageStatusPending   ImageStatus = "pending"
	ImageStatusProcessed ImageStatus = "processed"
	ImageStatusUploaded  ImageStatus = "uploaded"
	ImageStatusReady     ImageStatus = "ready"
	ImageStatusDeleted   ImageStatus = "deleted"
)

func (s ImageStatus) IsValid() bool {
	switch s {
	case ImageStatusPending, ImageStatusProcessed, ImageStatusUploaded, ImageStatusReady, ImageStatusDeleted:
		return true
	default:
		return false
	}
}

func (s ImageStatus) String() string {
	switch s {
	case ImageStatusPending:
		return "pending"
	case ImageStatusProcessed:
		return "processed"
	case ImageStatusUploaded:
		return "uploaded"
	case ImageStatusReady:
		return "ready"
	case ImageStatusDeleted:
		return "deleted"
	default:
		return ""
	}
}

func (s ImageStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}
