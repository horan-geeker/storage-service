package controllers

import (
    "net/http"
    "storage-service/services"
    "log"
    "io"
    "storage-service/library/response"
)

func COSUpload(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.COSUploadNormal(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "url": url,
    }))
}

func COSUploadSecure(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.COSUploadSecure(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "signed_url": url,
    }))
}

func COSUploadWithCDN(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.COSUploadWithCDN(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "cdn_url": url,
    }))
}