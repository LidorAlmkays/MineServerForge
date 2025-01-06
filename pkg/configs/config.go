package configs

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

type StructConstraint interface {
}

func GetConfig[T any](cfg *T, fileFullPath string, fileType enums.ConfigTypes) (*T, error) {
	var err error

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

	printConfigs(cfg)

	return cfg, nil
}

func printConfigs[T any](cfg *T) {
	// This is a way to print the configs
	result, _ := json.Marshal(cfg)
	fmt.Println(string(result))
}
