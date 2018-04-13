package shuffle

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
