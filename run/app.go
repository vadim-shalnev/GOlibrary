package run

import (
	"database/sql"
	"fmt"
	"github.com/vadim-shalnev/GOlibrary/config"
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/DBemptyChecker"
	"github.com/vadim-shalnev/GOlibrary/internal/infrastructures/responder"
	"github.com/vadim-shalnev/GOlibrary/internal/modules"
	"log"
	"os"
	"time"
)

func ConnectionDB(conf *config.AppConf) *sql.DB {
	time.Sleep(time.Second * 2)
	connect, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return connect
}

func Boostrup(connfig *config.AppConf) *modules.Controllers {

	db := ConnectionDB(connfig)
	DBemptyChecker.CreateTable(db)
	DBemptyChecker.CheckAndFillTables(db)
	//
	log := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	r := responder.NewResponder(log)
	//
	Repositories := modules.NewRepositories(db)
	Services := modules.NewServices(Repositories)
	Controllers := modules.NewControllers(Services, r)
	return Controllers
}
