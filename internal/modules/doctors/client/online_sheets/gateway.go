package online_sheets

import (
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
)

type SheetsGateway struct {
	sheetsService    *sheets.Service
	driveService     *drive.Service
	spreadsheetID    string
	spreadsheetTitle string
}

func New() *SheetsGateway {
	return &SheetsGateway{}
}
