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
