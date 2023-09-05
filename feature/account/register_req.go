package account

type RegisterReq struct {
	Username string `json:"username" validate:"max=3"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
