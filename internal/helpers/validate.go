package helpers

func ValidatePassword(password string) bool {
	// Thêm logic validate password nếu cần
	return len(password) >= 6 // ví dụ: ít nhất 6 ký tự
}
