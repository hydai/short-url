package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hydai/short-url/config"
	"github.com/hydai/short-url/utils"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	utils.ErrorOrNil(err)
	viper.WatchConfig()

	var configuration config.Configuration

	err = viper.Unmarshal(&configuration)
	utils.ErrorOrNil(err)

	fmt.Print(configuration.Server.Port)
	sqlConfig := fmt.Sprintf(
		"%s:%s@/%s",
		configuration.Database.User,
		configuration.Database.Password,
		configuration.Database.Dbname,
	)

	db, err := sql.Open("mysql", sqlConfig)
	defer db.Close()
	utils.ErrorOrNil(err)

	rows, err := db.Query("SELECT * FROM URLMAP")
	utils.ErrorOrNil(err)
	for rows.Next() {
		var uid int
		var surl string
		var ourl string
		err = rows.Scan(&uid, &surl, &ourl)
		utils.ErrorOrNil(err)
		fmt.Println("UID: " + strconv.Itoa(uid))
		fmt.Println("SURL: " + surl)
		fmt.Println("OURL: " + ourl)
	}
}
