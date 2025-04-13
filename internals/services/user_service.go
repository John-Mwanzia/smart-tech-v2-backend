package services

import (
	"fmt"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/repository"
	"github.com/John-Mwanzia/smart-tech-v2-backend/pkg"
)

func RegisterUser(user models.User) (*models.User, error) {
  //check if user already exists

  _, err := repository.FindUserByEmail(user.Email)

  if err == nil {
    return nil, fmt.Errorf("user already exists %w: ", err)
  }

  //hash passowrd
  hashedPassword, err := pkg.HashPassword(user.Password)

  if err != nil {
    return nil, fmt.Errorf("failed to hash password %w", err)
  }
  user.Password = hashedPassword
  
  //create user 
  createdUser, err := repository.CreateUser(user)
  if err != nil {
    return nil, err
  }
	return createdUser, nil  
}


func SignInUser(user models.User)(*models.User, error){
  //check if user exists
  existingUser, err := repository.FindUserByEmail(user.Email)

  if err != nil {
    return nil, fmt.Errorf("user not found ")
  }

  //validate password

  if !pkg.CheckPassword(existingUser.Password, user.Password) {
    return nil, fmt.Errorf("wrong credentials") 
  }

  return existingUser , nil 

}





















  
