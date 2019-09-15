package controllers

import (
    "net/http"
    "storage-service/services"
    "log"
    "io"
    "storage-service/library/response"
)

func Upload(w http.ResponseWriter, r *http.Request) {
    file, header , err := r.FormFile("file")
    if err != nil {
        log.Println(err.Error())
        io.WriteString(w, response.JsonParamError("formData Param 'file' is missing"))
        return
    }
    filename := header.Filename
    url, err := services.COSUpload(file, filename)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "url": url,
    }))
}

func UploadWithCDN(w http.ResponseWriter, r *http.Request) {

}