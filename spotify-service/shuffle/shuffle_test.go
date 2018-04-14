package shuffle

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
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

func (grouped GroupedCollection) len() int {
	sum := 0

	for _, value := range grouped {
		sum += len(value)
	}

	return sum
}

func TestGroupBy(t *testing.T) {
	coll := Collection{
		Track{artist: "Banana"},
		Track{artist: "Apple"},
		Track{artist: "Banana"},
		Track{artist: "Orange"},
		Track{artist: "Grapes"},
		Track{artist: "Apple"},
		Track{artist: "Apple"},
		Track{artist: "Grapes"},
		Track{artist: "Orange"},
	}

	grouped := GroupBy(coll, func(track Track) string {
		return track.artist
	})

	assert.Equal(t, 2, len(grouped["Banana"]))
	assert.Equal(t, 3, len(grouped["Apple"]))
	assert.Equal(t, 2, len(grouped["Orange"]))
	assert.Equal(t, 2, len(grouped["Grapes"]))
}

func TestDistribute(t *testing.T) {
	assert.Equal(
		t,
		[]string{"A", "", "", "", "", "", "", "", "", ""},
		Distribute(make([]string, 10), Bucket{"A", 1}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "", "", "A", "", "", "", ""},
		Distribute(make([]string, 10), Bucket{"A", 2}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "", "A", "", "", "A", "", ""},
		Distribute(make([]string, 10), Bucket{"A", 3}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "", "A", "", "A", "", "", "A", ""},
		Distribute(make([]string, 10), Bucket{"A", 4}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "", "A", "", "A", "", "A", ""},
		Distribute(make([]string, 10), Bucket{"A", 5}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "", "A", "", "A", "A", "", "A"},
		Distribute(make([]string, 10), Bucket{"A", 6}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "", "A", "A", "", "A", "A"},
		Distribute(make([]string, 10), Bucket{"A", 7}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "A", "A", "", "A", "A", "A"},
		Distribute(make([]string, 10), Bucket{"A", 8}),
	)

	assert.Equal(
		t,
		[]string{"A", "", "A", "A", "A", "A", "A", "A", "A", "A"},
		Distribute(make([]string, 10), Bucket{"A", 9}),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "C", "A", "", "A", "B", "A", "C"},
		Distribute([]string{"A", "B", "A", "", "A", "", "A", "B", "A", ""}, Bucket{"C", 2}),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "D", "A", "C", "A", "B", "A", "C"},
		Distribute([]string{"A", "B", "A", "D", "A", "", "A", "B", "A", ""}, Bucket{"C", 2}),
	)

	assert.Equal(
		t,
		[]string{"A", "B", "A", "C", "A", "C", "A", "B", "A", "D"},
		Distribute([]string{"A", "B", "A", "", "A", "", "A", "B", "A", "D"}, Bucket{"C", 2}),
	)
}

func TestDistributedShuffle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Log(DistributedShuffle(Collection{
		Track{artist: "Nirvana", genre: "Grunge", track: "Smells Like Teen Spirit", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Come as You Are", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Heart-Shaped Box", album: "In Utero"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Dive", album: "Incesticide"},
		Track{artist: "Pearl Jam", genre: "Grunge", track: "Once", album: "Ten"},
		Track{artist: "Pearl Jam", genre: "Grunge", track: "Daughter", album: "VS"},
		Track{artist: "Leviathan", genre: "Black_Metal", track: "Her Circle is the Noose", album: "True Traitor, True Whore"},
		Track{artist: "Leviathan", genre: "Black_Metal", track: "The Smoke of Their Torment", album: "Scar Sighted"},
		Track{artist: "Nile", genre: "Death_Metal", track: "Black Seeds of Vengeance", album: "Black Seeds of Vengeance"},
		Track{artist: "Deathspell Omega", genre: "Black_Metal", track: "Hetomaisa", album: "Si Circumspice"},
		Track{artist: "Execration", genre: "Death_Metal", track: "Eternal Recurrence", album: "Return to the Void"},
	}, []func(Track) string{
		func(c Track) string {
			return c.artist
		},
	}))

	t.Error("Oops")
}

func BenchmarkDistribute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distribute(make([]string, 10), Bucket{"A", 5})
	}
}
