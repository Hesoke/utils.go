package ensure

import (
	"fmt"
	"net/http"
)

type StatusCodes map[int]string

// i was bored
var requestFailedCodes = StatusCodes{
	http.StatusBadRequest:                   "bad request: 400",
	http.StatusUnauthorized:                 "unauthorized request: 401",
	http.StatusPaymentRequired:              "payment required: 402",
	http.StatusForbidden:                    "forbidden: 403",
	http.StatusNotFound:                     "not found: 404",
	http.StatusMethodNotAllowed:             "method not allowed: 405",
	http.StatusNotAcceptable:                "not acceptable: 406",
	http.StatusProxyAuthRequired:            "proxy authentication required: 407",
	http.StatusRequestTimeout:               "request timeout: 408",
	http.StatusConflict:                     "conflict: 409",
	http.StatusGone:                         "gone: 410",
	http.StatusLengthRequired:               "content-length required: 411",
	http.StatusPreconditionFailed:           "precondition failed: 412",
	http.StatusRequestEntityTooLarge:        "payload too large: 413",
	http.StatusRequestURITooLong:            "uri too long: 414",
	http.StatusUnsupportedMediaType:         "unsupported media type: 415",
	http.StatusRequestedRangeNotSatisfiable: "range not satisdiable: 416",
	http.StatusExpectationFailed:            "expectation failed: 417",
	http.StatusTeapot:                       "but i am a teapot :D",
	http.StatusMisdirectedRequest:           "misdirected request: 421",
	http.StatusUnprocessableEntity:          "unprocessable content: 422",
	http.StatusLocked:                       "locked: 423",
	http.StatusFailedDependency:             "failed dependency: 424",
	http.StatusTooEarly:                     "too early: 425",
	http.StatusUpgradeRequired:              "upgrade required: 426",
	http.StatusPreconditionRequired:         "precondition required: 428",
	http.StatusTooManyRequests:              "too many requests: 429",
	http.StatusRequestHeaderFieldsTooLarge:  "request header too large: 431",
	http.StatusUnavailableForLegalReasons:   "unavailable for legal reasons: 451",
}

// like really bored
var ServerErrorCodes = StatusCodes{
	http.StatusInternalServerError:           "internal server error: 500",
	http.StatusNotImplemented:                "not implemented: 501",
	http.StatusBadGateway:                    "bad gateway: 502",
	http.StatusServiceUnavailable:            "service unavailable: 503",
	http.StatusGatewayTimeout:                "gateway timeout: 504",
	http.StatusHTTPVersionNotSupported:       "http version not supported: 505",
	http.StatusVariantAlsoNegotiates:         "variant also negotiates: 506",
	http.StatusInsufficientStorage:           "insufficient storage: 507",
	http.StatusLoopDetected:                  "loop detected: 508",
	http.StatusNetworkAuthenticationRequired: "network authentication required",
}

func StatusGood(code int) error {
	if msg, ok := requestFailedCodes[code]; ok {
		return fmt.Errorf(msg)
	}
	if msg, ok := ServerErrorCodes[code]; ok {
		return fmt.Errorf(msg)
	}
	return nil
}
