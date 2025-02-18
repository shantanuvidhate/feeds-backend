package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shantanuvidhate/feeds-backend/internal/store"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user, err := app.store.User.GetById(ctx, userId)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}

	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerId, err := strconv.ParseInt(chi.URLParam(r, "followerId"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	followeeId, err := strconv.ParseInt(chi.URLParam(r, "followeeId"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if followerId == followeeId {
		app.badRequestResponse(w, r, errors.New("followerId and followeeId cannot be the same"))
		return
	}

	ctx := r.Context()
	err = app.store.User.Follow(ctx, followerId, followeeId)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()
		user, err := app.store.User.GetById(ctx, userId)
		if err != nil {
			switch err {
			case store.ErrRecordNotFound:
				app.notFoundResponse(w, r, err)
				return
			default:
				app.internalServerError(w, r, err)
				return
			}
		}
		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
