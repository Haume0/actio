package actio

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ReverseProxy handles HTTP requests by forwarding them to a target server.
// It takes an http.ResponseWriter, an http.Request, and a target URL as parameters.
// The function parses the target URL and creates a reverse proxy to forward the request.
// If the URL parsing fails, it logs a fatal error.
func ReverseProxy(w http.ResponseWriter, r *http.Request, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("URL parse error: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	// ReverseProxy
	proxy.ServeHTTP(w, r)
}

// CookieValue retrieves the value of a specified cookie from the HTTP request.
// If the cookie is not found or an error occurs, it returns an empty string.
//
// Example Usage:
//
//	val := CookieValue(r, "cookieName")
func CookieValue(r *http.Request, key string) string {
	cookie, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	val, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}
	return val
}

// GetLanguage retrieves the preferred language from the HTTP request.
// It first checks for the "accept-language" cookie, and if not found,
// it falls back to the "Accept-Language" header.
//
// Example Usage:
//
//	lang := GetLanguage(r)
func GetLanguage(r *http.Request) string {
	//check cookie
	cookie := CookieValue(r, "accept-language")
	if cookie != "" {
		return cookie
	}
	//check header
	language := r.Header.Get("Accept-Language")
	var lang = strings.Split(language, ",")[0] //get first language
	if strings.Contains(lang, "-") {
		lang = strings.Split(lang, "-")[0] //get first language
	}
	return lang
}
