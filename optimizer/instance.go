package optimizer

func CopyList(in []string) []string {
	var out []string
	out = append(out, in...)
	return out
}
