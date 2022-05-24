package auth0

import (
	"email-masks-service/src/business/entities"
	"github.com/auth0/go-auth0/management"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func mockManagementClient() *management.Management {
	body := struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: "token"}

	gock.New("https://domain.com").
		Post("/oauth/token").
		Reply(200).
		JSON(body)

	m, _ := management.New("domain.com", management.WithClientCredentials("id", "secret"))

	return m
}

func TestAuth0UsersService_GetUserByID(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		defer gock.Off()

		m := mockManagementClient()
		auth0UsersService := NewAuth0UsersService(m)

		userID := "longUserId"

		gock.New("https:/domain.com").
			Get("/api/v2/users/longUserId").
			Reply(404)

		_, err := auth0UsersService.GetUserByID(userID)

		if !gock.IsDone() {
			for _, re := range gock.GetUnmatchedRequests() {
				t.Logf("Failed to match %s", re.URL)
			}
			t.Fail()
			return
		}

		if err == nil {
			t.Error("Expected to return an error")
		}
	})

	t.Run("success", func(t *testing.T) {
		defer gock.Off()

		m := mockManagementClient()
		auth0UsersService := NewAuth0UsersService(m)

		userID := "longUserId"
		email := "userEmail"
		auth0User := management.User{
			ID:    &userID,
			Email: &email,
		}

		gock.New("https:/domain.com").
			Get("/api/v2/users/longUserId").
			Reply(200).
			JSON(auth0User)

		user, err := auth0UsersService.GetUserByID(userID)

		if !gock.IsDone() {
			for _, re := range gock.GetUnmatchedRequests() {
				t.Logf("Failed to match %s", re.URL)
			}
			t.Fail()
			return
		}

		if err != nil {
			t.Errorf("Expected not to return an error, given %v", err)
		}

		assert.Equal(t, user, &entities.User{
			ID:    userID,
			Email: email,
		})
	})
}
