package piscine

func IsPrime(nb int) bool {
	for i := 0; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
