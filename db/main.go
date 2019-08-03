package db

import (
	"fmt"
	"io/ioutil"

	//"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	User		string
	Password	string
	DBName		string
	Host		string
}

var config_loc string = "secrets/db_config.yml"
var dsn_fmt string = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=America%%2FNew_York"

func BuildCloudSQLDSN() string {
	config_raw, err := ioutil.ReadFile(config_loc)
	if err != nil {
		fmt.Println("Could not read db_config.yml file: ", err)
	}

	config := DBConfig{}

	err = yaml.Unmarshal([]byte(config_raw), &config)
	if err != nil {
		fmt.Println("Could not unmarshal db.Config data")
	}

	fmt.Println("%+v", config)

	dsn := fmt.Sprintf(dsn_fmt, config.User, config.Password, config.Host, config.DBName)

	fmt.Println(dsn)
	return dsn
}

func Migrate() {
	
}
