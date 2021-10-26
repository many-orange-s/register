package model

type Sign struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Department string `json:"department" binding:"required"`
}
