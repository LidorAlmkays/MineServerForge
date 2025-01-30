package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/api"
	grpc "github.com/LidorAlmkays/MineServerForge/internal/api/GRPC"
	rest "github.com/LidorAlmkays/MineServerForge/internal/api/REST"
	"github.com/LidorAlmkays/MineServerForge/internal/application"
	"github.com/LidorAlmkays/MineServerForge/internal/application/serverdatamanager"
	"github.com/LidorAlmkays/MineServerForge/internal/application/serverfeaturedatamanager"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/db"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/filesystem"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

// @title Swagger Example API
// @version 1.0
// @description Sample server for managing game servers

// @host localhost:5000
// @BasePath /
func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	errChan, servers, err := setUp()
	if err != nil {
		fmt.Println("Error accord during project setup: " + err.Error())
		os.Exit(1)
	}

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

func setUp() (chan error, []api.BaseServer, error) {
	var cfg *config.Config = &config.Config{}

	err := cfg.SetUp("./config/.env", enums.ENV, config.Flags.Mode == enums.DevelopmentMode)
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()

	var l logger.Logger = logger.NewStackedCustomLogger(config.Flags.Mode, cfg.ServiceConfig.ProjectName)

	fdm := filesystem.NewFileBasedFeatureDataStorage()

	var sFeatures application.ServerFeaturesDataManager = serverfeaturedatamanager.NewFilesBasedFeatureDataManager(cfg, l, fdm)
	dbF := db.GetDBFactory(l)
	mineDb, err := dbF.GetMinecraftServer(l, enums.Postgres, cfg.DbConfig)
	if err != nil {
		return nil, nil, err
	}

	var s application.ServerConfigDataManager = serverdatamanager.NewBaseServerConfigDataManager(mineDb)

	// Example of dynamically adding more servers (e.g., REST, GRPC)
	var servers = []api.BaseServer{
		rest.NewServer(ctx, cfg, l, s),
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
	return errCh, servers, nil
}
