package packr

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	st "github.com/golang-migrate/migrate/v4/source/testing"
)

func Test(t *testing.T) {
	box := packr.New("migrations", "./testdata")

	d, err := WithInstance(box, &Packr{})
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}

func TestWithInstance(t *testing.T) {
	box := packr.New("migrations", "./testdata")

	_, err := WithInstance(box, &Packr{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpen(t *testing.T) {
	p := &Packr{}
	_, err := p.Open("")
	if err == nil {
		t.Fatal("expected err, because it's not implemented yet")
	}
}

func TestClose(t *testing.T) {
	box := packr.New("migrations", "./testdata")

	d, err := WithInstance(box, &Packr{})
	if err != nil {
		t.Fatal(err)
	}

	err = d.Close()
	if err != nil {
		t.Fatal(err)
	}
}
