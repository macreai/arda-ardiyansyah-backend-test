package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	idleConnection := viper.GetInt("DB_POOL_IDLE")
	maxPool := viper.GetInt("DB_POOL_MAX")
	maxLifetime := viper.GetInt("DB_POOL_LIFETIME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbName, port)

	// m, err := migrate.New(
	// 	"file://db/migrations",
	// 	fmt.Sprintf(
	// 		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 		username, password, host, port, dbName,
	// 	),
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal(err)
	// }

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{
			Logger: log,
		}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxPool)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	return db

}

type logrusWriter struct {
	*logrus.Logger
}

func (l *logrusWriter) Print(messaage string, args ...interface{}) {
	l.Logger.Tracef(messaage, args...)
}
