package han

import (
	"database/sql/driver"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Phone    *Phone `json:"phone" gorm:"type:json"`
}

type Phone struct {
	Number string `json:"number"`
	Manu   string `json:"manu"`
}

var Sql *gorm.DB

func Tes(c *gin.Context) {
	data := new(User)
	err := c.ShouldBindJSON(data)
	if err != nil {
		log.Printf("binding error: %s", err)
	}
	Sql.Migrator().CreateTable(&User{})
	Sql.Create(data)
	op := Sql.First(&data)
	log.Println("OP:", op)
	c.JSON(200, op)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *Phone) Scan(value interface{}) error {
	log.Println("Scan value:", value)
	err := json.Unmarshal(value.([]byte), j)
	log.Println("Scan err:", err)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j Phone) Value() (driver.Value, error) {
	data, err := json.Marshal(j)
	if err != nil {
		log.Println("value error")
	}
	log.Println("Value", string(data))
	return string(data), nil
}

// func Phoness(s validator.FieldLevel) bool {
// 	if date, ok := s.Field().Interface().(string); ok {
// 		if len(date) > 5 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Val() {
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		err := v.RegisterValidation("phs", Phoness)
// 		if err != nil {
// 			fmt.Println("success")
// 		}
// 	}
// }
