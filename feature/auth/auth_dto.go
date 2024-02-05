package auth

type RegisterReq struct {
	Username string `json:"username" validate:"min=3,max=16"`
	Password string `json:"password" validate:"min=6,max=16"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	About    string `json:"about"`
	Role     string `json:"role"`
	Mobile   string `json:"mobile"`
}

type RegisterResp struct {
	ID       string
	Username string
	Password string
	Email    string
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccountId    string `json:"accountId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
