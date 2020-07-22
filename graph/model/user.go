package model

type User struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	PremiumID   string       `json:"premium_id"`
	Premiumtype *Premiumtype `json:"premiumtype"`
}
