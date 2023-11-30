package structs

import "os"

type fileSummaryResponse struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func FilesSummaryResponse(files []File) []fileSummaryResponse {
	listFiles := []fileSummaryResponse{}

	for _, file := range files {
		url := os.Getenv("APP_URL") + "images/" + file.Location + "/" + file.Name

		fileFormatter := fileSummaryResponse{
			ID: file.ID,
			URL: url,
		}

		listFiles = append(listFiles, fileFormatter)
	}

	return listFiles
}