package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/api"
	grpc "github.com/LidorAlmkays/MineServerForge/internal/api/GRPC"
	rest "github.com/LidorAlmkays/MineServerForge/internal/api/REST"
	"github.com/LidorAlmkays/MineServerForge/internal/application/serverfeaturedatamanager"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/featuresdatamanager"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/validators"
	"github.com/go-playground/validator"
)

type ProgramFlags struct {
	Mode enums.ProgramMode `validate:"required,programmode"`
}

var programFlags ProgramFlags

func init() {
	validate := validator.New()
	validate.RegisterValidation("programmode", validators.ProgramModeValidator)

	var mode string

	flag.StringVar(&mode, "Mode", "development", "This flags changes the program mode")
	flag.Parse()

	mode = strings.ToLower(mode)
	programFlags.Mode = enums.ProgramMode(mode)
	err := validate.Struct(programFlags)
	if err != nil {
		panic(err)
	}
	fmt.Println("Project Mode: " + programFlags.Mode)
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	errChan, servers := setUp()

	// Wait for either an OS signal or an error from one of the servers
	select {
	case <-signalChan:
		log.Println("Received termination signal")
	case err := <-errChan:
		fmt.Printf("Error occurred: %v. Shutting down all servers...\n", err)
	}
	for _, s := range servers {
		err := s.Shutdown() // Shutdown servers gracefully
		if err != nil {
			fmt.Println("Error when shutdown :" + err.Error())
			os.Exit(1)
		}
	}
	os.Exit(0)
}

func setUp() (chan error, []api.BaseServer) {
	ctx := context.Background()

	var cfg *config.Config = &config.Config{}

	cfg.SetUp("./config/.env", enums.ENV)
	var l logger.Logger = logger.NewStackedCustomLogger(programFlags.Mode, cfg.ServiceConfig.ProjectName)

	fdm := featuresdatamanager.NewFileBasedFeatureDataManager()

	var sFeatures serverfeaturedatamanager.ServerFeaturesDataManager = serverfeaturedatamanager.NewFilesBasedFeatureDataManager(cfg, l, fdm)

	// Example of dynamically adding more servers (e.g., REST, GRPC)
	var servers = []api.BaseServer{
		rest.NewServer(ctx, cfg, l),
		grpc.NewServer(ctx, cfg, l, sFeatures),
	}

	// Channel to capture errors
	errCh := make(chan error, len(servers))
	wg := sync.WaitGroup{}

	// Start servers in separate goroutines
	for _, s := range servers {
		wg.Add(1)
		go func(server api.BaseServer) {
			defer wg.Done()
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				errCh <- err
			}
		}(s)
	}

	// Wait for an error or shutdown signal
	go func() {
		wg.Wait()
		close(errCh)
	}()

	return errCh, servers
}
