package hasty

// ImageUploadExternalParams is used for uploading images from external sources
type ImageUploadExternalParams struct {
	Project  *string `json:"-"`
	Dataset  *string `json:"dataset_id"`
	URL      *string `json:"url"`
	Copy     *bool   `json:"copy_original,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

// Image describes an image information that API may return
type Image struct {
	ID *string `json:"image_id"`
}
