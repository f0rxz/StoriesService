package varsizedint

const MaxSize = 9

func CalcSize(x uint64) int {
	var size = 0

	for {
		size++
		x >>= 7
		if x == 0 {
			break
		}
	}

	if size > MaxSize {
		return MaxSize
	}

	return size
}

func ParseSize(src []uint8) int {
	var x = src[0]

	if (x & 128) == 0 {
		return 1
	} else if (x & (128 | 64)) == 128 {
		return 2
	} else if (x & (128 | 64 | 32)) == (128 | 64) {
		return 3
	} else if (x & (128 | 64 | 32 | 16)) == (128 | 64 | 32) {
		return 4
	} else if (x & (128 | 64 | 32 | 16 | 8)) == (128 | 64 | 32 | 16) {
		return 5
	} else if (x & (128 | 64 | 32 | 16 | 8 | 4)) == (128 | 64 | 32 | 16 | 8) {
		return 6
	} else if (x & (128 | 64 | 32 | 16 | 8 | 4 | 2)) == (128 | 64 | 32 | 16 | 8 | 4) {
		return 7
	} else if (x & (128 | 64 | 32 | 16 | 8 | 4 | 2 | 1)) == (128 | 64 | 32 | 16 | 8 | 4 | 2) {
		return 8
	} else if (x & (128 | 64 | 32 | 16 | 8 | 4 | 2 | 1)) == (128 | 64 | 32 | 16 | 8 | 4 | 2 | 1) {
		return 9
	}

	return -1
}

func Encode(dest []uint8, src uint64) int {
	var size = CalcSize(src)

	switch size {
	case 1:
		dest[0] = uint8(src)
	case 2:
		dest[0] = 128 | uint8(src>>8)
		dest[1] = uint8(src)
	case 3:
		dest[0] = 128 | 64 | uint8(src>>16)
		dest[1] = uint8(src >> 8)
		dest[2] = uint8(src)
	case 4:
		dest[0] = 128 | 64 | 32 | uint8(src>>24)
		dest[1] = uint8(src >> 16)
		dest[2] = uint8(src >> 8)
		dest[3] = uint8(src)
	case 5:
		dest[0] = 128 | 64 | 32 | 16 | uint8(src>>32)
		dest[1] = uint8(src >> 24)
		dest[2] = uint8(src >> 16)
		dest[3] = uint8(src >> 8)
		dest[4] = uint8(src)
	case 6:
		dest[0] = 128 | 64 | 32 | 16 | 8 | uint8(src>>40)
		dest[1] = uint8(src >> 32)
		dest[2] = uint8(src >> 24)
		dest[3] = uint8(src >> 16)
		dest[4] = uint8(src >> 8)
		dest[5] = uint8(src)
	case 7:
		dest[0] = 128 | 64 | 32 | 16 | 8 | 4 | uint8(src>>48)
		dest[1] = uint8(src >> 40)
		dest[2] = uint8(src >> 32)
		dest[3] = uint8(src >> 24)
		dest[4] = uint8(src >> 16)
		dest[5] = uint8(src >> 8)
		dest[6] = uint8(src)
	case 8:
		dest[0] = 128 | 64 | 32 | 16 | 8 | 4 | 2
		dest[1] = uint8(src >> 48)
		dest[2] = uint8(src >> 40)
		dest[3] = uint8(src >> 32)
		dest[4] = uint8(src >> 24)
		dest[5] = uint8(src >> 16)
		dest[6] = uint8(src >> 8)
		dest[7] = uint8(src)
	case 9:
		dest[0] = 128 | 64 | 32 | 16 | 8 | 4 | 2 | 1
		dest[1] = uint8(src >> 56)
		dest[2] = uint8(src >> 48)
		dest[3] = uint8(src >> 40)
		dest[4] = uint8(src >> 32)
		dest[5] = uint8(src >> 24)
		dest[6] = uint8(src >> 16)
		dest[7] = uint8(src >> 8)
		dest[8] = uint8(src)
	}

	return size
}

func Decode(src []uint8) uint64 {
	var size = ParseSize(src)

	switch size {
	case 1:
		return uint64(src[0]) & 127
	case 2:
		return (uint64(src[0]&63) << 8) | uint64(src[1])
	case 3:
		return ((uint64(src[0]) & 31) << 16) |
			(uint64(src[1]) << 8) |
			uint64(src[2])
	case 4:
		return ((uint64(src[0]) & 15) << 24) |
			(uint64(src[1]) << 16) |
			(uint64(src[2]) << 8) |
			uint64(src[3])
	case 5:
		return ((uint64(src[0]) & 7) << 32) |
			(uint64(src[1]) << 24) |
			(uint64(src[2]) << 16) |
			(uint64(src[3]) << 8) |
			uint64(src[4])
	case 6:
		return ((uint64(src[0]) & 3) << 40) |
			(uint64(src[1]) << 32) |
			(uint64(src[2]) << 24) |
			(uint64(src[3]) << 16) |
			(uint64(src[4]) << 8) |
			uint64(src[5])
	case 7:
		return ((uint64(src[0]) & 1) << 48) |
			(uint64(src[1]) << 40) |
			(uint64(src[2]) << 32) |
			(uint64(src[3]) << 24) |
			(uint64(src[4]) << 16) |
			(uint64(src[5]) << 8) |
			uint64(src[6])
	case 8:
		return (uint64(src[1]) << 48) |
			(uint64(src[2]) << 40) |
			(uint64(src[3]) << 32) |
			(uint64(src[4]) << 24) |
			(uint64(src[5]) << 16) |
			(uint64(src[6]) << 8) |
			uint64(src[7])
	case 9:
		return (uint64(src[1]) << 56) |
			(uint64(src[2]) << 48) |
			(uint64(src[3]) << 40) |
			(uint64(src[4]) << 32) |
			(uint64(src[5]) << 24) |
			(uint64(src[6]) << 16) |
			(uint64(src[7]) << 8) |
			uint64(src[8])
	}

	return 0xffffffffffffffff
}
