package internal

func validateAddress(address string) bool {
	if len(address) == 34 && (string(address[0]) == "1" || string(address[0]) == "3") {
		return true
	}

	if len(address) == 42 && (string(address[:3]) == "bc1") {
		return true
	}

	return false
}
