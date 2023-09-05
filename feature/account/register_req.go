package account

type RegisterReq struct {
	Username string `json:"username" validate:"max=16"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
