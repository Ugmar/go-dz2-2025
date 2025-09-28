package main

import (
	task1 "dz2/internal/task_1"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnimalFeeding(t *testing.T) {
	t.Parallel()
	t.Run("Обычный случай: все типы животных", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `q--3wer4301-4+r_e**
	   dog d1
	   cat c1
	   cat c2
	   bird b1
	   end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 4)
		assert.Equal(t, "d1 qwe", results[0])
		assert.Equal(t, "c1 0", results[1])
		assert.Equal(t, "c2 1", results[2])
		assert.Equal(t, "b1 -_", results[3])
	})

	t.Run("Пустая кормушка", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `
	   dog d1
	   cat c1
	   end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 2)
		assert.Equal(t, "d1 ", results[0])
		assert.Equal(t, "c1 ", results[1])
	})

	t.Run("Нет животных", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `q--3wer4301-4+r_e**
	   end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Empty(t, results)
	})

	t.Run("Только один тип животных", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `abc123!@#
	   cat cat1
	   cat cat2
	   cat cat3
	   end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "cat1 1", results[0])
		assert.Equal(t, "cat2 2", results[1])
		assert.Equal(t, "cat3 3", results[2])
	})
	t.Run("Unicode символы в кормушке", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `Привет1234🎉🐱!abc
dog d1
cat c1
bird b1
end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "d1 Пвc", results[0]) // буквы
		assert.Equal(t, "c1 2", results[1])   // цифры
		assert.Equal(t, "b1 🐱", results[2])   // спецсимволы
	})

	t.Run("Разные регистры названий животных", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `q--3wer4301-4+r_e**
	   Dog d1
	   CAT c1
	   cAt c2
	   biRD b1
	   end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 4)
		assert.Equal(t, "d1 qwe", results[0])
		assert.Equal(t, "c1 0", results[1])
		assert.Equal(t, "c2 1", results[2])
		assert.Equal(t, "b1 -_", results[3])
	})

	t.Run("Пустая строка корма", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `

		dog d1
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "d1 ", results[0])
		assert.Equal(t, "c1 ", results[1])
		assert.Equal(t, "b1 ", results[2])
	})

	t.Run("Много животных, мало корма", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `ab12

		dog d1
		dog d2
		dog d3
		cat c1
		cat c2
		bird b1
		bird b2
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 7)
		// Корм "ab12", очередь: d1, d2, d3, c1, c2, b1, b2
		// d1: a (съел), d2: b (съел), d3: 1 (не съел), c1: 2 (съел)
		assert.Equal(t, "d1 a", results[0])
		assert.Equal(t, "d2 b", results[1])
		assert.Equal(t, "d3 ", results[2])
		assert.Equal(t, "c1 2", results[3])
		assert.Equal(t, "c2 ", results[4])
		assert.Equal(t, "b1 ", results[5])
		assert.Equal(t, "b2 ", results[6])
	})

	t.Run("Только специальные символы", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `!@#$%^&*()_+-=

		dog d1
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "d1 ", results[0])     // собака не ест спецсимволы
		assert.Equal(t, "c1 ", results[1])     // кошка не ест спецсимволы
		assert.Equal(t, "b1 #^(+", results[2]) // птица съела все
	})

	t.Run("Только цифры", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `1234567890

		dog d1
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "d1 ", results[0])
		assert.Equal(t, "c1 258", results[1])
		assert.Equal(t, "b1 ", results[2])
	})

	t.Run("Только буквы", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `abcABCaBC

		dog d1
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		assert.Equal(t, "d1 aAa", results[0]) // собака съела все
		assert.Equal(t, "c1 ", results[1])    // кошка не ест буквы
		assert.Equal(t, "b1 ", results[2])    // птица не ест буквы
	})

	t.Run("Порядок вывода соответствует порядку ввода", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `a1b2c3

		bird b1
		cat c1
		dog d1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		// Очередь: b1, c1, d1
		// b1: a (не съела), c1: 1 (съела), d1: b (съел), b1: 2 (не съела), c1: c (не съела), d1: 3 (не съел)
		assert.Equal(t, "b1 ", results[0])
		assert.Equal(t, "c1 1", results[1])
		assert.Equal(t, "d1 b", results[2])
	})

	t.Run("Большое количество животных", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `abc123!@#

		dog dog1
		cat cat1
		bird bird1
		dog dog2
		cat cat2
		bird bird2
		dog dog3
		cat cat3
		bird bird3
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 9)
		// Очередь: dog1, cat1, bird1, dog2, cat2, bird2, dog3, cat3, bird3
		// dog1: a, cat1: 1, bird1: !, dog2: b, cat2: 2, bird2: @, dog3: c, cat3: 3, bird3: #
		assert.Equal(t, "dog1 a", results[0])
		assert.Equal(t, "cat1 ", results[1])
		assert.Equal(t, "bird1 ", results[2])
		assert.Equal(t, "dog2 ", results[3])
		assert.Equal(t, "cat2 2", results[4])
		assert.Equal(t, "bird2 ", results[5])
		assert.Equal(t, "dog3 ", results[6])
		assert.Equal(t, "cat3 ", results[7])
		assert.Equal(t, "bird3 #", results[8])
	})

	t.Run("Один символ корма", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `a

		dog d1
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		// d1: a (съел), c1: -, b1: -
		assert.Equal(t, "d1 a", results[0])
		assert.Equal(t, "c1 ", results[1])
		assert.Equal(t, "b1 ", results[2])
	})

	t.Run("Один символ корма, но никто не сьел", func(t *testing.T) {
		t.Parallel()
		// Given
		input := `a

		cat c0
		cat c1
		bird b1
		end`

		// When
		results := task1.AnimalFeeding(strings.NewReader(input))

		// Then
		require.Len(t, results, 3)
		// d1: a (съел), c1: -, b1: -
		assert.Equal(t, "c0 ", results[0])
		assert.Equal(t, "c1 ", results[1])
		assert.Equal(t, "b1 ", results[2])
	})

}
