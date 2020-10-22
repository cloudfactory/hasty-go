package hasty

// ImageUploadParams is used for uploading images
type ImageUploadParams struct {
	Project *string `json:"-"`
	Dataset *string `json:"dataset_id,omitempty"`
	URL     *string `json:"url,omitempty"`
	Copy    *bool   `json:"copy_original,omitempty"`
}
