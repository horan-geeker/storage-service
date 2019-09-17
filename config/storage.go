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
    COS_CDN_DOMAIN string

    OSS_BUCKET string
    OSS_BUCKET_DIR string
    OSS_HTTPS_ENABLE bool
    OSS_ENDPOINT string
    OSS_ACCESS_KEY_ID string
    OSS_ACCESS_KEY_SECRET string
    OSS_CDN_DOMAIN string
}{}

func init() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
    configor.Load(&Config)
}