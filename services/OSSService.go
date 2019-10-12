package services

import (
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "storage-service/config"
    "strings"
    "mime/multipart"
    "io/ioutil"
)

func OSSUpload(file multipart.File, filename string) (string, error) {
    client, err := oss.New(config.Config.OSS_ENDPOINT, config.Config.OSS_ACCESS_KEY_ID, config.Config.OSS_ACCESS_KEY_SECRET)
    if err != nil {
        return "", err
    }

    bucket, err := client.Bucket(config.Config.OSS_BUCKET)
    if err != nil {
        return "", err
    }
    content, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }
    fileUri := config.Config.COS_BUCKET_DIR + "/" + MD5HashFileName(filename, content, "")
    f := strings.NewReader(string(content))
    err = bucket.PutObject(fileUri, f)
    if err != nil {
        return "", err
    }
    return fileUri, err
}

func OSSUploadNormal(file multipart.File, filename string) (string, error) {
    fileUri, err := OSSUpload(file, filename)
    if err != nil {
        return "", nil
    }
    if config.Config.OSS_HTTPS_ENABLE {
        return "https://" + config.Config.OSS_BUCKET + "." + config.Config.OSS_ENDPOINT + "/" + fileUri, nil
    }
    return "http://" + config.Config.OSS_BUCKET + "." + config.Config.OSS_ENDPOINT + "/" + fileUri, nil
}

func OSSUploadWithCDN(file multipart.File, filename string) (string, error) {
    fileUri, err := OSSUpload(file, filename)
    if err != nil {
        return "", nil
    }
    if config.Config.OSS_HTTPS_ENABLE {
        return "https://" + config.Config.OSS_CDN_DOMAIN + "/" + fileUri, nil
    }
    return "http://" + config.Config.OSS_CDN_DOMAIN + "/" + fileUri, nil
}