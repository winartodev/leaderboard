package helper

import (
	"math"
	"math/bits"
)

func ExtractBitToInt64(score int64, indexStart int, length int) (result int64) {
	totalBits := bits.Len((1<<64 - 1))
	if indexStart >= 0 && int32(indexStart)+int32(length) > int32(totalBits) || indexStart < 0 && math.Abs(float64(indexStart))+1 < float64(length) {
		return 0
	}

	mask := (1 << length) - 1
	var maskShift int

	if indexStart >= 0 {
		maskShift = totalBits + indexStart - length

	} else {
		maskShift = int(math.Abs(float64(indexStart) + 1 - float64(length)))
	}

	mask = mask << maskShift

	result = (score & int64(mask)) >> int64(maskShift)
	return result
}
