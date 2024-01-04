package kanthorsdk

import (
	"errors"
	"strings"
)

func host(credentials string, proj *Project) (string, error) {
	segments := strings.Split(credentials, ":")
	if len(segments) != 2 || len(segments[1]) == 0 {
		return "", errors.New("malform credentials")
	}

	parts := strings.Split(segments[1], ".")
	if len(parts) == 2 {
		if h, exist := proj.Hosts[parts[0]]; exist {
			return h, nil
		}
	}

	return proj.Host, nil
}

func scheme(host string) string {
	if strings.HasPrefix(host, "localhost") {
		return "http"
	}
	if strings.HasPrefix(host, "127.0.0.1") {
		return "http"
	}
	return "https"
}

func max[T int | int32 | int64](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func min[T int | int32 | int64](x, y T) T {
	if x > y {
		return y
	}
	return x
}
