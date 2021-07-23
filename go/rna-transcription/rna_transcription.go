package strand

func ToRNA(dna string) string {
	rnaSeq := ""
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			rnaSeq = rnaSeq + "C"
		case 'C':
			rnaSeq = rnaSeq + "G"
		case 'T':
			rnaSeq = rnaSeq + "A"
		case 'A':
			rnaSeq = rnaSeq + "U"
		}
	}
	return rnaSeq
}
