type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "empty"
	}
	if len(strs) == 1 {
		return strs[0]
	}
	start := "---"
	builder := strings.Join(strs, "|||")
	builder = start + builder
	return builder
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "empty" {
		return []string{}
	}
	str, _ := strings.CutPrefix(encoded, "---")
	if len(str) == 0 {
		return []string{str}
	}
	if len(str) == 1 {
		return []string{string(str[0])}
	}
	finale := strings.Split(str, "|||")
	return finale
}