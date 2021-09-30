package variadic

func Max(vals...int) int {
	if len(vals) == 0 {
		return 0
	}

	max := vals[0]

	for _, num := range vals {
		if num > max {
			max = num
		}
	}

	return max
}

func Min(vals...int) int {
	if len(vals) == 0 {
		return 0
	}

	min := vals[0]

	for _, num := range vals {
		if num < min {
			min = num
		}
	}

	return min
}
