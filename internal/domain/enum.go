package domain

type DeletedParam string

const (
	DeletedExcludeParam DeletedParam = "exclude"
	DeletedOnlyParam    DeletedParam = "only"
	DeletedAllParam     DeletedParam = "all"
)

type DeleteImageURL struct {
	URL string `json:"url" binding:"required,url"`
}

type UploadURLImage struct {
	URL string `json:"url" binding:"required,url"`
	Key string `json:"key" binding:"required"`
}
