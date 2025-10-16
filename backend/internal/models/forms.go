package models


// LoginMetadata is sent to /login handler
type LoginMetadata struct {
	UserAgent string
	IP string
}
