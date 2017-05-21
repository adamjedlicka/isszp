// Install package provides installation process of the application.
// From creating database to settings default user passwords
package install

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
)

var (
	DefaultConfig = "./config/config.default.json"
	CustomConfig  = "./config/config.json"
)

// InstallConfig copies the default config.default.json, generates new data and saves it as config.json
func InstallConfig() error {
	err := common.CopyFile(CustomConfig, DefaultConfig)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(CustomConfig)
	if err != nil {
		return err
	}

	var cfg interface{}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return err
	}

	cfgMap := cfg.(map[string]interface{})
	cfgController := cfgMap["Controller"].(map[string]interface{})

	log.Println("Generating new controller secret...")
	cfgController["Secret"] = common.RandomString(32)

	bytes, err = json.MarshalIndent(cfgMap, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(CustomConfig, bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
