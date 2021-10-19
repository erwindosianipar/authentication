package infra

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	Config() *viper.Viper
	SetMode() string
	GormDB() *gorm.DB
	Migrate(values ...interface{})
	Port() string
}

type infra struct {
	configFile string
}

func New(configFile string) Infra {
	return &infra{configFile: configFile}
}

var (
	vprOnce sync.Once
	vpr     *viper.Viper
)

func (i *infra) Config() *viper.Viper {
	vprOnce.Do(func() {
		viper.SetConfigFile(i.configFile)
		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatalf("[infra][Config][viper.ReadInConfig] %v", err)
		}

		vpr = viper.GetViper()
	})

	return vpr
}

var (
	modeOnce    sync.Once
	mode        string
	development = "development"
	production  = "production"
)

func (i *infra) SetMode() string {
	modeOnce.Do(func() {
		env := i.Config().Sub("environment").GetString("mode")
		if env == development {
			mode = gin.DebugMode
		} else if env == production {
			mode = gin.ReleaseMode
		} else {
			logrus.Fatalf("[infa][SetMode] %v", errors.New("environment not setup"))
		}

		gin.SetMode(mode)
	})

	return mode
}

var (
	grmOnce sync.Once
	grm     *gorm.DB
)

func (i *infra) GormDB() *gorm.DB {
	grmOnce.Do(func() {
		config := i.Config().Sub("database")
		user := config.GetString("user")
		pass := config.GetString("pass")
		host := config.GetString("host")
		port := config.GetString("port")
		name := config.GetString("name")

		dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			logrus.Fatalf("[infra][GormDB][gorm.Open] %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			logrus.Fatalf("[infra][GormDB][db.DB] %v", err)
		}

		if err := sqlDB.Ping(); err != nil {
			logrus.Fatalf("[infra][GormDB][sqlDB.Ping] %v", err)
		}

		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		grm = db
	})

	return grm
}

var (
	migrateOnce sync.Once
)

func (i *infra) Migrate(values ...interface{}) {
	migrateOnce.Do(func() {
		if i.SetMode() == gin.DebugMode {
			if err := i.GormDB().Debug().AutoMigrate(values...); err != nil {
				logrus.Fatalf("[infra][Migrate][GormDB.Debug.AutoMigrate] %v", err)
			}
		} else if i.SetMode() == gin.ReleaseMode {
			if err := i.GormDB().AutoMigrate(values...); err != nil {
				logrus.Fatalf("[infra][Migrate][GormDB.AutoMigrate] %v", err)
			}
		}
	})
}

var (
	portOnce sync.Once
	port     string
)

func (i *infra) Port() string {
	portOnce.Do(func() {
		port = i.Config().Sub("server").GetString("port")
	})

	return ":" + port
}
