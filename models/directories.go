package models

type Directories struct {
	Name                string
	status              bool //whether be downloaded
	DownloadedTimestamp int64
}
