package hello

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	t.Run("saying hello to people ", func(t *testing.T) {
		got := Hello("Chris", englishLanguageText)
		want := englishHelloPrefix + "Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World'", func(t *testing.T) {
		got := Hello("", englishLanguageText)
		want := englishHelloPrefix + "World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanishLanguageText)
		want := spanishHelloPrefix + "Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func TestGreetingPrefix(t *testing.T) {
	t.Run("Test Greeting Prifix", func(t *testing.T) {
		got := greetingPrefix(frenchLanguageText)
		want := frenchHelloPrefix
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
