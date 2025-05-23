package utils

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type real interface {
	~float32 | ~float64
}

func IntDivCeil[intTypes integer](number intTypes, divisor intTypes) intTypes {
	result := number / divisor
	if number%divisor != 0 {
		result++
	}

	return result
}

func TruncateFloat[floatTypes real](num floatTypes) floatTypes {
	// Truncate floating point number to 2 digits behind comma
	return floatTypes(int(num*100)) / 100
}
