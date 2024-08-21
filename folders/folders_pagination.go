package folders

import (
	"github.com/gofrs/uuid"
)
// EXPLANATION
// To facilitate pagination, we add page and pageSize parameters
// to GetAllFolders and FetchAllFoldersByOrgID to control pagination
//  The start and end indexes are calculated based on page and pageSize
// and based on these values, we slice the filtered folders to only return the
// desired subset.
// We must also check the validity of these variables by checking if
// start and end are in bounds to prevent any errors. If the start index
// exceeds the length of the filtered folders, we just return an empty list.

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
func GetAllFoldersPaginated(req *FetchFolderRequest, page, pageSize int) (*FetchFolderResponse, error) {
	// Fetch paginated folders with given organisation ID
	folders, err := FetchAllFoldersByOrgIDPaginated(req.OrgID, page, pageSize)
	// Error Check after fetching folders
	if err != nil {
		return nil, err
	}

	// Return the response
	response := &FetchFolderResponse{Folders: folders}
	return response, nil
}

func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, page, pageSize int) ([]*Folder, error) {
	folders := GetSampleData()

	// Filter based on OrgID
	var filteredFolders []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			filteredFolders = append(filteredFolders, folder)
		}
	}

	// Pagination Logic
	start := (page - 1) * pageSize
	end := start + pageSize

	// Make sure that the end does not exceed sliced length
	if start >= len(filteredFolders) {
		// Start index out of bounds
		return []*Folder{}, nil
	}
	if end > len(filteredFolders) {
		end = len(filteredFolders)
	}

	return filteredFolders[start:end], nil
}