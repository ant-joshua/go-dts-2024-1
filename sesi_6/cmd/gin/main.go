package main

import (
	"fmt"
	"sesi_6/pkg/database"
	"sesi_6/pkg/routers"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // path to look for the config file in
	// viper.AddConfigPath("../..") // path to look for the config file in

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	port := ":" + viper.GetString("APP_PORT")

	db := database.SQLInit()

	gorm := database.GormInit(db)

	defer db.Close()

	start := routers.StartServer(db, gorm)
	start.Run(port)

}
