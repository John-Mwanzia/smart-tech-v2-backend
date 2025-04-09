package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Env struct {
  jwt_secret                string
  password                  string
  mongodb_url               string
  safaricom_consumer_secret string
  safaricom_consumer_key    string
  pass_key                  string
  business_short_code       string
  port                      string
  stripe_cancel_url         string
  stripe_success_url        string
  stripe_secret_key         string
  jwtsecret                 string
}

var AppEnv Env

func LoadEnv() {
  err := godotenv.Load("../../env") 

  if err != nil {
    log.Fatal("Error trying to load env file", err)
  }


  AppEnv = Env{
    jwt_secret: os.Getenv("JWT_SECRET"),
    password : os.Getenv("PASSWORD"),
    mongodb_url: os.Getenv("MONGODB_URL"),
    safaricom_consumer_secret: os.Getenv("SAFARICOM_CONSUMER_SECRET"),
    safaricom_consumer_key: os.Getenv("SAFARICOM_CONSUMER_KEY"),
    pass_key: os.Getenv("PASS_KEY"),
    business_short_code: os.Getenv("BUSINESS_SHORT_CODE"),
    port: os.Getenv("PORT"),
    stripe_cancel_url: os.Getenv("STRIPE_CANCEL_URL"),
    stripe_success_url: os.Getenv("STRIPE_SUCCESS_URL"),
    stripe_secret_key: os.Getenv("STRIPE_SECRET_KEY"),
    jwtsecret: os.Getenv("JWTSECRET"),
  }
}
