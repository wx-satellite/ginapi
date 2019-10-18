package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	Byn *gorm.DB
}

var DB *Database

func Init() {
	DB = &Database{}
	DB.InitByn()
}

func Close() {
	DB.CloseByn()
}

func (db *Database) InitByn() {
	db.Byn = GetDB("byn")
}
func (db *Database) CloseByn() {
	var (
		err error
	)
	if err = db.Byn.Close(); err != nil {
		fmt.Println("必应鸟数据库关闭失败！")
		panic(err)
	}
}


func  GetDB(t string) *gorm.DB {
	var (
		env string
	)
	env = viper.GetString("env")
	username := viper.GetString(fmt.Sprintf("mysql.%s.%s.username", t, env))
	password := viper.GetString(fmt.Sprintf("mysql.%s.%s.password", t, env))
	address := viper.GetString(fmt.Sprintf("mysql.%s.%s.address", t, env))
	fmt.Println(address)
	database := viper.GetString(fmt.Sprintf("mysql.%s,%s.database", t,env))
	return openDB(username, password,address,database)
}

func openDB(username, password, address, name string) *gorm.DB {
	var (
		config string
		db *gorm.DB
		err error
	)
	// parseTime不加的话，模型的时间解析会失败（	storing driver.Value type []uint8 into type *time.Time ）
	config = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		username,
		password,
		address,
		name,
	)
	db, err = gorm.Open("mysql", config)
	if err != nil {
		panic(err)
	}

	// 设置数据库
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {

	//db.DB().SetMaxOpenConns(100); 设置最大连接数
	//db.DB().SetMaxIdleConns(10) // 设置最大空闲连接数

	if viper.GetString("env") == "local" {
		db.LogMode(true)
	}
}
