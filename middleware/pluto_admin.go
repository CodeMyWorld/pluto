package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/leeif/pluto/utils/general"

	"github.com/urfave/negroni"

	"github.com/gorilla/context"
	perror "github.com/leeif/pluto/datatype/pluto_error"
	"github.com/leeif/pluto/utils/jwt"
	"github.com/wxnacy/wgo/arrays"
)

func PlutoAdmin() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		auth := strings.Fields(r.Header.Get("Authorization"))

		if len(auth) < 2 && strings.ToLower(auth[0]) != "jwt" {
			context.Set(r, "pluto_error", perror.Forbidden)
			next(w, r)
			return
		}

		jwtToken, perr := jwt.VerifyJWT(auth[1])
		if perr != nil {
			context.Set(r, "pluto_error", perr)
			next(w, r)
			return
		}

		accessPayload := &jwt.AccessPayload{}

		if err := json.Unmarshal(jwtToken.Payload, &accessPayload); err != nil {
			context.Set(r, "pluto_error", perror.ServerError.Wrapper(err))
			next(w, r)
			return
		}

		if accessPayload.Type != jwt.ACCESS {
			context.Set(r, "pluto_error", perror.InvalidJWTToekn)
			next(w, r)
			return
		}

		if time.Now().Unix() > accessPayload.Expire {
			context.Set(r, "pluto_error", perror.JWTTokenExpired)
			next(w, r)
			return
		}

		if accessPayload.AppID != general.PlutoAdminApplication {
			context.Set(r, "pluto_error", perror.Forbidden)
			next(w, r)
			return
		}

		if arrays.ContainsString(accessPayload.Scopes, general.PlutoAdminScope) == -1 {
			context.Set(r, "pluto_error", perror.NotPlutoAdmin)
			next(w, r)
			return
		}

		context.Set(r, "payload", accessPayload)

		next(w, r)
	}
}