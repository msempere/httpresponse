package httpresponse

import (
    "strconv"
)

type HttpResponse struct {
    responseCode   int
    responseStatus string
    contentType    string
    body           string
    extraHeaders   []string
}

func NewHttpResponse(responseCode int, responseStatus, body string) *HttpResponse {
    p := new(HttpResponse)
    p.responseCode = responseCode
    p.responseStatus = responseStatus
    p.body = body
    p.contentType = "application/json"
    //p.extraHeaders = extraHeaders
    return p
}

func (h *HttpResponse) Get() string {
    responseStr := "HTTP/1.1 "
    responseStr += strconv.Itoa(h.responseCode) + " " + h.responseStatus + "\r\n"

    if h.contentType != "" {
        responseStr += "Content-Type: " + h.contentType + "\r\n"
    }
    responseStr += "Content-Length: " + strconv.Itoa(len(h.body)) + "\r\nConnection: Keep-Alive\r\n" + h.body
    return responseStr
}
