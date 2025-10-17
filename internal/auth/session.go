package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

const sessionCookieName = "blog_admin_session"
const sessionSecretLength = 32

// initSessionKey generates a secure key for session encryption. (Not strictly needed for this simple demo,
// but kept for future expansion of secured cookies).
func InitSessionKey() {
	// In a real application, you would use this key to encrypt the session cookie value.
	// For this simple demo, we just generate a random token and rely on its presence.
}

// SetAuthCookie sets a cookie marking the user as authenticated.
func SetAuthCookie(w http.ResponseWriter) {
	token := make([]byte, 16)
	rand.Read(token)
	encodedToken := hex.EncodeToString(token)

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    encodedToken,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour), // Session lasts 1 hour
		HttpOnly: true,
		Secure:   false, // Set to true in a production HTTPS environment
	})
}

// ClearAuthCookie clears the session cookie (logout).
func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
		HttpOnly: true,
	})
}

// IsAuthenticated checks for the presence of the admin session cookie.
func IsAuthenticated(r *http.Request) bool {
	_, err := r.Cookie(sessionCookieName)
	return err == nil
}

// RequireAuth is a middleware function for protecting admin routes.
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}