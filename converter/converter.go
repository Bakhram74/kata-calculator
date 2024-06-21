package converter

type Converter struct {
	Roman   map[string]int
	Arabian map[int]string
}

func NewConverter() *Converter {
	roman := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"L":    50,
		"C":    100,
		"D":    500,
		"M":    1000,
	}
	arabian := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	return &Converter{
		Roman:   roman,
		Arabian: arabian,
	}
}

func (converter *Converter) IsRoman(num string) bool {
	_, ok := converter.Roman[num]
	return ok
}

func (converter *Converter) ToNumber(n string) int {
	out := 0
	ln := len(n)
	for i := 0; i < ln; i++ {
		str := string(n[i])
		vc := converter.Roman[str]
		if i < ln-1 {
			cnext := string(n[i+1])
			vcnext := converter.Roman[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}
func (converter *Converter) ToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += converter.Arabian[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}
