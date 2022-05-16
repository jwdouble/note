package yamlx

func findFirst(in string) int {
	for n, v := range in {
		if v > 32 {
			return n
		}
	}
	return 0
}

func TrimAllSpace(s string) string {
	var buf []byte
	for _, v := range s {
		if v == 32 {
			continue
		}
		buf = append(buf, byte(v))
	}
	return string(buf)
}
