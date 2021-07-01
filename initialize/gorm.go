package initialize

import (
	"github.com/TualatinX/durian-go/global"
	"github.com/TualatinX/durian-go/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQL() {
	// host, port, username, password, database := FetchMySQLConfig()
	ADDR:=os.Getenv("ADDR")
	USR:=os.Getenv("USR")
	PWD:=os.Getenv("PWD")
	DB:=os.Getenv("DB")
	dsn := USR + ":" + PWD + "@tcp(" + ADDR + ")/" + DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
	)
}
