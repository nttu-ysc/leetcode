package problem443

func compress(chars []byte) int {
	mem := chars[0]
	start := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == mem {
			continue
		}
		if i-start == 1 {
			start = i
			mem = chars[start]
			continue
		}
		count := int2Byte(i - start)
		chars = append(chars[:start], append([]byte{mem}, append(count, chars[i:]...)...)...)
		start += 1 + len(count)
		mem = chars[start]
		i = start
	}
	if len(chars)-start > 1 {
		count := int2Byte(len(chars) - start)
		chars = append(chars[:start], append([]byte{mem}, count...)...)
	}

	return len(chars)
}

func int2Byte(d int) []byte {
	var res []byte
	for i := d; i > 0; i /= 10 {
		res = append([]byte{byte(i%10) + '0'}, res...)
	}
	return res
}
