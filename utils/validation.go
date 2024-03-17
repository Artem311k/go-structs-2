package utils

func ValidateTitle(title string) bool {
	return len(title) < MAX_TITLE_LENGTH
}

func ValidateContent(content string) bool {
	return len(content) < MAX_CONTENT_LENGTH
}
