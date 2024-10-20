package printer

import "testing"

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

func TestPrintHello(t *testing.T) {
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
