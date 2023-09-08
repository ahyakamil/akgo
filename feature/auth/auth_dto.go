package auth

type RegisterReq struct {
	Username string `json:"username" validate:"min=3,max=16"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
