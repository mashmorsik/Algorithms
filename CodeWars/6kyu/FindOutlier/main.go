package main

func FindOutlier(integers []int) int {
	var arrEven []int
	var arrOdd []int
	for _, num := range integers {
		if num%2 == 0 {
			arrEven = append(arrEven, num)
		} else {
			arrOdd = append(arrOdd, num)
		}
	}
	if len(arrEven) == 1 {
		return arrEven[0]
	} else {
		return arrOdd[0]
	}
}

func FindOutlierFast(integers []int) int {
	var arrEven []int
	var arrOdd []int
	for _, num := range integers {
		if num%2 == 0 {
			arrEven = append(arrEven, num)
		} else {
			arrOdd = append(arrOdd, num)
		}
		if len(arrEven) > 1 && len(arrOdd) == 1 {
			return arrOdd[0]
		}
		if len(arrOdd) > 1 && len(arrEven) == 1 {
			return arrEven[0]
		}
	}
	return 0
}

func FindOutlierFaster(integers []int) int {
	var arrEven, arrOdd, countEven, countOdd int
	for _, num := range integers {
		if num%2 == 0 {
			arrEven = num
			countEven++
		} else {
			arrOdd = num
			countOdd++
		}
		if countEven > 1 && countOdd == 1 {
			return arrOdd
		}
		if countOdd > 1 && countEven == 1 {
			return arrEven
		}
	}
	return 0
}
