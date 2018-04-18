package shuffle

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type Track struct {
	artist string
	album  string
	track  string
	genre  string
}

func (t Track) GroupingKey(attribute string) string {
	switch attribute {
	case "genre":
		return t.genre
	case "artist":
		return t.artist
	case "album":
		return t.album
	default:
		return t.track
	}
}

func genres(tracks []Item) []string {
	res := make([]string, len(tracks))

	for i, track := range tracks {
		res[i] = track.(Track).genre
	}

	return res
}

func artists(tracks []Item) []string {
	res := make([]string, len(tracks))

	for i, track := range tracks {
		res[i] = track.(Track).artist
	}

	return res
}

func albums(tracks []Item) []string {
	res := make([]string, len(tracks))

	for i, track := range tracks {
		res[i] = track.(Track).album
	}

	return res
}

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

func (g GroupedItems) len() int {
	sum := 0

	for _, items := range g {
		sum += len(items)
	}

	return sum
}

func TestGroupBy(t *testing.T) {
	coll := []Item{
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

	grouped := GroupBy(coll, func(i Item) string {
		return i.GroupingKey("artist")
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

func playlist() []Item {
	return []Item{
		Track{artist: "Nirvana", genre: "Grunge", track: "Smells Like Teen Spirit", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Come as You Are", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Heart-Shaped Box", album: "In Utero"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Something in the Way", album: "Nevermind"},
		Track{artist: "Pearl Jam", genre: "Grunge", track: "Once", album: "Ten"},
		Track{artist: "Pearl Jam", genre: "Grunge", track: "Daughter", album: "VS"},
		Track{artist: "Leviathan", genre: "BM", track: "Dawn Vibration", album: "Scar Sighted"},
		Track{artist: "Leviathan", genre: "BM", track: "The Smoke of Their Torment", album: "Scar Sighted"},
		Track{artist: "Nile", genre: "DM", track: "Black Seeds of Vengeance", album: "Black Seeds of Vengeance"},
		Track{artist: "Deathspell Omega", genre: "BM", track: "Abscission", album: "Paracletus"},
		Track{artist: "Execration", genre: "DM", track: "Eternal Recurrence", album: "Return to the Void"},
	}
}

func TestDistributeByArtist(t *testing.T) {
	assert.Equal(
		t,
		[]string{"Nirvana", "Pearl Jam", "Nirvana", "Nirvana", "Pearl Jam", "Nirvana"},
		artists(DistributeBy(playlist()[:6], []func(Item) string{
			func(c Item) string {
				return c.GroupingKey("artist")
			},
		})),
	)
}

func TestDistributeByArtistAttributeAccessor(t *testing.T) {
	assert.Equal(
		t,
		[]string{"Nirvana", "Pearl Jam", "Nirvana", "Nirvana", "Pearl Jam", "Nirvana"},
		artists(DistributeBy(playlist()[:6], AttributeAccessors([]string{"artist"}))),
	)
}

func TestDistributeByGenreArtist(t *testing.T) {
	byGenreArtist := DistributeBy(playlist(), []func(Item) string{
		func(c Item) string {
			return c.GroupingKey("genre")
		},
		func(c Item) string {
			return c.GroupingKey("artist")
		},
	})

	assert.Equal(
		t,
		[]string{"Grunge", "BM", "Grunge", "DM", "Grunge", "BM", "Grunge", "DM", "Grunge", "BM", "Grunge"},
		genres(byGenreArtist),
	)

	assert.Equal(
		t,
		[]string{
			"Nirvana",
			"Leviathan",
			"Pearl Jam",
			"Nile",
			"Nirvana",
			"Deathspell Omega",
			"Nirvana",
			"Execration",
			"Pearl Jam",
			"Leviathan",
			"Nirvana",
		},
		artists(byGenreArtist),
	)
}

func TestDistributeByGenreArtistAttributeAccessors(t *testing.T) {
	byGenreArtist := DistributeBy(playlist(), AttributeAccessors([]string{"genre", "artist"}))

	assert.Equal(
		t,
		[]string{
			"Nirvana",
			"Leviathan",
			"Pearl Jam",
			"Nile",
			"Nirvana",
			"Deathspell Omega",
			"Nirvana",
			"Execration",
			"Pearl Jam",
			"Leviathan",
			"Nirvana",
		},
		artists(byGenreArtist),
	)
}

func TestDistributeByGenreArtistAlbum(t *testing.T) {
	assert.Equal(
		t,
		[]string{
			"Nevermind",
			"Scar Sighted",
			"Ten",
			"Black Seeds of Vengeance",
			"In Utero",
			"Paracletus",
			"Nevermind",
			"Return to the Void",
			"VS",
			"Scar Sighted",
			"Nevermind",
		},
		albums(DistributeBy(playlist(), []func(Item) string{
			func(c Item) string {
				return c.GroupingKey("genre")
			},
			func(c Item) string {
				return c.GroupingKey("artist")
			},
			func(c Item) string {
				return c.GroupingKey("album")
			},
		})),
	)
}

func TestDistributeByGenreArtistAlbumAttributeAccessor(t *testing.T) {
	assert.Equal(
		t,
		[]string{
			"Nevermind",
			"Scar Sighted",
			"Ten",
			"Black Seeds of Vengeance",
			"In Utero",
			"Paracletus",
			"Nevermind",
			"Return to the Void",
			"VS",
			"Scar Sighted",
			"Nevermind",
		},
		albums(DistributeBy(playlist(), AttributeAccessors([]string{"genre", "artist", "album"}))),
	)
}

func TestDistributedShuffle(t *testing.T) {
	rand.Seed(1)

	playlist := []Item{
		Track{artist: "Nirvana", genre: "Grunge", track: "Smells Like Teen Spirit", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Come as You Are", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Something in the Way", album: "Nevermind"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Heart-Shaped Box", album: "In Utero"},
		Track{artist: "Nirvana", genre: "Grunge", track: "Serve the Servants", album: "In Utero"},
		Track{artist: "Nirvana", genre: "Grunge", track: "About a Girl", album: "Bleach"},
		Track{artist: "Leviathan", genre: "BM", track: "Dawn Vibration", album: "Scar Sighted"},
		Track{artist: "Leviathan", genre: "BM", track: "The Smoke of Their Torment", album: "Scar Sighted"},
		Track{artist: "Nile", genre: "DM", track: "Black Seeds of Vengeance", album: "Black Seeds of Vengeance"},
	}

	assert.Equal(
		t,
		[]string{
			"Nevermind",
			"Scar Sighted",
			"In Utero",
			"Nevermind",
			"Scar Sighted",
			"In Utero",
			"Nevermind",
			"Black Seeds of Vengeance",
			"Bleach",
		},
		albums(ShuffleBy(playlist, AttributeAccessors([]string{"genre", "artist", "album"}))),
	)
}

func BenchmarkDistribute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distribute(make([]string, 10), Bucket{"A", 5})
	}
}

var benchPlaylist []Item = []Item{
	Track{artist: "Nirvana", genre: "Grunge", track: "Smells Like Teen Spirit", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "In Bloom", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Come As You Are", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Breed", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Lithium", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Polly", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Territorial Pissings", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Drain You", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Lounge Act", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "On a Plain", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Something in the Way", album: "Nevermind"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Serve the Servants", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Scentless Apprentice", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Heart-Shaped Box", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Rape Me", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Frances Farmer Will Have Her Revenge on Seattle", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Dumb", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Very Ape", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Milk It", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Pennyroyal Tea", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Radio Friendly Unit Shifter", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "Tourette's", album: "In Utero"},
	Track{artist: "Nirvana", genre: "Grunge", track: "All Apologies", album: "In Utero"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Once", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Even Flow", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Alive", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Why Go", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Black", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Jeremy", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Oceans", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Porch", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Garden", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Deep", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Release", album: "Ten"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Go", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Animal", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Daughter", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Glorified G", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Dissident", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "W.M.A", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Blood", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Rearviewmirror", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Rats", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Elderly Woman Behind The Counter in a Small Town", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Leash", album: "VS"},
	Track{artist: "Pearl Jam", genre: "Grunge", track: "Indifference", album: "VS"},
	Track{artist: "Leviathan", genre: "BM", track: "Intro", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "The Smoke of their Torment", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Dawn Vibration", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Gardens of Coprolite", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Wicked Fields of Calm", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Within Thrall", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "A Veil is Lifted", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Scar Sighted", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "All Tongues Toward", album: "Scar Sighted"},
	Track{artist: "Leviathan", genre: "BM", track: "Aphonos", album: "Scar Sighted"},
	Track{artist: "Nile", genre: "DM", track: "Invocation of the Gate of Aat-Ankh-Es-Amenti", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Black Seeds of Vengeance", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Defiling the Gates of Ishtar", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "The Black Flame", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Libation Unto the Shades Who Lurk in the Shadows of the Temple of Anhur", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Masturbating the War God", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Multitude of Foes", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Chapter for Transforming Into a Snake", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Nas Akhu Khan She En Asbiu", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "To Dream of Ur", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "The Nameless City of the Accused", album: "Black Seeds of Vengeance"},
	Track{artist: "Nile", genre: "DM", track: "Khetti Satha Shemsu", album: "Black Seeds of Vengeance"},
}

func BenchmarkShuffleBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShuffleBy(benchPlaylist, AttributeAccessors([]string{"genre", "artist", "album"}))
	}
}
