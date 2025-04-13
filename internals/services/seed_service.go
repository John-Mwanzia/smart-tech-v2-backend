package services

import (
	"github.com/John-Mwanzia/smart-tech-v2-backend"
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/repository"
	"github.com/John-Mwanzia/smart-tech-v2-backend/pkg"
)

func SeedCollections() error {

  //users
  var users []interface{}

  for _, user := range data.InitialData.Users{
    hashedPassword, err := pkg.HashPassword(user.Password)
    if err != nil {
      return err
    }
    
    user.Password = hashedPassword
    users = append(users, user)
  }

  var products []interface{}

  for _, product := range data.InitialData.Products{
    products = append(products, product)
  }

  var fps []interface{}

  for _, fp := range data.InitialData.FeaturedProducts{
    fps = append(fps, fp)
  }



  //clear all collections 

  err := repository.ClearCollections()
  if err != nil {
    return err
  }

  //seed data 

  err = repository.SeedInitialdata(users, products, fps)
  if err != nil {
    return err
  }



  return nil
}
