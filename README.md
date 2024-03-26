# Hasty API Go client

[![Build](https://github.com/cloudfactory/hasty-go/workflows/Build/badge.svg)](https://github.com/cloudfactory/hasty-go/actions)

This is a client for Hasty API for gophers.

Hasty is an image annotation tool that makes it easier, faster and more accurate to teach machines the context of what they are looking at. Learn more more at:
- Hasty website [hasty.ai](https://hasty.ai)
- Hasty application [app.hasty.ai](https://app.hasty.ai)
- Hasty API documentation [docs.hasty.ai](https://docs.hasty.ai)

## Usage

First go to workspace settings, then "API accounts". Create new API account for the workspace and then create a new API key for that account.

Get Hasty client:
```
go get -u github.com/cloudfactory/hasty-go
```
Import it into your package:
```
import "github.com/cloudfactory/hasty-go"
```
Obtain and provide API key, and instantiate the client:
```
key := "hK0WXPfeS..."
client := hasty.NewClient(key)

params := &hasty.ImageUploadExternalParams{
	Project:  hasty.String("b72ef832-2bf6-4509-a95a-54c9038bf848"),
	Dataset:  hasty.String("c66779fe-043e-40c6-a419-a9ddcd5ccbff"),
	URL:      hasty.String("https://example.com/cats/one.jpg"),
	Copy:     hasty.Bool(true),
	Filename: hasty.String("cat-one.jpg"),
}
image, err := client.Image.UploadExternal(context.TODO(), params)
fmt.Printf("%v, %v", image, err)
```
