package actio

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ReverseProxy is the reverse proxy for the API
func ReverseProxy(w http.ResponseWriter, r *http.Request, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("URL parse error: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	// ReverseProxy
	proxy.ServeHTTP(w, r)
}

// CookieValue : handles getting cookie value and returning base of key
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

// GetLanguage : returns language from header
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
