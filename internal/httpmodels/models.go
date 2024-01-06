package httpmodels

import "github.com/DanilaNik/IU5_RIP2023/internal/service/role"

type Item struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
	Quantity uint64 `json:"quantity"`
	Height   uint64 `json:"height"`
	Width    uint64 `json:"width"`
	Depth    uint64 `json:"depth"`
	Barcode  uint64 `json:"barcode"`
}

type User struct {
	ID       uint64    `json:"id"`
	UserName string    `json:"userName"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     role.Role `json:"role"`
	ImageURL string    `json:"image_url"`
}