package crypto_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/scheiblingco/gofn/crypto"
)

func TestLoadPfx(t *testing.T) {
	data, err := os.ReadFile("test.pfx")
	if err != nil {
		t.Fatal(err)
	}

	cert, err := crypto.X509FromPkcs12(data, "qwerty123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cert.Leaf.Subject)
}

func TestLoadP12(t *testing.T) {
	data, err := os.ReadFile("test.p12")
	if err != nil {
		t.Fatal(err)
	}

	cert, err := crypto.X509FromPkcs12(data, "qwerty123")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cert.Leaf.Subject)
}
