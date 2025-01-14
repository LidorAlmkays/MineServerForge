package featuresdatamanager

type FeaturesDataManager interface {
	SaveFile(fileName string, data []byte, folderSaveDirectory string) error
}
