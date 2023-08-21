package repositories_test

import (
	"encoder/src/application/repositories"
	"encoder/src/domain"
	"encoder/src/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	created_video, err := repo.Find(video.ID)

	require.NotEmpty(t, created_video.ID)
	require.Nil(t, err)
	require.Equal(t, created_video.ID, video.ID)
}
