package orm

import (
	"highfaner.com/utils"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func TestUsers(t *testing.T) {
	db, err := gorm.Open("mysql", "root:root@tcp(:3306)/highfaner?collation=utf8mb4_general_ci&parseTime=true&loc=Local")
	if err != nil {
		t.Error(err)
		return
	}

	db.LogMode(true)

	seller := Seller{
		Username: "shen19911",
		Password: utils.MD5CreateStrings("123456"),
		// State:    true,
		// Remark:   "创建成功",
	}

	if err := seller.Find(db); err != nil {
		t.Error(err)
		return
	}

	t.Log(seller)

	// if err := seller.Create(db); err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// t.Log(seller)
}
