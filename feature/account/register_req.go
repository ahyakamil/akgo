package account

type RegisterReq struct {
	Username string `json:"username" validate:"min=3,max=16"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
