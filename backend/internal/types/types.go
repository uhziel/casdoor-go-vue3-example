// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RenewAfter   int64  `json:"renewAfter"`
}

type UserDetailResp struct {
	Name string `json:"name"`
}