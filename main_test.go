package printer

import "testing"

// run all of the tests -> | go test -v ./...

// func TestPrintHello(t *testing.T) {
// 	got := PrintHello("Igor")
// 	expected := "Hello, Igor!"

// 	if got != expected {
// 		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
// 	}
// }

// func TestPrintHello(t *testing.T) {
//     got := PrintHello("Igor")
//     expected := "Hello, Igor!"

//     if got != expected {
//         t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
//     }

//     for name := range names {
//         if name != "Igor" {
//             t.Fatalf(`got %q, want %q`, name, "Igor")
//         } else {
//             break
//         }
//     }
// }

func TestPrintHello(t *testing.T) { //go test -v -run=^TestPrintHello$
	names["Petr"] = "Petr"
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}

	for name := range names {
		if name != "Igor" {
			t.Fatalf(`got %q, want %q`, name, "Igor")
		} else {
			break
		}
	}
}

func TestPrintHelloIvan(t *testing.T) { // go test -v -run=TestPrintHelloIvan
	got := PrintHello("Ivan")
	expected := "Hello, Ivan!"

	if got != expected {
		t.Fatalf(`PrintHello("Ivan") = %q, want %q`, got, expected)
	}
}
