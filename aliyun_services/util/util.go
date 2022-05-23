package util

import "math/rand"

func Get_rand_string() string {
	nn := make([]byte, 6)

	for i := 0; i < len(nn); i++ {
		t := rand.Intn(36)
		if t >= 10 {
			nn[i] = byte('a' + t - 10)
		} else {
			nn[i] = byte('0' + t)
		}
	}
	return string(nn)
}
