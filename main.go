package main

import (
    "net/http"
    "storage-service/controllers"
    "log"
)

func main() {
    http.HandleFunc("/v1/cos/upload", controllers.Upload)
    http.HandleFunc("/v1/cos/upload/cdn", controllers.UploadWithCDN)
    log.Println("server start")
    log.Fatal(http.ListenAndServe(":80", nil))
}
