package lib

func Convert(bytes uint64) (uint64, string) {
	var tmp float64 = float64(bytes)
	var s string = " "

	bytes = uint64(tmp)

	switch {
	case bytes < uint64(2<<9):
		s = "B"
	case bytes < uint64(2<<19):
		tmp = tmp / float64(2<<9)
		s = "K"

	case bytes < uint64(2<<29):
		tmp = tmp / float64(2<<19)
		s = "M"

	case bytes < uint64(2<<39):
		tmp = tmp / float64(2<<29)
		s = "G"

	case bytes < uint64(2<<49):
		tmp = tmp / float64(2<<39)
		s = "T"

	}
	return uint64(tmp), s
}