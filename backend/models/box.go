package models

import "time"

type BoxUserResponse struct {
	Type              string        `json:"type"`
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Login             string        `json:"login"`
	CreatedAt         time.Time     `json:"created_at"`
	ModifiedAt        time.Time     `json:"modified_at"`
	Language          string        `json:"language"`
	Timezone          string        `json:"timezone"`
	SpaceAmount       int64         `json:"space_amount"`
	SpaceUsed         int           `json:"space_used"`
	MaxUploadSize     int64         `json:"max_upload_size"`
	Status            string        `json:"status"`
	JobTitle          string        `json:"job_title"`
	Phone             string        `json:"phone"`
	Address           string        `json:"address"`
	AvatarURL         string        `json:"avatar_url"`
	NotificationEmail []interface{} `json:"notification_email"`
}
