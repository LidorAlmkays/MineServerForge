package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/LidorAlmkays/MineServerForge/config"
	rest "github.com/LidorAlmkays/MineServerForge/internal/api/REST"
	"github.com/LidorAlmkays/MineServerForge/pkg/configs"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/validators"
	"github.com/go-playground/validator"
)

type ProgramFlags struct {
	Mode string `validate:"required,programmode"`
}

var programFlags ProgramFlags

func init() {
	flag.StringVar(&programFlags.Mode, "Mode", "development", "This flags changes the program mode")
	flag.Parse()
	programFlags.Mode = strings.ToLower(programFlags.Mode)
	validate := validator.New()
	validate.RegisterValidation("programmode", validators.ProgramModeValidator)
	err := validate.Struct(programFlags)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}

func setUp() error {
	ctx := context.Background()
	var err error
	var cfg *config.Config = &config.Config{}

	cfg, err = configs.GetConfig(cfg, "../config/.env", enums.ENV)
	if err != nil {
		return err
	}

	var l logger.Logger = logger.NewStackedCustomLogger(cfg.ServiceConfig.ProjectName)

	err = rest.NewServer(ctx, cfg, l).ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
