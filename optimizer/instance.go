package optimizer

func CopyList(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}
	return out
}
