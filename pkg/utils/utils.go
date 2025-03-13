package utils

import (
	"io"
	"mime/multipart"
	"os"
)

// SaveUploadedFile saves an uploaded file
func SaveUploadedFile(file multipart.File, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
