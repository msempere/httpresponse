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

func getResponseReasonPhrase(code int) string {
    switch code {
    case 100:
        return "Continue"
    case 101:
        return "Switching Protocols"
    case 200:
        return "OK"
    case 201:
        return "Created"
    case 202:
        return "Accepted"
    case 203:
        return "Non-Authoritative Information"
    case 204:
        return "No Content"
    case 205:
        return "Reset Content"
    case 206:
        return "Partial Content"
    case 300:
        return "Multiple Choices"
    case 301:
        return "Moved Permanently"
    case 302:
        return "Found"
    case 303:
        return "See Other"
    case 304:
        return "Not Modified"
    case 305:
        return "Use Proxy"
    case 307:
        return "Temporary Redirect"
    case 400:
        return "Bad Request"
    case 401:
        return "Unauthorized"
    case 402:
        return "Payment Required"
    case 403:
        return "Forbidden"
    case 404:
        return "Not Found"
    case 405:
        return "Method Not Allowed"
    case 406:
        return "Not Acceptable"
    case 407:
        return "Proxy Authentication Required"
    case 408:
        return "Request Time-out"
    case 409:
        return "Conflict"
    case 410:
        return "Gone"
    case 411:
        return "Length Required"
    case 412:
        return "Precondition Failed"
    case 413:
        return "Request Entity Too Large"
    case 414:
        return "Request-URI Too Large"
    case 415:
        return "Unsupported Media Type"
    case 416:
        return "Requested range not satisfiable"
    case 417:
        return "Expectation Failed"
    case 500:
        return "Internal Server Error"
    case 501:
        return "Not Implemented"
    case 502:
        return "Bad Gateway"
    case 503:
        return "Service Unavailable"
    case 504:
        return "Gateway Time-out"
    case 505:
        return "HTTP Version not supported"
    default:
        return "unknown response code"
    }
}

func NewHttpResponse(responseCode int, contentType, body string) *HttpResponse {
    p := new(HttpResponse)
    p.responseCode = responseCode
    p.responseStatus = getResponseReasonPhrase(responseCode)
    p.body = body
    p.contentType = contentType
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
