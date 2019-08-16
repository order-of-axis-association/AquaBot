package db

import (
	"fmt"
	"io/ioutil"
	"time"

	"crypto/tls"
	"crypto/x509"
	"github.com/go-sql-driver/mysql"

	"github.com/order-of-axis-association/AquaBot/types"

	_ "github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var config_loc string = "secrets/db_config.yml"

var ca_cert_path string = "secrets/server-ca.pem"
var client_cert_path string = "secrets/client-cert.pem"
var client_key_path string = "secrets/client-key.pem"

var cloud_sql_server_name string = "oa-aquabot:aquabot-master"

var tls_config_name string = "custom"
var new_york_location string = "America/New_York"

func registerAquabotTLSConfig() {
	root_cert_pool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(ca_cert_path)
	if err != nil {
		fmt.Println("Could not read ca-cert.pem:", err)
	}
	if ok := root_cert_pool.AppendCertsFromPEM(pem); !ok {
		fmt.Println("Failed to append PEM.")
	}

	client_cert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair(client_cert_path, client_key_path)
	if err != nil {
		fmt.Println("Could not load X509 Keypair:", err)
	}

	client_cert = append(client_cert, certs)

	mysql.RegisterTLSConfig(tls_config_name, &tls.Config{
		RootCAs:      root_cert_pool,
		Certificates: client_cert,
		ServerName:   cloud_sql_server_name,
	})
}

func BuildCloudSQLDSN() string {
	config_raw, err := ioutil.ReadFile(config_loc)
	if err != nil {
		fmt.Println("Could not read db_config.yml file: ", err)
	}

	config := types.DBConfig{}

	err = yaml.Unmarshal([]byte(config_raw), &config)
	if err != nil {
		fmt.Println("Could not unmarshal db.Config data")
	}

	fmt.Println("%+v", config)

	registerAquabotTLSConfig()

	new_york_loc, err := time.LoadLocation(new_york_location)
	if err != nil {
		fmt.Println("Could not load America/New_York location!:", err)
	}

	cfg := mysql.Config{
		User:                 config.User,
		Passwd:               config.Password,
		Addr:                 config.Host + ":3306",
		Net:                  "tcp",
		DBName:               config.DBName,
		Loc:                  new_york_loc,
		AllowNativePasswords: true,

		ParseTime: true,

		TLSConfig: tls_config_name,
	}

	dsn := cfg.FormatDSN()
	fmt.Println(dsn)

	dsn = dsn + "&charset=utf8mb4" // Lol the mysql config source code literally doesn't have logic for a charset option

	fmt.Println(dsn)
	return dsn
}

func Migrate(global_state types.G_State) {
	fmt.Println("%+v", global_state)
	g_db := global_state.DBConn

	g_db.AutoMigrate(&Guild{})
	g_db.AutoMigrate(&Channel{})
	g_db.AutoMigrate(&User{})
	g_db.AutoMigrate(&Reminder{})
	g_db.AutoMigrate(&Todo{})
	g_db.AutoMigrate(&Config{})

	fmt.Println("Done migrating.")

}
