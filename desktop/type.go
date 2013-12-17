package desktop

// Type is the type of desktop file.
type Type uint8

const (
	Unknown Type = iota
	Application
	Link
	Directory
)

// ParseType converts the given string s into a Type.
func ParseType(s string) Type {
	switch s {
	case Application.String():
		return Application
	case Link.String():
		return Link
	case Directory.String():
		return Directory
	default:
		return Unknown
	}
}

// String returns the Type as a string.
func (t Type) String() string {
	switch t {
	case Application:
		return "Application"
	case Link:
		return "Link"
	case Directory:
		return "Directory"
	default:
		return "Unknown"
	}
}
