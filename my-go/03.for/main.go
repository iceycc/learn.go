package main

func main() {
	// 1
	for i := 0; i < 5; i++ {
		println(i, "golang3")
	}
	// 2
	j := 0
	for ; j < 5; j++ {
		println(j, "golang2")
	}
	// 3
	k := 0
	for k < 5 {
		println(k, "golang3")
		k++
	}

	// 4
	m := 0
	for m < 5 {
		m++
		if m%3 == 0 {
			continue
		}
		println(m, "golang4")
	}
}
