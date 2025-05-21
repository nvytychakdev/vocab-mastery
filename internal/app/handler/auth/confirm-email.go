package auth

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
	"gopkg.in/mail.v2"
)

type ConfirmEmailRequest struct {
	Token string `json:"token"`
}

func (s *ConfirmEmailRequest) Bind(r *http.Request) error {
	if s.Token == "" {
		return errors.New("token is missing")
	}
	return nil
}

type ConfirmEmailResponse struct {
	AccessToken           string     `json:"accessToken"`
	AccessTokenExpiresIn  int64      `json:"accessTokenExpiresIn"`
	RefreshToken          string     `json:"refreshToken"`
	RefreshTokenExpiresIn int64      `json:"refreshTokenExpiresIn"`
	User                  SignInUser `json:"user"`
}

func (s *ConfirmEmailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (auth *AuthHandler) ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	data := &ConfirmEmailRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInvalidPayload, nil))
		return
	}

	userId, usedAt, err := auth.Deps.DB.GetNonExpiredUserToken(data.Token, model.EMAIL_CONFIRM_TOKEN)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrConfirmTokenExpired, err))
		return
	}

	if usedAt != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusUnauthorized, httpError.ErrConfirmTokenAlreadyUsed, err))
		return
	}

	err = auth.Deps.DB.SetUserTokenUsed(data.Token)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	err = auth.Deps.DB.SetUserEmailConfirmed(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	user, err := auth.Deps.DB.GetUserByID(userId)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	auth.signInComplete(w, r, user)
}

func (auth *AuthHandler) sendEmailConfirm(w http.ResponseWriter, r *http.Request, userId string, email string) {
	_, token, err := auth.Deps.DB.CreateUserToken(userId, model.EMAIL_CONFIRM_TOKEN)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	err = auth.sendEmailConfirmMessage(email, token)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}
	response := &EmailConfirmResponse{
		Sent: true,
	}

	render.Render(w, r, response)
}

type EmailTemplateData struct {
	UserName string
	Link     string
}

func (auth *AuthHandler) sendEmailConfirmMessage(email string, token string) error {
	tmpl, err := template.ParseFiles("templates/email-confirm-tpl.html")
	if err != nil {
		slog.Error("Erorr", "err", err)
		return err
	}

	link := fmt.Sprintf("http://localhost:4200/auth/confirm-email?token=%s", token)
	data := &EmailTemplateData{
		UserName: "Test",
		Link:     link,
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return err
	}

	htmlBody := buf.String()
	// Create a new message
	message := mail.NewMessage()

	// Set email headers
	message.SetHeader("From", "no-reply@vocab-mastery.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Email Confirmation")

	// Set the plain-text version of the email
	message.SetBody("text/plain", "This is a Test Email\n\nHello!\nThis is a test email with plain-text formatting.\nThanks,\nMailtrap")

	// Set the HTML version of the email
	message.AddAlternative("text/html", htmlBody)

	// Set up the SMTP dialer
	dialer := mail.NewDialer("sandbox.smtp.mailtrap.io", 587, os.Getenv("MAILTRAP_USER"), os.Getenv("MAILTRAP_PASSWORD"))

	// Send the email
	return dialer.DialAndSend(message)
}
