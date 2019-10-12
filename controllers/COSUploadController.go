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
    url, err := services.COSUploadNormal(file, filename,  "")
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
    url, err := services.COSUploadSecure(file, filename, "")
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
    url, err := services.COSUploadWithCDN(file, filename, "")
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "cdn_url": url,
    }))
}

func COSUploadByUrl(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    fileUrls, ok := params["file_url"]
    if !ok || len(fileUrls[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'file_url' is missing"))
        return
    }
    fileContent, fileType, err := services.DownloadByUrl(fileUrls[0])
    if err != nil {
        io.WriteString(w, response.JsonError(response.FILE_URL_DOWNLOAD_ERR))
        return
    }
    url, err := services.COSUploadFileContentNormal(fileContent, "", fileType)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "url": url,
    }))
}

func COSUploadByUrlWithCDN(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    fileUrls, ok := params["file_url"]
    if !ok || len(fileUrls[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'file_url' is missing"))
        return
    }
    fileContent, fileType, err := services.DownloadByUrl(fileUrls[0])
    if err != nil {
        io.WriteString(w, response.JsonError(response.FILE_URL_DOWNLOAD_ERR))
        return
    }
    url, err := services.COSUploadFileContentWithCDN(fileContent, "", fileType)
    if err != nil {
        log.Println(err.Error())
    }
    io.WriteString(w, response.Json(response.SUCCESS, "success", map[string]interface{}{
        "cdn_url": url,
    }))
}
