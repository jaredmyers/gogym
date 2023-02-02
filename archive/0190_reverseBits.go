package main

// LEEDCODE: 190

func reverseBits(num uint32) uint32 {

	reversed := uint32(0)
	power := uint32(31)

	for num != 0 {
		reversed += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return reversed
}
