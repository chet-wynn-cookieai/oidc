package op

import (
	"context"
	"time"

	"gopkg.in/square/go-jose.v2"

	"github.com/caos/oidc/pkg/oidc"
)

type AuthStorage interface {
	CreateAuthRequest(context.Context, *oidc.AuthRequest, string) (AuthRequest, error)
	AuthRequestByID(context.Context, string) (AuthRequest, error)
	AuthRequestByCode(context.Context, string) (AuthRequest, error)
	SaveAuthCode(context.Context, string, string) error
	DeleteAuthRequest(context.Context, string) error

	CreateToken(context.Context, TokenRequest) (string, time.Time, error)

	TerminateSession(context.Context, string, string) error

	GetSigningKey(context.Context, chan<- jose.SigningKey, chan<- error, <-chan time.Time)
	GetKeySet(context.Context) (*jose.JSONWebKeySet, error)
	SaveNewKeyPair(context.Context) error
}

type OPStorage interface {
	GetClientByClientID(ctx context.Context, clientID string) (Client, error)
	AuthorizeClientIDSecret(ctx context.Context, clientID, clientSecret string) error
	SetUserinfoFromScopes(ctx context.Context, userinfo oidc.UserInfoSetter, userID, clientID string, scopes []string) error
	SetUserinfoFromToken(ctx context.Context, userinfo oidc.UserInfoSetter, tokenID, subject, origin string) error
	GetPrivateClaimsFromScopes(ctx context.Context, userID, clientID string, scopes []string) (map[string]interface{}, error)
	GetKeyByIDAndUserID(ctx context.Context, keyID, userID string) (*jose.JSONWebKey, error)
	//ValidateJWTProfileScopes(ctx context.Context, userID string, scope oidc.Scopes) (oidc.Scopes, error)

	//deprecated: use GetUserinfoFromScopes instead
	GetUserinfoFromScopes(ctx context.Context, userID, clientID string, scopes []string) (oidc.UserInfo, error)
	//deprecated: use SetUserinfoFromToken instead
	GetUserinfoFromToken(ctx context.Context, tokenID, subject, origin string) (oidc.UserInfo, error)
}

type Storage interface {
	AuthStorage
	OPStorage
	Health(context.Context) error
}

type StorageNotFoundError interface {
	IsNotFound()
}

type EndSessionRequest struct {
	UserID      string
	Client      Client
	RedirectURI string
}
