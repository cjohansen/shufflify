package shuffle

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/quick"
)

func TestIndexOf(t *testing.T) {
	coll := []string{"Some", "String", "Is", "Repeated", "Some", "Times", ""}

	assert.Equal(t, IndexOf(coll, "Some", 0), 0)
	assert.Equal(t, IndexOf(coll, "String", 0), 1)
	assert.Equal(t, IndexOf(coll, "Is", 0), 2)
	assert.Equal(t, IndexOf(coll, "Some", 1), 4)
	assert.Equal(t, IndexOf(coll, "Times", 1), 5)
	assert.Equal(t, IndexOf(coll, "Some", 4), 4)
	assert.Equal(t, IndexOf(coll, "Times", 4), 5)
	assert.Equal(t, IndexOf(coll, "", 0), 6)
	assert.Equal(t, IndexOf(coll, "Times", 6), 5)
}

func groupedItemCount(grouped map[string][]map[string]string) int {
	sum := 0

	for _, value := range grouped {
		sum += len(value)
	}

	return sum
}

func TestGroupBy(t *testing.T) {
	coll := []map[string]string{
		map[string]string{"id": "1", "fruit": "Banana"},
		map[string]string{"id": "2", "fruit": "Apple"},
		map[string]string{"id": "3", "fruit": "Banana"},
		map[string]string{"id": "4", "fruit": "Orange"},
		map[string]string{"id": "5", "fruit": "Grapes"},
		map[string]string{"id": "6", "fruit": "Apple"},
		map[string]string{"id": "7", "fruit": "Apple"},
		map[string]string{"id": "8", "fruit": "Grapes"},
		map[string]string{"id": "9", "fruit": "Orange"},
	}

	grouped := GroupBy(coll, "fruit")

	assert.Equal(t, 2, len(grouped["Banana"]))
	assert.Equal(t, 3, len(grouped["Apple"]))
	assert.Equal(t, 2, len(grouped["Orange"]))
	assert.Equal(t, 2, len(grouped["Grapes"]))
}

func TestGroupByProperties(t *testing.T) {
	f := func(coll []map[string]string, key string) bool {
		grouped := GroupBy(coll, key)
		return len(coll) == groupedItemCount(grouped)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDistribute(t *testing.T) {
	assert.Equal(
		t,
		[]string{"A", "", "", "", "", "", "", "", "", ""},
		Distribute(make([]string, 10), "A", 1),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "", "", "A", "", "", "", ""},
		Distribute(make([]string, 10), "A", 2),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "", "A", "", "", "A", "", ""},
		Distribute(make([]string, 10), "A", 3),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "A", "", "A", "", "", "A", ""},
		Distribute(make([]string, 10), "A", 4),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "", "A", "", "A", "", "A", ""},
		Distribute(make([]string, 10), "A", 5),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "", "A", "", "A", "A", "", "A"},
		Distribute(make([]string, 10), "A", 6),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "", "A", "A", "", "A", "A"},
		Distribute(make([]string, 10), "A", 7),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "A", "A", "", "A", "A", "A"},
		Distribute(make([]string, 10), "A", 8),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "A", "A", "A", "A", "A", "A"},
		Distribute(make([]string, 10), "A", 9),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "C", "A", "", "A", "B", "A", "C"},
		Distribute([]string{"A", "B", "A", "", "A", "", "A", "B", "A", ""}, "C", 2),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "D", "A", "C", "A", "B", "A", "C"},
		Distribute([]string{"A", "B", "A", "D", "A", "", "A", "B", "A", ""}, "C", 2),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "C", "A", "C", "A", "B", "A", "D"},
		Distribute([]string{"A", "B", "A", "", "A", "", "A", "B", "A", "D"}, "C", 2),
	)
}

func BenchmarkDistribute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distribute(make([]string, 10), "A", 5)
	}
}

func BenchmarkDistributeRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DistributeRecursive(make([]string, 10), "A", 5)
	}
}
