package oraconn

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
)

// Do should return a Oracle connection string from a preset configuration file
func Do(vars ...string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("invalid user environment: %s", err)
	}

	configName := "database"
	if len(vars) >= 1 {
		configName = vars[1]
	}

	// read configuration for database
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath(usr.HomeDir)
	err = viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("config file not found: %s", err)
	}
	// build oracle connection string
	credentials := viper.GetStringMap("credentials")
	database := viper.GetStringMap("database")
	oraconn := credentials["user"].(string) + "/" + 
				credentials["password"].(string) + "@" + 
				database["tns"].(string)

	return oraconn, err
}
