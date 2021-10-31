package models

import (
	"devsmake/models"
	"devsmake/util"
)

type Users []User

type User struct {
	ID         uint64 `json:"id"`
	ProviderID uint64 `json:"provider_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Points     uint64 `json:"points"`
	Created    string `json:"created"`
}

func (u *User) PublicUser() *models.Profile {
	return &models.Profile{
		Identifier: int64(u.ID),
		Username:   u.Username,
		Points:     int64(u.Points),
		AvatarURL:  util.GenerateAvatarUrl(u.Username),
	}
}
func (users Users) PublicUsers() []*models.Profile {
	result := make([]*models.Profile, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}
