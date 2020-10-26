package hasty

import (
	"context"
	"testing"
)

type dummyBackend struct{}

func (b *dummyBackend) Request(ctx context.Context, method, path string, payload, response interface{}) error {
	return nil
}

func TestImageClient_UploadExternal(t *testing.T) {
	c := &ImageClient{
		backend: &dummyBackend{},
	}
	ctx := context.TODO()
	tests := []struct {
		name    string
		project *string
		wantErr bool
	}{
		{"with project", String("foo"), false},
		{"with empty project", String(""), true},
		{"with nil project", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.UploadExternal(ctx, &ImageUploadExternalParams{Project: tt.project})
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageClient.UploadExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
