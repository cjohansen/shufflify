package shuffle

import (
	"math/rand"
	"sort"
)

type Item interface {
	GroupingKey(string) string
}

type GroupedItems map[string][]Item

type Bucket struct {
	label       string
	occurrences int
}

type byOccurrences []Bucket

func (s byOccurrences) Len() int {
	return len(s)
}

func (s byOccurrences) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byOccurrences) Less(i, j int) bool {
	return s[i].occurrences > s[j].occurrences
}

func IndexOf(haystack []string, needle string, start int) int {
	l := len(haystack)

	for i := start; i < start+l; i++ {
		index := i % l
		if haystack[index] == needle {
			return index
		}
	}

	return -1
}

func GroupBy(items []Item, fn func(Item) string) GroupedItems {
	res := make(GroupedItems)

	for _, item := range items {
		key := fn(item)

		if res[key] == nil {
			res[key] = make([]Item, 1)
			res[key][0] = item
		} else {
			res[key] = append(res[key], item)
		}
	}

	return res
}

func Distribute(distribution []string, bucket Bucket) []string {
	var stepSize float64 = float64(len(distribution)) / float64(bucket.occurrences)
	index := 0
	n := bucket.occurrences
	remainder := stepSize

	for {
		if n == 0 {
			return distribution
		}

		if remainder >= stepSize {
			index = IndexOf(distribution, "", index)
			n--
			remainder -= stepSize
			distribution[index] = bucket.label
		}

		index++
		remainder++
	}
}

func bucketsByOccurrences(grouped GroupedItems) []Bucket {
	res := make([]Bucket, len(grouped))
	i := 0

	for bucket, bucketItems := range grouped {
		res[i] = Bucket{bucket, len(bucketItems)}
		i++
	}

	sort.Sort(byOccurrences(res))
	return res
}

// Given a distribution and items grouped by the distribution keys, returns
// a collection of items:
//
// []string{"a", "b", "a", "c"}
// GroupedItems{"a": []Item{"A item 1", "A item 2"},
//              "b": []Item{"B item 1"}
//              "c": []Item{"C item 1"}}
// =>
// []Item{"A item 1", "B item 1", "A item 2", "C item 1"}
func ReifyDistribution(distribution []string, groups GroupedItems) []Item {
	items := make([]Item, len(distribution))

	for i, key := range distribution {
		items[i], groups[key] = groups[key][0], groups[key][1:]
	}

	return items
}

func DistributeBy(items []Item, fns []func(Item) string) []Item {
	if len(fns) == 0 {
		return items
	}

	distributedGrouped := make(GroupedItems)
	fn := fns[0]
	grouped := GroupBy(items, fn)
	buckets := bucketsByOccurrences(grouped)
	distribution := make([]string, len(items))

	for _, bucket := range buckets {
		distributedGrouped[bucket.label] = DistributeBy(grouped[bucket.label], fns[1:])
		Distribute(distribution, bucket)
	}

	return ReifyDistribution(distribution, distributedGrouped)
}

func AttributeAccessors(attrs []string) []func(Item) string {
	res := make([]func(Item) string, len(attrs))

	for i, attr := range attrs {
		res[i] = func(a string) func(Item) string {
			return func(i Item) string {
				return i.GroupingKey(a)
			}
		}(attr)
	}

	return res
}

func shuffle(src []Item) []Item {
	dest := make([]Item, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

func startRandomly(items []Item) []Item {
	index := rand.Intn(len(items))
	return append(items[index:], items[:index]...)
}

func ShuffleBy(items []Item, fns []func(Item) string) []Item {
	return startRandomly(DistributeBy(shuffle(items), fns))
}
