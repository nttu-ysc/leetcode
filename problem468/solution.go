package problem468

import "strconv"

func validIPAddress(queryIP string) string {
	if validIPv4(queryIP) {
		return "IPv4"
	}
	if validIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(ip string) bool {
	var count int
	var tmp string
	ip += "."
	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			if len(tmp) == 0 {
				return false
			}
			count++
			if tmp[0] == '0' && len(tmp) > 1 {
				return false
			}
			d, _ := strconv.Atoi(tmp)
			if d < 0 || d > 255 {
				return false
			}
			tmp = ""
		} else if ip[i] >= '0' && ip[i] <= '9' {
			tmp += string(ip[i])
		} else {
			return false
		}
	}
	return count == 4
}

func validIPv6(ip string) bool {
	var count int
	var tmp string
	ip += ":"
	for i := 0; i < len(ip); i++ {
		if ip[i] == ':' {
			if len(tmp) < 1 || len(tmp) > 4 {
				return false
			}
			count++
			tmp = ""
		} else if (ip[i] >= '0' && ip[i] <= '9') || (ip[i] >= 'a' && ip[i] <= 'f') || (ip[i] >= 'A' && ip[i] <= 'F') {
			tmp += string(ip[i])
		} else {
			return false
		}
	}
	return count == 8
}
