package counter

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	Increment()

	if counter != 1 {
		t.Error("counter is expected to be incremented")
	}
}

// func TestSum(t *testing.T) {
// 	// набор тестов
// 	cases := []struct {
// 		// имя теста
// 		name string
// 		// значения на вход тестируемой функции
// 		values []int
// 		// желаемый результат
// 		want int
// 	}{
// 		// тестовые данные № 1
// 		{
// 			name: "positive values",
// 			values: []int{1, 2, 3},
// 			want: 6,
// 		},
// 		// тестовые данные № 2
// 		{
// 			name: "mixed values",
// 			values: []int{-1, 2, -3},
// 			want: -2,
// 		},
// 	}
// 	// перебор всех тестов
// 	for _, tc := range cases {
// 		tc := tc
// 		// запуск отдельного теста
// 		t.Run(tc.name, func(t *testing.T) {
// 			// тестируем функцию Sum
// 			got := Sum(tc.values...)
// 			// проверим полученное значение
// 			if got != tc.want {
// 					t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
// 			}
// 		})
// 	}
// }

func TestSum(t *testing.T) {
	a, b := 1, 4
	result := Sum(a, b)

	if (a + b) != result {
		t.Fatalf(`Sum(%d, %d) = %d, want %d`, a, b, result, (a + b))
	}
}

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
	{0, "zero"},
	{99, "long"},
	{101, "very long"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

func TestMultiply(t *testing.T) {
	a, b := 1, 4
	result := Multiply(a, b)

	if (a * b) != result {
		t.Fatalf(`Multiply(%d, %d) = %d, want %d`, a, b, result, (a * b))
	}
}

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"No vowels", "bcdfghjklmnpqrstvwxyz", "bcdfghjklmnpqrstvwxyz"},
		{"Only vowels", "aeiouAEIOU", ""},
		{"Mixed string", "hello world", "hll wrld"},
		{"Empty string", "", ""},
		{"With capital letters", "GoLang", "GLng"},
		{"Vowels at the end", "testee", "tst"},
		{"Vowels at the beginning", "oogee", "g"},
		{"Numbers and punctuation", "123 hello!", "123 hll!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeleteVowels(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
		err      error
	}{
		{"Valid ASCII", []byte("hello"), 5, nil},
		{"Valid UTF-8", []byte("Привет"), 6, nil},                                // русский текст
		{"Empty string", []byte(""), 0, nil},                                     // пустая строка
		{"Invalid UTF-8", []byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},           // некорректная последовательность байтов
		{"Mixed valid and invalid", []byte("hello\xffworld"), 0, ErrInvalidUTF8}, // смешанные корректные и некорректные байты
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetUTFLength(tt.input)
			if result != tt.expected || err != tt.err {
				t.Errorf("expected (result: %d, err: %v), got (result: %d, err: %v)", tt.expected, tt.err, result, err)
			}
		})
	}
}
