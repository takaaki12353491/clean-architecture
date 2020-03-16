package handler

import (
	"cln-arch/errs"
	"net/http"
)

func statusCode(err error) int {
	switch errs.GetType(err) {
	case errs.Invalidated:
		return http.StatusBadRequest
	case errs.Forbidden:
		return http.StatusForbidden
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Conflict:
		return http.StatusConflict
	case errs.Failed:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
