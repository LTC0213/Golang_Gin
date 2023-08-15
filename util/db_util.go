package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() error {
	dsn := "root:Nhss@3120@tcp(192.168.1.213:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"
	// 替换上述 dsn 中的参数为实际的数据库连接信息

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func CloseDB() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

//func main() {
//	ConnectDB()
//}
