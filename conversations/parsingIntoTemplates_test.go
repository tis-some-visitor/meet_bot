package conversations

import (
	"testing"
)

func TestCapitalizingDestinations(t *testing.T) {

	t.Run("владивосток хабаровск ростов-на-дону", func(t *testing.T) {

		got := CapitalizingDestinations("владивосток хабаровск ростов-на-дону")
		want := "Владивосток Хабаровск Ростов-на-Дону"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("кот-д'ивуар|сент-винсент и гренадины|сша", func(t *testing.T) {
		got := CapitalizingDestinations("кот-д'ивуар|сент-винсент и гренадины|сша")
		want := "Кот-д'Ивуар, Сент-Винсент и Гренадины, США"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Кот-д'Ивуар, Сент-Винсент и Гренадины, США", func(t *testing.T) {
		got := CapitalizingDestinations("Кот-д'Ивуар, Сент-ВИНСЕНТ и Гренадины, США")
		want := "Кот-д'Ивуар, Сент-Винсент и Гренадины, США"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	/*
		t.Run("dates are long, so have weekends", func(t *testing.T) {
			date1, _ := time.Parse("02.01.2006", "20.08.2022")
			date2, _ := time.Parse("02.01.2006", "30.08.2022")

			got := HasWeekends(date1, date2)
			want := int64(30)
			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})*/
}
