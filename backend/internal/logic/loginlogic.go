package logic

import (
	"context"
	"fmt"
	"net/http"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	url          = "http://192.168.31.64:8000/api/login/oauth/access_token"
	clientID     = "2db6dfeba919f743247a"
	clientSecret = "28a245520ef717e005acc9e1c4ba2ef4c5533869"
)

type TokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type TokenResp struct {
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	tokenReq := TokenReq{
		GrantType:    "authorization_code",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         req.Code,
	}

	httpResp, err := httpc.Do(l.ctx, http.MethodPost, url, tokenReq)
	if err != nil {
		return
	}

	if httpResp.StatusCode >= 400 {
		err = fmt.Errorf("Login fail statusCode=%d", httpResp.StatusCode)
		return
	}

	tokenResp := TokenResp{}
	if err = httpc.Parse(httpResp, &tokenResp); err != nil {
		return
	}

	return &types.LoginResp{
		AccessToken: tokenResp.AccessToken,
	}, nil
}
