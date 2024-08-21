package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
	"github.com/gofrs/uuid"
)

func Test_GetAllFolders(t *testing.T) {
	// new UUID for testing
	orgID := uuid.Must(uuid.NewV4())
	// sample request
	req := &folders.FetchFolderRequest{OrgID: orgID}

	// Mock GetSampleData to return dummy data
	originalGetSampleData := folders.GetSampleData
	folders.GetSampleData = func() []*folders.Folder {
		return []*folders.Folder{
			{OrgId: orgID, Name: "Folder 1"},
			{OrgId: orgID, Name: "Folder 2"},
			{OrgId: uuid.Must(uuid.NewV4()), Name: "Dummy Folder"},
		}
	}
	defer func() { folders.GetSampleData = originalGetSampleData }()


	t.Run("Matching Organisation ID", func(t *testing.T) {
		response, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 2, len(response.Folders))
		assert.Equal(t, "Folder 1", response.Folders[0].Name)
		assert.Equal(t, "Folder 2", response.Folders[1].Name)
	})

	t.Run("No Matching Organisation ID", func(t *testing.T) {
		randomOrgID := uuid.Must(uuid.NewV4())
		req := &folders.FetchFolderRequest{OrgID: randomOrgID}

		response, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 0, len(response.Folders))
	})
}
