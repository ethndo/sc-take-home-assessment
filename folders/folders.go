package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders:
// Retrieves folders with associated organisation ID and converts them
// into a list of folder pointers and returns them in a FetchFolderResponse object

// FetchAllFoldersByOrgID:
// Filters and returns a list of folder pointers from a sample data set based
// on if the organisation ID matches

// Improvements:
// - f1 and fs are never used
// - FetchAllFoldersByOrgID ignores returned errors
// - Conversion of the slice f to fp has an unnecessary second loop
// - more readable variable names
// - Use dependency injection

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Fetch all folders with given organisation ID
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	// Error Check after fetching folders
	if err != nil {
		return nil, err
	}

	// Return the response
	response := &FetchFolderResponse{Folders: folders}
	return response, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	// Filter based on OrgID
	var filteredFolders []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			filteredFolders = append(filteredFolders, folder)
		}
	}

	return filteredFolders, nil
}