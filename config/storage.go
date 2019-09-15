package config

import (
    "github.com/jinzhu/configor"
    "github.com/joho/godotenv"
)

var Config = struct {
    APP_ENV     string `default:"local"`
    COS_SECRET_ID  string
    COS_SECRET_KEY string
    COS_BUCKET_DOMAIN string
    COS_BUCKET_DIR string
}{}

func init() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
    configor.Load(&Config)
}