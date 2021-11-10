package models

import (
	"devsmake/models"
	"devsmake/util"
)

type Users []User

type User struct {
	ID               uint64 `json:"id"`
	ProviderID       uint64 `json:"provider_id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Points           uint64 `json:"points"`
	TotalPosts       uint64 `json:"total_posts"`
	TotalComments    uint64 `json:"total_comments"`
	TotalSubmissinos uint64 `json:"total_submissions"`
	TotalRatings     uint64 `json:"total_ratings"`
	Created          string `json:"created"`
}

func (u *User) PublicUser() *models.Profile {
	return &models.Profile{
		Identifier:       int64(u.ID),
		Username:         u.Username,
		Points:           int64(u.Points),
		TotalPosts:       int64(u.TotalPosts),
		TotalComments:    int64(u.TotalComments),
		TotalRatings:     int64(u.TotalRatings),
		TotalSubmissions: int64(u.TotalSubmissinos),
		AvatarURL:        util.GenerateAvatarUrl(u.Username),
	}
}
func (users Users) PublicUsers() []*models.Profile {
	result := make([]*models.Profile, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}
