package model

type User struct {
	//gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	//當這個 struct 被轉換為 JSON 時，對應的 JSON 鍵名應該是 "username"。
	//如果沒加，就會使用原生的name，大小寫敏感
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsername(user *User) string {
	return user.Username
}
