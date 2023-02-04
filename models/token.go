package models

import "time"

const (
	InviteToken                = "invite"
	SessionToken               = "session"
	InviteTokenExpiryDuration  = time.Hour           // One hour
	SessionTokenExpiryDuration = time.Hour * 24 * 14 // 14 days
)

type Token struct {
	ID        string    `bson:"_id"`
	OwnerID   string    `bson:"ownerId"`
	Type      string    `bson:"type"`
	ExpiresAt time.Time `bson:"expiresAt"`
}
