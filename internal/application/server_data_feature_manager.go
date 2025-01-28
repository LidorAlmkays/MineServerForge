package application

type ServerFeaturesDataManager interface {
	SaveMode(ownerEmail, fileName string, data []byte) error
	SavePlugin(ownerEmail, fileName string, data []byte) error
}
