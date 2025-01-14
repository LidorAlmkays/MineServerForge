package serverfeaturedatamanager

import (
	"errors"
	"path/filepath"

	"github.com/LidorAlmkays/MineServerForge/config"
	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure/featuresdatamanager"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)

type filesBasedFeatureDataManager struct {
	cfg *config.Config
	l   logger.Logger
	fdm featuresdatamanager.FeaturesDataManager
}

func NewFilesBasedFeatureDataManager(cfg *config.Config, l logger.Logger, fDM featuresdatamanager.FeaturesDataManager) ServerFeaturesDataManager {
	return &filesBasedFeatureDataManager{cfg, l, fDM}
}

// SaveModeChunk saves the data into a file in the specified directory.
// It ensures the directory exists and handles file creation and writing.
func (m *filesBasedFeatureDataManager) SaveMode(ownerEmail string, fileName string, data []byte) error {
	if m.cfg.StorageConfig.MinecraftPluginsStorage == nil {
		return errors.New("using file based saving, mode storage wasn't set in config")
	}
	err := m.fdm.SaveFile(fileName, data, filepath.Join(*m.cfg.StorageConfig.MinecraftModesStorage, ownerEmail))
	if err != nil {
		return err
	}
	return nil
}

// SaveModeChunk saves the data into a file in the specified directory.
// It ensures the directory exists and handles file creation and writing.
func (m *filesBasedFeatureDataManager) SavePlugin(ownerEmail string, fileName string, data []byte) error {
	if m.cfg.StorageConfig.MinecraftPluginsStorage == nil {
		return errors.New("using file based saving, plugin storage wasn't set in config")
	}
	err := m.fdm.SaveFile(fileName, data, filepath.Join(*m.cfg.StorageConfig.MinecraftPluginsStorage, ownerEmail))
	if err != nil {
		return err
	}
	return nil
}
