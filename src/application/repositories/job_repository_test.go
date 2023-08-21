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

func TestNewJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, insert_job_err := domain.NewJob("path", "converted", video)

	repo := repositories.NewJobRepository(db)
	repo.Insert(job)

	created_job, find_repo_err := repo.Find(job.ID)

	require.Nil(t, insert_job_err)
	require.Nil(t, find_repo_err)
	require.NotEmpty(t, job)
	require.NotEmpty(t, created_job.ID)
	require.Equal(t, created_job.ID, job.ID)
}

func TestNewJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, insert_job_err := domain.NewJob("path", "converted", video)

	repo := repositories.NewJobRepository(db)
	repo.Insert(job)

	created_job, find_repo_err := repo.Find(job.ID)

	require.Nil(t, insert_job_err)
	require.Nil(t, find_repo_err)
	require.NotEmpty(t, job)
	require.NotEmpty(t, created_job.ID)
	require.Equal(t, created_job.ID, job.ID)

	created_job.Status = "completed"
	updated_job, update_repo_err := repo.Update(created_job)

	require.Nil(t, update_repo_err)
	require.NotEmpty(t, updated_job)
	require.Equal(t, updated_job.Status, "completed")
}
