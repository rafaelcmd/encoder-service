package services_test

import (
	"log"
	"testing"
	"time"
	"github.com/joho/godotenv"
	"github.com/rafaelcmd/encoder-service/application/repositories"
	"github.com/rafaelcmd/encoder-service/domain"
	"github.com/rafaelcmd/encoder-service/framework/database"
	"github.com/rafaelcmd/encoder-service/application/services"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "vikings.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db:db}
	
	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("encoder-service-bucket")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}