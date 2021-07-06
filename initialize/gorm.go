package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TualatinX/durian-go/global"
	"github.com/TualatinX/durian-go/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 初始化 MySQL 的相关配置
func InitMySQL() {
	addr, port, username, password, database := "rm-uf6ji600qianqe6921o.mysql.rds.aliyuncs.com", "3306", "buaase2021", "buaase(2021)", "durian"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addr, port, database)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
		&model.Subscription{},
		&model.News{},
		&model.Notice{},
		&model.Rumor{},
		&model.Question{},
		&model.Comment{},
		&model.HighRiskArea{},
		&model.DirectData{},
		&model.CovidCases{},
		&model.CovidDeaths{},
		&model.CovidRecovered{},
	)
}
