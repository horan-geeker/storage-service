package services

import (
    "mime/multipart"
    "net/http"
    "context"
    "net/url"
    "io/ioutil"
    "encoding/json"
    "github.com/tencentyun/cos-go-sdk-v5"
    "log"
    "strings"
    "storage-service/config"
    "crypto/md5"
    "encoding/hex"
    "time"
)

type COSUploadResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data struct {
        AccessUrl string `json:"access_url"`
        PreviewUrl string `json:"preview_url"`
        ResourcePath string `json:"resource_path"`
        SourceUrl string `json:"source_url"`
        Url string `json:"url"`
    }
}

func MD5HashFileName(filename string, fileContent []byte) string {
    hasher := md5.New()
    hasher.Write(fileContent)
    md5String := hex.EncodeToString(hasher.Sum(nil))
    fileFullName := strings.Split(filename, ".")
    if len(fileFullName) >= 2 {
        filename = md5String + "." + fileFullName[len(fileFullName)-1]
    }
    return filename
}

func COSUpload(file multipart.File, filename string) (string, error) {
    u, _ := url.Parse(config.Config.COS_BUCKET_DOMAIN)
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  config.Config.COS_SECRET_ID,
            SecretKey: config.Config.COS_SECRET_KEY,
        },
    })
    content, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }
    fileUrl := config.Config.COS_BUCKET_DIR + "/" + MD5HashFileName(filename, content)
    f := strings.NewReader(string(content))
    httpResponse, err := client.Object.Put(context.Background(), fileUrl, f, nil)
    if err != nil {
        return "", err
    }
    body, err := ioutil.ReadAll(httpResponse.Body)
    defer httpResponse.Body.Close()
    log.Println(string(body), httpResponse.StatusCode)
    uploadResponse := COSUploadResponse{}
    json.Unmarshal(body, &uploadResponse)
    if uploadResponse.Code != 0 {
        return "", err
    }
    presignedURL, err := client.Object.GetPresignedURL(context.Background(), http.MethodGet, fileUrl, config.Config.COS_SECRET_ID, config.Config.COS_SECRET_KEY, time.Hour, nil)
    if err != nil {
        log.Println(err.Error())
    }
    return presignedURL.String(), nil
}