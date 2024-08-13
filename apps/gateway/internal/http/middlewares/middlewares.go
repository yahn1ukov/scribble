package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/yahn1ukov/scribble/apps/gateway/internal/config"
	"github.com/yahn1ukov/scribble/libs/jwt"
	"github.com/yahn1ukov/scribble/libs/respond"
)

const USER_ID_KEY = "id"

type Middleware struct {
	cfg *config.Config
}

func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{
		cfg: cfg,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			respond.Error(w, http.StatusUnauthorized, ErrInvalidFormat)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := jwt.Validate(token, m.cfg.JWT.Secret)
		if err != nil {
			if errors.Is(err, jwt.ErrInvalidToken) {
				respond.Error(w, http.StatusUnauthorized, jwt.ErrInvalidToken)
				return
			}

			if _, err = uuid.Parse(claims.UserID); err != nil {
				respond.Error(w, http.StatusUnauthorized, jwt.ErrInvalidToken)
				return
			}

			respond.Error(w, http.StatusForbidden, ErrAccessDenied)
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, claims.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *Middleware) GetUserIDFromCtx(ctx context.Context) string {
	userID, ok := ctx.Value(USER_ID_KEY).(string)
	if !ok {
		return ""
	}

	return userID
}
