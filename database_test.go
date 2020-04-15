package mtp

import "testing"

func TestDatabaseMockImplementsInterfaces(t *testing.T) {
	implTester := func(DatabaseTester) {}
	implConnecter := func(DatabaseConnecter) {}
	implTester(&DatabaseMock{})
	implConnecter(&DatabaseMock{})
}

func TestIsGoArticle(t *testing.T) {
	articlePython := Article{Title: "I prefer Python"}
	articleGo := Article{Title: "Go go gadget"}

	if isGoArticle(articlePython) || !isGoArticle(articleGo) {
		t.Fatalf("Wrong classification")
	}
}
