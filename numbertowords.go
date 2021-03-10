package numbertowords

import "errors"

var words = [20]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tenwords = [10]string{
	"",
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

// maxNumber is the max. limit.
const maxNumber = 99999

// Convert converts an integer to words.
func Convert(index int) (string, error) {

	// Invalid condiiton
	if index < 0 || index > maxNumber {
		return "", errors.New("number not in valid range") //fmt.Errorf()
	}

	// Valid ones

	result := ""
	result1 := ""
	result2 := ""

	tenthousands := index / 10000

	if tenthousands > 0 {
		result2 = tenwords[tenthousands]
		index = index % 10000
		if index == 0 {
			return result2, nil
		}
	}

	thousands := index / 1000

	if thousands > 0 {
		result1 = words[thousands] + " thousand "
		index = index % 1000
		if index == 0 {
			return result1, nil
		}

	}

	hundreds := index / 100

	if hundreds > 0 {
		result = words[hundreds] + " hundred "
		index = index % 100
		if index == 0 {
			return result, nil
		}
	}

	tens := index / 10
	units := index % 10

	if tens < 2 {
		return result2 + " " + result1 + result + words[index], nil
	}

	if units == 0 {
		return result2 + " " + result1 + result + tenwords[tens], nil
	}

	return result2 + " " + result1 + result + tenwords[tens] + " " + words[units], nil
}
