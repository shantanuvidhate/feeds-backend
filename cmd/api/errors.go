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
