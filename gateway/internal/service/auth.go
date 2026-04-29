package service

import (
	"context"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/SynKolbasyn/bank/gateway/internal/repository"
	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Auth struct {
	repositoryUser repository.IUser
	secretKey []byte
}

func NewAuth(repositoryUser repository.IUser, secretKey []byte) *Auth {
	return &Auth{
		repositoryUser: repositoryUser,
		secretKey: secretKey,
	}
}

func (a *Auth) SignUp(ctx context.Context, user *model.SignRequest) (string, error) {
	passwordHash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	
	id, err := a.repositoryUser.Create(ctx, user.Email, passwordHash);
	if err != nil {
		return "", err
	}
	
	return a.generateToken(id)
}

func (a *Auth) SignIn(ctx context.Context, user *model.SignRequest) (string, error) {
	userID, passwordHash, err := a.repositoryUser.Get(ctx, user.Email)
	if err != nil {
		return "", err
	}

	match, err := argon2id.ComparePasswordAndHash(user.Password, passwordHash)
	if err != nil {
		return "", domain.NewAppError(http.StatusInternalServerError, err)
	}

	if !match {
		return "", domain.NewAppError(http.StatusBadRequest)
	}
	
	return a.generateToken(userID);	
}

func (a *Auth) generateToken(userID uuid.UUID) (string, error) {
	claims := model.JWTAuthData{
		UserID: userID, 
	}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString(a.secretKey)

    if err != nil {
        return "", err
    }
    
    return signed, nil
}
