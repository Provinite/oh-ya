package serverutils

import (
	"net/http"
	"strconv"
)

func GetPagination(req *http.Request) (skip, limit int, err error) {
	skipStr := req.URL.Query().Get("skip")
	limitStr := req.URL.Query().Get("limit")

	if skipStr == "" {
		skipStr = "0"
	}
	if limitStr == "" {
		limitStr = "100"
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, err
	}
	skip, err = strconv.Atoi(skipStr)
	if err != nil {
		return 0, 0, err
	}
	return skip, limit, err
}
