package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Uid       int    `json:"uid" gorm:"primaryKey"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	GID       int    `json:"gid"`
}

type Tabler interface {
	TableName() string
}

// Override default table name
func (User) TableName() string {
	return "user_"
}

func GetAllInfo(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		users := make([]User, 0)

		db.Table("user_").Find(&users)

		c.JSON(http.StatusOK, gin.H{"users": users})
	}

	return gin.HandlerFunc(fn)
}

// Get user by user ID
func GetUserById(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")

		var u User

		db.Model(&User{}).Select("user_.username, user_.firstname, user_.lastname, user_.email, group_.gid").Joins("inner join group_ on user_.gid = group_.gid").Where("user_.uid = ?", id).Scan(&u)

		c.IndentedJSON(http.StatusOK, u)
	}

	return gin.HandlerFunc(fn)
}
