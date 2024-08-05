package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contextKey string

var contextUserKey contextKey

// returns the value stored in context
func (h *handler) ipFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

func (h *handler) addIpToContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userIP, err := getIP(ctx.Request)
		if err != nil {
			IP, _, _ := net.SplitHostPort(ctx.Request.RemoteAddr)
			if len(IP) == 0 {
				IP = "unknown"
			}
			ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), contextUserKey, IP))
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), contextUserKey, userIP))

		ctx.Next()
	}
}

func getIP(r *http.Request) (string, error) {
	IP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", err
	}

	// validating if the given ip  is a valid ip address
	userIP := net.ParseIP(IP)
	if userIP == nil {
		return "", fmt.Errorf("userIP: %q is not a ip:port", r.RemoteAddr)
	}

	// check if request is coming thruogh a proxy server
	proxyUserip := r.Header.Get("X-Forwarded-For")
	if len(proxyUserip) > 0 {
		IP = proxyUserip
	}

	if len(IP) == 0 {
		IP = "unknown"
	}

	return IP, nil
}
