package ladderLength

import (
	"reflect"
	"testing"
)

func TestLadderLength(t *testing.T) {
	t.Run("test ladder length", func(t *testing.T) {
		beginWord := "hit"
		endWord := "cog"
		wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

		got := LadderLength(beginWord, endWord, wordList)
		want := 5

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
