package envar

import (
	"testing"
)

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHashID(t *testing.T) {

	t.Run("123123123", func(t *testing.T) {
		got, _ := HashID(5109919104)
		want := "0PuJ_DoKXekSitiI5tTNpmvRSR_NYZ9SeiWz_ySimUY=1"
		if got != want {
			t.Errorf("got %s,\n want %s", got, want)
		}
	})
}

func BenchmarkHashID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashID(123123123)
	}
}
