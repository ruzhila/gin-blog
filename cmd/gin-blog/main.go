package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/ruzhila/gin-blog/internal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var GitCommit string
var BuildTime string

func main() {
	var addr string
	var runMigration bool
	var logFile string = os.Getenv("LOG_FILE")
	var traceSql bool = os.Getenv("TRACE_SQL") != ""
	var logerLevel string = os.Getenv("LOG_LEVEL")
	var dbDriver string = os.Getenv("DB_DRIVER")
	var dsn string = os.Getenv("DSN")

	var superUserEmail string
	var superUserPassword string

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		},
	})

	flag.StringVar(&superUserEmail, "superuser", "", "Create an super user with email")
	flag.StringVar(&superUserPassword, "password", "", "Super user password")
	flag.StringVar(&addr, "addr", ":8080", "HTTP Serve address")
	flag.StringVar(&logFile, "log", logFile, "Log output file name, default is os.Stdout")
	flag.StringVar(&logerLevel, "level", logerLevel, "Log level debug|info|warn|error")
	flag.BoolVar(&runMigration, "m", false, "Run migration only")
	flag.StringVar(&dbDriver, "db", dbDriver, "DB Driver, sqlite|mysql")
	flag.StringVar(&dsn, "dsn", dsn, "DB DSN")
	flag.BoolVar(&traceSql, "tracesql", traceSql, "Trace sql execution")
	flag.Parse()
	var lw io.Writer = os.Stdout
	var err error

	if logFile != "" {
		lw, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("open %s fail, %v\n", logFile, err)
		} else {
			logrus.SetOutput(lw)
		}
	} else {
		logFile = "console"
	}

	fmt.Println("GitCommit   =", GitCommit)
	fmt.Println("BuildTime   =", BuildTime)
	fmt.Println("addr        =", addr)
	fmt.Println("logfile     =", logFile)
	fmt.Println("logerLevel  =", logerLevel)
	fmt.Println("DB Driver   =", dbDriver)
	fmt.Println("DSN         =", dsn)
	fmt.Println("traceSql    =", traceSql)
	fmt.Println("migration   =", runMigration)

	// Init Database
	if runMigration {
		fmt.Println("migration done")
		return
	}

	r := gin.New()
	var db *gorm.DB
	m := internal.NewBlogApp(db)

	logConfig := gin.LoggerConfig{
		Output: lw,
	}
	r.Use(gin.LoggerWithConfig(logConfig), gin.Recovery())

	if err = m.Prepare(r); err != nil {
		panic(err)
	}

	fmt.Println("ruzhila blog is running on", addr)
	r.Run(addr)
}
