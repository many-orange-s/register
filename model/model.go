package model

type Sign struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Password string `json:"password" binding:"required" db:"password"`
}

type Register struct {
	Name       string `json:"name" binding:"required" db:"name"`
	Password   string `json:"password" binding:"required" db:"password"`
	Department string `json:"department" binding:"required" db:"department"`
}

type AllMsg struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Gender    string `json:"gender" db:"gender"`
	Grade     string `json:"grade" db:"grade"`
	Birth     string `json:"birth" db:"birth"`
	Telephone string `json:"telephone" db:"telephone"`
	GroupName string `json:"group_name" db:"group_name"`
}

type Update struct {
	Target     string `json:"target"`
	UpdateDate string `json:"update_date"`
}
