package utils

import "mime/multipart"

func IsPDF(file *multipart.FileHeader) bool {
	f, err := file.Open()
	if err != nil {
		return false
	}
	defer f.Close()

	// PDF files start with "%PDF"
	buf := make([]byte, 4)
	_, err = f.Read(buf)
	if err != nil {
		return false
	}

	return string(buf) == "%PDF"
}
