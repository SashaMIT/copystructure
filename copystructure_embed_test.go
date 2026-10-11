package copystructure_test

import (
	"testing"

	"github.com/mitchellh/copystructure"
)

func TestCopy_exportedFieldOfUnexportedEmbed(t *testing.T) {
	type inner struct {
		Field string
		Tags  []string
		priv  string
	}
	type mid struct {
		inner
	}
	type outer struct {
		mid
		Name string
	}

	in := outer{
		mid:  mid{inner: inner{Field: "111", Tags: []string{"a"}, priv: "secret"}},
		Name: "n",
	}
	copied, err := copystructure.Copy(in)
	if err != nil {
		t.Fatal(err)
	}
	got := copied.(outer)
	if got.Field != "111" {
		t.Fatalf("Field=%q", got.Field)
	}
	if got.Name != "n" {
		t.Fatalf("Name=%q", got.Name)
	}
	if len(got.Tags) != 1 || got.Tags[0] != "a" {
		t.Fatalf("Tags=%v", got.Tags)
	}
	if got.priv != "" {
		t.Fatalf("priv=%q", got.priv)
	}

	in.Tags[0] = "changed"
	if got.Tags[0] != "a" {
		t.Fatal("tags were aliased")
	}
}
