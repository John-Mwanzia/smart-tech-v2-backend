package pkg

import (
	"time"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
	"github.com/golang-jwt/jwt"
)



func GenerateJWT(user models.User)(string, error){
  //create claims

  claims := jwt.MapClaims{
    "id": user.ID,
    "name": user.Name,
    "email": user.Email,
    "isAdmin": user.IsAdmin,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  //generate token
  
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  //sign the token

  signedToken, err := token.SignedString([]byte(config.AppEnv.Jwtsecret))
  if err != nil {
    return "", err
  }

  return signedToken, nil 

}




