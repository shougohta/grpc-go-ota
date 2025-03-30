package calc_test

import (
	"testing"

	calc "example.com/mod/test/calc"
)

// サブテスト
func TestMaxSubtests(t *testing.T) {
	t.Run("正の数", func(t *testing.T) {
		got := calc.Max(1, 2)
		want := 2
		if got != want {
			t.Errorf("Max(1, 2) == %d, want %d", got, want)
		}
	})
	t.Run("負の数", func(t *testing.T) {
		got := calc.Max(-1, -2)
		want := -1
		if got != want {
			t.Errorf("Max(-1, -2) == %d, want %d", got, want)
		}
	})
	t.Run("両方", func(t *testing.T) {
		got := calc.Max(1, -2)
		want := 1
		if got != want {
			t.Errorf("Max(1, -2) == %d, want %d", got, want)
		}
	})
}

// ベンチマークテスト：b.Nは繰り返し回数でよしなに決定される
func BenchmarkMax(b *testing.B) {
	// ログ出力
	b.Log(b.Name())

	// 整形ログ出力
	b.Logf("整形: %s", b.Name())

	for i := 0; i < b.N; i++ {
		calc.Max(1, 2)
	}
}
