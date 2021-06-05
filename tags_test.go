package language_test

import (
	"testing"

	"github.com/muonsoft/language"
	textlanguage "golang.org/x/text/language"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name          string
		tag1          textlanguage.Tag
		tag2          textlanguage.Tag
		expectedEqual bool
	}{
		{
			name:          "equal tags",
			tag1:          language.Russian,
			tag2:          language.Russian,
			expectedEqual: true,
		},
		{
			name:          "not equal tags",
			tag1:          language.English,
			tag2:          language.Russian,
			expectedEqual: false,
		},
		{
			name:          "equal by parent tags",
			tag1:          textlanguage.MustParse("ru-RU"),
			tag2:          language.Russian,
			expectedEqual: true,
		},
		{
			name:          "equal by parent tags (reverse)",
			tag1:          language.Russian,
			tag2:          textlanguage.MustParse("ru-RU"),
			expectedEqual: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isEqual := language.Equal(test.tag1, test.tag2)

			if isEqual != test.expectedEqual {
				t.Error("failed asserting that tags are equal")
			}
		})
	}
}
