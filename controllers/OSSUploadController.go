package controllers

import (
    "net/http"
    "log"
    "io"
    "storage-service/services"
    "storage-service/library/response"
)

func OSSUpload(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.OSSUploadNormal(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "url": url,
    }))
}

func OSSUploadSecure(w http.ResponseWriter, r *http.Request) {

}

func OSSUploadWithCDN(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.OSSUploadWithCDN(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "url": url,
    }))
}