package main

func FindSimilarity(left, right []int) int {
	similarityMap := make(map[int]int)
	similarity := 0

	for _, v := range left {
		if val, exist := similarityMap[v]; exist {
			if val == 0 {
				similarityMap[v] = 2
			} else {
				similarityMap[v] = val + 1
			}
		} else {
			similarityMap[v] = 0
		}
	}

	for _, v := range right {
		if val, exist := similarityMap[v]; exist {
			similarityMap[v] = val + 1
		}
	}

	for k, v := range similarityMap {
		similarity += k * v
	}
	return similarity
}