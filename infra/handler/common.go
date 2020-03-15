package handler

import (
	"cln-arch/errs"
	"net/http"
)

func statusCode(err error) int {
	switch errs.GetType(err) {
	case errs.Invalidated:
		return http.StatusBadRequest
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Forbidden:
		return http.StatusForbidden
	case errs.Failed:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
