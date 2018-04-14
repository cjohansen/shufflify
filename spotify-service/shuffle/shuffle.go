package shuffle

import (
	"bytes"
	"math/rand"
	"sort"
)

type Track struct {
	artist string
	album  string
	track  string
	genre  string
}

type Collection []Track

type GroupedCollection map[string]Collection

type Bucket struct {
	label       string
	occurrences int
}

func (coll Collection) String() string {
	var buffer bytes.Buffer

	for _, track := range coll {
		buffer.WriteString("\n")
		buffer.WriteString(track.String())
	}

	return buffer.String()
}

func (track Track) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	buffer.WriteString("Artist: ")
	buffer.WriteString(track.artist)
	buffer.WriteString(", Album: ")
	buffer.WriteString(track.album)
	buffer.WriteString(", Track: ")
	buffer.WriteString(track.track)
	buffer.WriteString(", Genre: ")
	buffer.WriteString(track.genre)
	buffer.WriteString("}")
	return buffer.String()
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

func GroupBy(coll Collection, fn func(Track) string) GroupedCollection {
	res := make(GroupedCollection)

	for _, track := range coll {
		key := fn(track)

		if res[key] == nil {
			res[key] = make(Collection, 1)
			res[key][0] = track
		} else {
			res[key] = append(res[key], track)
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

func shuffle(src Collection) Collection {
	dest := make(Collection, len(src))
	perm := rand.Perm(len(src))

	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

func bucketsByOccurrences(grouped GroupedCollection) []Bucket {
	res := make([]Bucket, len(grouped))
	i := 0

	for bucket, bucketItems := range grouped {
		res[i] = Bucket{bucket, len(bucketItems)}
		i++
	}

	sort.Sort(byOccurrences(res))
	return res
}

func FlattenDistribution(distribution []string, groups map[string]Collection) Collection {
	coll := make(Collection, len(distribution))

	for i, key := range distribution {
		coll[i], groups[key] = groups[key][0], groups[key][1:]
	}

	return coll
}

func ShuffleBy(coll Collection, fns []func(Track) string) Collection {
	res := make([]string, len(coll))

	if len(fns) == 0 {
		return shuffle(coll)
	}

	distributedGrouped := make(map[string]Collection)
	fn := fns[0]
	grouped := GroupBy(coll, fn)
	buckets := bucketsByOccurrences(grouped)

	for _, bucket := range buckets {
		distributedGrouped[bucket.label] = ShuffleBy(grouped[bucket.label], fns[1:])
		Distribute(res, bucket)
	}

	return FlattenDistribution(res, distributedGrouped)
}

func StartRandomly(coll Collection) Collection {
	index := rand.Intn(len(coll))
	return append(coll[index:], coll[:index]...)
}

func DistributedShuffle(coll Collection, fns []func(Track) string) Collection {
	return StartRandomly(ShuffleBy(coll, fns))
}
