package filesystem

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/LidorAlmkays/MineServerForge/internal/infrastructure"
)

type fileBasedFeatureDataStorage struct {
}

func NewFileBasedFeatureDataStorage() infrastructure.FeaturesDataStorage {
	return &fileBasedFeatureDataStorage{}
}

// SaveFile saves the data into a file in the specified directory.
// It ensures the directory exists and handles file creation and writing.
func (f *fileBasedFeatureDataStorage) SaveFile(fileName string, data []byte, folderSaveDirectory string) error {
	// Ensure the storage directory exists
	if err := f.ensureDirectoryExists(folderSaveDirectory); err != nil {
		return err
	}

	// Construct the full file path
	filePath := filepath.Join(folderSaveDirectory, fileName)

	// Open or create the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("failed to create or open file: " + err.Error())
	}
	defer file.Close()

	// Write the data to the file
	if _, err := file.Write(data); err != nil {
		return errors.New("failed to write data to file: " + err.Error())
	}

	return nil
}

// ensureDirectoryExists checks if a directory exists, and creates it if not.
func (*fileBasedFeatureDataStorage) ensureDirectoryExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Directory doesn't exist; create it
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return errors.New("failed to create directory: " + err.Error())
		}
	}
	return nil
}
