package auth

import (
	"net/http"
	"os"
	"text/template"

	"github.com/coreos/go-oidc"
	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

var clientID = os.Getenv("GOOGLE_CLIENT_ID")
var config = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/api/v1/auth/oauth/google/callback",
	Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	Endpoint:     google.Endpoint,
}

func (auth *AuthHandler) HandleGooglePopup(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (auth *AuthHandler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	var claims = &GoogleOAuthClaims{}

	err := auth.Deps.AuthService.HandleGoogleOAuth(config, code, claims)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	user, err := auth.getOrCreateUserWithOAuth(claims)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	tokens, err := auth.signInGenerateTokens(user)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	tmpl := template.Must(template.New("callback").Parse(`
		<!DOCTYPE html>
		<html>
		<head><title>OAuth Callback</title></head>
		<body>
		<script>
			window.opener.postMessage({
				accessToken: "{{.AccessToken}}",
				accessTokenExpiresIn: {{.AccessTokenExpiresIn}},
				refreshToken: "{{.RefreshToken}}",
				refreshTokenExpiresIn: {{.RefreshTokenExpiresIn}}
			}, "http://localhost:4200");
			window.close();
		</script>
		</body>
		</html>
	`))

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tokens)
}

func (auth *AuthHandler) getOrCreateUserWithOAuth(claims *GoogleOAuthClaims) (*model.User, error) {
	userExists, err := auth.Deps.DB.UserExists(claims.Email)
	if err != nil {
		return nil, err
	}

	if !userExists {
		userId, err := auth.Deps.DB.CreateUserOAuth(claims.Email, claims.Name, "google", claims.Sub, claims.Picture, claims.EmailVerified)
		if err != nil {
			return nil, err
		}

		user, err := auth.Deps.DB.GetUserByID(userId)
		return user, err
	}

	user, err := auth.Deps.DB.GetUserByEmail(claims.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
