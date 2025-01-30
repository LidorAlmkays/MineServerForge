package configs

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

func GetConfig[T ConfigObject](fileFullPath string, fileType enums.ConfigTypes, debugMode bool) (*T, error) {
	var err error
	cfg := new(T)
	switch fileType {
	case enums.ENV:
		{
			err = loadConfigFromEnvOrFile(cfg, fileFullPath)
			if err != nil {
				fmt.Print(err.Error())
				return nil, err
			}
		}
	default:
		{
			err := errors.New("no config file type was selected, ENV is also path variables")
			fmt.Print(err)
			return nil, err
		}
	}

	if debugMode {
		printConfigs(cfg)
	}

	if err = (*cfg).ValidateSelf(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func printConfigs[T any](cfg *T) {
	// This is a way to print the configs
	result, _ := json.Marshal(cfg)
	fmt.Println(string(result))
}
