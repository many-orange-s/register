package model

// Sign 登录结构体
type Sign struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Password string `json:"password" binding:"required" db:"password"`
}

// Register 注册结构体
type Register struct {
	Name       string `json:"name" binding:"required" db:"name"`
	Password   string `json:"password" binding:"required" db:"password"`
	Department string `json:"department" binding:"required" db:"department"`
}

// AllMsg 成员插入结构体
// 注意都要打*号 可以防止空的时候无法传递
// 用户输入是不用输入id
type AllMsg struct {
	Id        *int    `json:"id" db:"id"`
	Name      *string `json:"name" db:"name"`
	Gender    *string `json:"gender" db:"gender"`
	Grade     *string `json:"grade" db:"grade"`
	Birth     *string `json:"birth" db:"birth"`
	Telephone *string `json:"telephone" db:"telephone"`
	GroupName *string `json:"group_name" db:"group_name"`
}

// Update 更新信息时的结构体
type Update struct {
	Target     string `json:"target" binding:"required"`
	UpdateData string `json:"update_data" binding:"required"`
}
