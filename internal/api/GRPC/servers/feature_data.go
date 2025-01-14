package servers

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/LidorAlmkays/MineServerForge/internal/api/GRPC/pb"
	"github.com/LidorAlmkays/MineServerForge/internal/application/serverfeaturedatamanager"
	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)

type featureDataServer struct {
	l logger.Logger
	s serverfeaturedatamanager.ServerFeaturesDataManager
	pb.UnimplementedUploadFeatureDataServer
}

func NewFeatureDataServer(l logger.Logger, s serverfeaturedatamanager.ServerFeaturesDataManager) pb.UploadFeatureDataServer {
	return &featureDataServer{l: l, s: s}
}

func (u *featureDataServer) SaveFeature(stream pb.UploadFeatureData_SaveFeatureDataServer) error {
	var filename string
	fileData := []byte{}
	var ownerEmail string
	receivedData := false // Flag to track if any data is received
	var featureType pb.FeatureType

	for {
		streamChunk, err := stream.Recv()
		if err == io.EOF {
			if !receivedData {
				u.l.Info("Received a file save with no data.")
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Empty file upload not allowed",
					Success: false,
				})
			}

			// Save the file based on the feature type
			var err error
			switch featureType {
			case pb.FeatureType_MODE:
				err = u.s.SaveMode(ownerEmail, filename, fileData)
			case pb.FeatureType_PLUGIN:
				err = u.s.SavePlugin(ownerEmail, filename, fileData)
			default:
				u.l.Info("Received a save file with an invalid feature type.")
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Invalid feature type",
					Success: false,
				})
			}
			if err != nil {
				u.l.Error(errors.New("Failed to save file: " + err.Error()))
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Failed to save file",
					Success: false,
				})
			}

			u.l.Message(fmt.Sprintf("File uploaded successfully, file name: %s", filename))
			return stream.SendAndClose(&pb.UploadStatus{
				Message: "File uploaded successfully",
				Success: true,
			})
		}

		if err != nil {
			log.Printf("Error receiving file chunk: %v", err)
			return err
		}

		// Validate and process the first chunk
		if !receivedData {
			if streamChunk.GetFileChunk().GetFilename() == "" {
				log.Println("Received chunk without filename")
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Filename is required",
					Success: false,
				})
			}
			// Check if ownerEmail is provided
			ownerEmail = streamChunk.GetFileChunk().GetOwnerEmail()
			if ownerEmail == "" {
				log.Println("Owner email is missing")
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Owner email is required",
					Success: false,
				})
			}
			filename = streamChunk.GetFileChunk().GetFilename()
			featureType = streamChunk.GetFeatureType()
			if featureType != pb.FeatureType_MODE && featureType != pb.FeatureType_PLUGIN {
				log.Printf("Invalid feature type received: %v", featureType)
				return stream.SendAndClose(&pb.UploadStatus{
					Message: "Invalid feature type",
					Success: false,
				})
			}
		}
		fileData = append(fileData, streamChunk.FileChunk.Data...)
		receivedData = true
	}
}
