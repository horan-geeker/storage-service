package main

import (
    "net/http"
    "storage-service/controllers"
    "log"
)

func main() {
    // tencent cos storage
    http.HandleFunc("/v1/cos/upload", controllers.COSUpload)
    http.HandleFunc("/v1/cos/upload/secure", controllers.COSUploadSecure)
    http.HandleFunc("/v1/cos/upload/cdn", controllers.COSUploadWithCDN)
    // aliyun oss storage
    http.HandleFunc("/v1/oss/upload", controllers.OSSUpload)
    http.HandleFunc("/v1/oss/upload/secure", controllers.OSSUploadSecure)
    http.HandleFunc("/v1/oss/upload/cdn", controllers.OSSUploadWithCDN)

    log.Println("server start")
    log.Fatal(http.ListenAndServe(":80", nil))
}
