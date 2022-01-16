package services

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/consts"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net"
	"net/smtp"
	"strings"
)

func dsn(settings models.DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4,utf8", settings.Username, settings.Password, settings.Hostname, settings.Dbname)
}

//The following variables are defined for Database services

//The following variables are defined for email services
var db *gorm.DB
var client *smtp.Client
var account string
var password string
var servername string

//For Redis services
var RedisClient *redis.Client

func init() {
	databaseInit()
	emailInit()
	redisInit()
}

func databaseInit() {
	conf := util.ReadSettingsFromFile("Config.json")
	settings := conf.DbSettings
	connStr := dsn(settings)

	dbStr := strings.Replace(connStr, settings.Dbname, "", 1)
	msdb, e := sql.Open("mysql", dbStr)
	util.CheckError(e)
	msdb.Exec("create database if not exists " + settings.Dbname + " character set utf8")
	msdb.Close()

	var err1 error
	db, err1 = gorm.Open("mysql", connStr)
	//db.DB().SetMaxIdleConns(0)
	util.CheckError(err1)

	var temp []interface{}
	var holeUserType models.Auth
	var blogType models.Blog
	var commentType models.Comment
	//创建blogForbidden
	var blogForbiddenType models.BlogForbidden
	temp = append(temp, &holeUserType, &blogType, &commentType, &blogForbiddenType)
	util.CreateTableIfNotExist(db, temp)

	passwordHash := util.HashWithSalt("root")
	var u models.Auth
	db.Where("email=?", "root").Find(&u)
	if (u == models.Auth{}) {
		user := models.Auth{
			Email: "root", Password: passwordHash,
			RegisterTimestamp: util.GetTimeStamp(),
			Role:              consts.ROOT}
		db.Create(user)
	}
}
func emailInit() {

	conf := util.ReadSettingsFromFile("Config.json")
	account = conf.EmailSenderSettings.Email
	password = conf.EmailSenderSettings.Password
	servername = conf.EmailSenderSettings.Servername
	// Connect to the SMTP Server

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", account, password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		//log.Panic(err)
		fmt.Println(err)
	}

	client, err = smtp.NewClient(conn, host)
	if err != nil {
		//log.Panic(err)
		fmt.Println(err)
	}

	// Auth
	if err = client.Auth(auth); err != nil {
		//log.Panic(err)
		fmt.Println(err)
	}
	//go HandleMultipleEmail()
}
func redisInit() {
	redisSettings := util.ReadSettingsFromFile(consts.CONFIG_FILE_NAME).RedisSettings
	RedisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     redisSettings.Address + ":" + redisSettings.Port,
		Password: redisSettings.Password,
		DB:       0,
	})
}
