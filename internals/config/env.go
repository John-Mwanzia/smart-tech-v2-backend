package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Env struct {
  Jwt_secret                string
  Password                  string
  Mongodb_url               string
  Safaricom_consumer_secret string
  Safaricom_consumer_key    string
  Pass_key                  string
  Business_short_code       string
  Port                      string
  Stripe_cancel_url         string
  Stripe_success_url        string
  Stripe_secret_key         string
  Jwtsecret                 string
}

var AppEnv Env

func LoadEnv() {
  err := godotenv.Load(".env") 

  if err != nil {
    log.Fatal("Error trying to load env file", err)
  }


  AppEnv = Env{
    Jwt_secret: os.Getenv("JWT_SECRET"),
    Password : os.Getenv("PASSWORD"),
    Mongodb_url: os.Getenv("MONGODB_URL"),
    Safaricom_consumer_secret: os.Getenv("SAFARICOM_CONSUMER_SECRET"),
    Safaricom_consumer_key: os.Getenv("SAFARICOM_CONSUMER_KEY"),
    Pass_key: os.Getenv("PASS_KEY"),
    Business_short_code: os.Getenv("BUSINESS_SHORT_CODE"),
    Port: os.Getenv("PORT"),
    Stripe_cancel_url: os.Getenv("STRIPE_CANCEL_URL"),
    Stripe_success_url: os.Getenv("STRIPE_SUCCESS_URL"),
    Stripe_secret_key: os.Getenv("STRIPE_SECRET_KEY"),
    Jwtsecret: os.Getenv("JWTSECRET"),
  }
}
