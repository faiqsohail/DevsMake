package models

type User struct {
	ID         uint64 `json:"id"`
	ProviderID uint64 `json:"provider_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Points     uint64 `json:"points"`
	Created    string `json:"created"`
}
