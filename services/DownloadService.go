package services

import (
    "storage-service/library/httpclient"
    "io/ioutil"
    "strings"
)

func DownloadByUrl(url string) ([]byte, string, error) {
    response, err := httpclient.Download(url)
    if (err != nil || response.StatusCode != 200) {
        return nil, "", err
    }
    body, err := ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    contentType := response.Header.Get("Content-Type")
    fileType := ""
    if len(contentType) != 0 {
        contentTypes := strings.Split(contentType, "/")
        if len(contentTypes) >= 2 {
            fileType = contentTypes[len(contentTypes) - 1]
        }
    }
    return body, fileType, nil;
}