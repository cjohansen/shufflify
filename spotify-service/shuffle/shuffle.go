package shuffle

func IndexOf(coll []string, item string, start int) int {
	for i := start; i < len(coll); i++ {
		if coll[i] == item {
			return i
		}
	}

	for i := 0; i < start; i++ {
		if coll[i] == item {
			return i
		}
	}

	return -1
}

func GroupBy(coll []map[string]string, key string) map[string][]map[string]string {
	res := make(map[string][]map[string]string)

	for _, item := range coll {
		if res[item[key]] == nil {
			res[item[key]] = make([]map[string]string, 1)
			res[item[key]][0] = item
		} else {
			res[item[key]] = append(res[item[key]], item)
		}
	}

	return res
}

func Distribute(distribution []string, bucket string, items int) []string {
	var stepSize float64 = float64(len(distribution)) / float64(items)
	index := 0
	n := items
	remainder := stepSize

	for {
		if n == 0 {
			return distribution
		}

		if remainder >= stepSize {
			index = IndexOf(distribution, "", index)
			n--
			remainder -= stepSize
			distribution[index] = bucket
		}

		index++
		remainder++
	}
}
