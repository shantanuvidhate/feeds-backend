package main

import (
	"net/http"

	"go.uber.org/zap"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Error("internal server error",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem!")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warn("bad request error",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))
	writeJSONError(w, http.StatusBadRequest, error.Error(err))
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warn("not found error",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))
	writeJSONError(w, http.StatusNotFound, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("resourse already exists: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Error("resource already exists",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))
	writeJSONError(w, http.StatusConflict, "resource already exists")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("unauthorized error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warn("unauthorized basic error",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("unauthorized error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warn("unauthorized error",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", err.Error()))
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	// log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	app.logger.Warn("forbidden",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("error", "forbidden"),
	)
	writeJSONError(w, http.StatusForbidden, "forbidden")
}
