package api

//gorm实践
import (
	"_/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Mygorm struct{}
type User struct {
	ID int64 `json:"id"`
	//Name         string `json:"name"`
	//Email        *string `json:"email"`
	//Age          uint8 `json:"age"`
	Birthday utils.NullTime `json:"birthday"`
	//MemberNumber utils.NullString `json:"member_number"`
	//Test sql.NullString `json:"test"`
	//ActivatedAt  sql.NullTime `json:"activated_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g Mygorm) Migration(c *gin.Context) {
	db.AutoMigrate(&User{})
}

func (g Mygorm) Create(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, err.Error())
		return
	}
	fmt.Println(user)

	//if err:=db.Create(&user).Error;err!=nil{
	//	c.JSON(400,err.Error())
	//	return
	//}
	db.Model(&User{}).Where("id = ?", 13).Update("name", "hello")
	c.JSON(200, user)
	return
}
