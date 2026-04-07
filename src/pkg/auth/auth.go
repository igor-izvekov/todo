package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "ВАШ_CLIENT_ID.apps.googleusercontent.com",
		ClientSecret: "ВАШ_CLIENT_SECRET",
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	state = generateStateOauthCookie()
)

func generateStateOauthCookie() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func HandleLogin(c *gin.Context) {
	url := oauthConf.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleCallback(c *gin.Context) {
	// Проверяем state для защиты от CSRF
	if c.Query("state") != state {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		log.Println("State mismatch")
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	ctx := context.Background()
	token, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		log.Printf("Token exchange error: %v", err)
		return
	}

	client := oauthConf.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		log.Printf("User info error: %v", err)
		return
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		log.Printf("Parse error: %v", err)
		return
	}

	log.Printf("User logged in: %s (%s)", userInfo.Name, userInfo.Email)

	c.HTML(http.StatusOK, "success.html", gin.H{
		"user": userInfo,
	})
}

func HandleHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Go OAuth2 with Google",
	})
}

func OpenBrowser(url string) {
	commands := []struct {
		cmd  string
		args []string
	}{
		{"cmd", []string{"/c", "start", "chrome", url}},
		{"cmd", []string{"/c", "start", url}},
		{"open", []string{"-a", "Google Chrome", url}},
		{"open", []string{url}},
		{"google-chrome", []string{url}},
		{"google-chrome-stable", []string{url}},
		{"chromium-browser", []string{url}},
		{"xdg-open", []string{url}},
	}

	for _, cmd := range commands {
		err := exec.Command(cmd.cmd, cmd.args...).Start()
		if err == nil {
			log.Printf("Browser opened with: %s %v", cmd.cmd, cmd.args)
			return
		}
	}
	log.Println("Could not open browser automatically, please visit http://localhost:8080")
}
