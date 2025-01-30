package infrastructure

type FeaturesDataStorage interface {
	SaveFile(fileName string, data []byte, folderSaveDirectory string) error
}
