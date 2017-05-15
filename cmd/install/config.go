package install

import (
	"gitlab.fit.cvut.cz/isszp/isszp/src/common"
)

func InstallConfig() error {
	err := common.CopyFile("./config/config.json", "./config/config.default.json")
	if err != nil {
		return err
	}

	return nil
}
