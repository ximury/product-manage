// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Gender   int64  `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type UserInfoResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}

type GetUseResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}
