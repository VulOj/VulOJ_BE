package models

type Directories struct {
	Name                string
	Path                string
	Status              bool //whether be downloaded
	DownloadedTimestamp int64
}
