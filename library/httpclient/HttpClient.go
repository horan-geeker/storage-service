package httpclient

import (
    "net/http"
    "io/ioutil"
    "bytes"
    "time"
    "log"
)

var client = &http.Client{}

func Download(url string) (*http.Response, error) {
    return Get(url, nil)
}

func Get(url string, params []byte) (*http.Response, error) {
    var (
        response *http.Response
        err      error
    )
    before := time.Now()
    log.Println("http client request begin ", url)
    response, err = http.Get(url)
    if err != nil {
        return nil, err
    }
    log.Println("http client request end, time usage ", time.Now().Sub(before))
    return response, err
}

func Request(httpMethod string, url string, data []byte) ([]byte, int, error) {
    var (
        request  *http.Request
        response *http.Response
        err      error
        body     []byte
    )
    request, err = http.NewRequest(httpMethod, url, bytes.NewReader(data))
    if err != nil {
        log.Println(err)
        return nil, 0, err
    }
    request.Header.Set("Content-Type", "application/json")
    before := time.Now()
    log.Println("http client request begin ", url)
    response, err = client.Do(request)
    log.Println("http client request end, time usage ", time.Now().Sub(before))
    if err != nil {
        return nil, 0, err
    }
    body, err = ioutil.ReadAll(response.Body)
    defer response.Body.Close()
    return body, response.StatusCode, err
}
