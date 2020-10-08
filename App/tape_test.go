package poker

import (
	"io/ioutil"
	"testing"
)

func TestTape_Writer(t *testing.T) {
	file, clean := createTempFile(t, "123456")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
