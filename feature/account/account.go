package account

type Account struct {
	ID     string
	Name   string
	About  string
	Role   Role
	Mobile string
	AuthID string
}

type Role string

const (
	ROLE_ADMINISTRATOR Role = "administrator"
	ROLE_EDITOR        Role = "editor"
	ROLE_USER          Role = "user"
	ROLE_UNKNOWN       Role = "unknown"
)

func MapStringToRole(roleStr string) Role {
	switch roleStr {
	case "administrator":
		return ROLE_ADMINISTRATOR
	case "editor":
		return ROLE_EDITOR
	case "user":
		return ROLE_USER
	default:
		return ROLE_UNKNOWN
	}
}
