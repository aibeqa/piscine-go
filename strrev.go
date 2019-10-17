package piscine

func StrRev(s string) string {
	var tmp_arr []byte = []byte(s)
	var l int = 0
	var r int = len(tmp_arr) - 1

	for l < r {
		tmp := tmp_arr[l]
		tmp_arr[l] = tmp_arr[r]
		tmp_arr[r] = tmp
		l++
		r--
	}
	return string(tmp_arr)
}
