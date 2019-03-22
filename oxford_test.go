package oxforddict_test

import (
	"testing"

	"github.com/suguanyang/oxforddict"
)

func TestRequestOxfordDictionaries(t *testing.T) {
	id := oxforddict.RequestOxfordDictionaries([]string{"waga"}, "en")
	expected := "https://od-api.oxforddictionaries.com/api/v1"
	if id != expected {
		t.Errorf("RequestOxfordDictionaries() = %s; want %s", id, expected)
	}
}

// var _ = Describle("test api", func() {
// 	id := oxforddict.RequestOxfordDictionaries("waga", "en")
// 	expected := "febc0f05"
// 	Expecte(id).To(Equal(expected))
// })
