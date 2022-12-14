package blackjack

import "testing"

func TestHandNumberOfAces(t *testing.T) {

	hand := Hand{
		Cards: []Card{
			Ace,
			Ace,
			Two,
			Two,
		},
	}

	var want = 2
	var got = hand.NumberOfAces()

	if want != got {
		t.Fatalf("got %d, wanted %d\n", got, want)
	}
}

func TestHandTotal(t *testing.T) {

	hand := Hand{
		Cards: []Card{
			Ace, // 1
			Ace, // 2
			Ace, // 3
			Ace, // 4
			Ace, // 5
			Ace, // 6
			Ace, // 7
			Ace, // 8
			Two, // 10
			Two, // 12
		},
	}

	var want = 12
	var got = hand.Total()

	if want != got {
		t.Fatalf("got %d, wanted %d\n", got, want)
	}

	hand = Hand{
		Cards: []Card{
			Ace,  // 11
			King, // 21
		},
	}

	want = 21
	got = hand.Total()

	if want != got {
		t.Fatalf("got %d, wanted %d\n", got, want)
	}
}

func TestHandBust(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			King,  // 10
			Queen, // 20
			Five,  // 25
			Ace,   // 26
		},
	}

	var want = true
	var got = hand.Bust()

	if want != got {
		t.Fatalf("got %t, wanted %t\n", got, want)
	}
}

func TestHandIsHard(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			Ace, // 1
			Ace, // 2
			Ace, // 3
			Ace, // 4
			Ace, // 5
			Ace, // 6
			Ace, // 7
			Ace, // 8
			Two, // 10
			Two, // 12
		},
	}

	var want = true
	var got = hand.IsHard()

	if want != got {
		t.Fatalf("got %t, wanted %t\n", got, want)
	}
}

func TestHandIsBlackjack(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			King, // 10
			Ace,  // 21
		},
	}

	var want = true
	var got = hand.IsBlackjack()

	if want != got {
		t.Fatalf("got %t, wanted %t\n", got, want)
	}

	hand = Hand{
		Cards: []Card{
			King, // 10
			Ten,  // 20
		},
	}

	want = false
	got = hand.IsBlackjack()

	if want != got {
		t.Fatalf("got %t, wanted %t\n", got, want)
	}
}

func TestHandEmpty(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			King, // 10
			Ace,  // 21
		},
	}

	want := 2
	got := len(hand.Cards)

	if want != got {
		t.Fatalf("got %d, wanted %d\n", got, want)
	}

	hand.Empty()
	want = 0
	got = len(hand.Cards)

	if want != got {
		t.Fatalf("got %d, wanted %d\n", got, want)
	}
}
