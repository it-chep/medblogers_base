package pipe

import (
	"context"
	"fmt"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testObject struct {
}

func (t *testObject) Do(err error) func(_ context.Context, value int) (int, error) {
	return func(_ context.Context, value int) (int, error) {
		return value * value, err
	}
}

func (t *testObject) Break() func(_ context.Context, value int) (int, error) {
	return func(_ context.Context, value int) (int, error) {
		return value, Break()
	}
}

func TestPipe(t *testing.T) {
	t.Parallel()

	t.Run("успешный ответ", func(t *testing.T) {
		t.Parallel()
		test := &testObject{}
		result, err := With(test.Do(nil)).With(test.Do(nil)).Run(context.Background(), 2).Get()

		assert.NoError(t, err)
		assert.Equal(t, 16, result)
	})

	t.Run("выход из пайп", func(t *testing.T) {
		t.Parallel()
		test := &testObject{}
		testErr := fmt.Errorf("idk")
		result, err := With(test.Do(nil)).With(test.Break()).With(test.Do(testErr)).Run(context.Background(), 2).Get()

		assert.NoError(t, err)
		assert.Equal(t, 4, result)
	})

	t.Run("ошибка в одной из секций", func(t *testing.T) {
		t.Parallel()
		test := &testObject{}
		result, err := With(test.Do(fmt.Errorf("err"))).
			With(test.Do(nil)).Run(context.Background(), 2).
			Anyway(context.Background(), func(_ context.Context, _ int, err error) {
				require.EqualError(t, err, "err")
			}).Get()

		assert.EqualError(t, err, "err")
		assert.Equal(t, 4, result)
	})

	t.Run("проверка фильтрации", func(t *testing.T) {
		t.Parallel()

		vals := IntWrappers{{}, {}, {}, {}, {}, {}}

		result, err :=
			With(incWrappers(nil)).
				OverFiltered(filterEvenFunc(nil),
					With(incWrappers(nil)),
				).
				Run(context.Background(), vals).Get()

		assert.NoError(t, err)
		assert.Equal(t, IntWrappers{{2}, {1}, {2}, {1}, {2}, {1}}, result)
	})

	t.Run("проверка отсутствия вызова подпайпа, если фильтр ничего не вернул", func(t *testing.T) {
		t.Parallel()

		vals := IntWrappers{{}, {}, {}, {}, {}, {}}

		result, err :=
			With(incWrappers(nil)).
				OverFiltered(
					func(_ context.Context, value IntWrappers) (context.Context, IntWrappers, bool, error) {
						return context.Background(), value, true, nil
					},
					With(incWrappers(nil)),
				).
				Run(context.Background(), vals).Get()

		assert.NoError(t, err)
		assert.Equal(t, IntWrappers{{1}, {1}, {1}, {1}, {1}, {1}}, result)
	})

	t.Run("проверка фильтрации, ошибка в функции фильтрации", func(t *testing.T) {
		t.Parallel()

		vals := IntWrappers{{}, {}, {}, {}, {}, {}}

		result, err :=
			With(incWrappers(nil)).
				OverFiltered(filterEvenFunc(fmt.Errorf("filter error")),
					With(incWrappers(nil)),
				).
				Run(context.Background(), vals).Get()

		assert.Error(t, err)
		assert.Equal(t, IntWrappers{{1}, {1}, {1}, {1}, {1}, {1}}, result)
	})

	t.Run("проверка фильтрации, ошибка во внутреннем пайпе", func(t *testing.T) {
		t.Parallel()

		vals := IntWrappers{{}, {}, {}, {}, {}, {}}

		result, err :=
			With(incWrappers(nil)).
				OverFiltered(filterEvenFunc(nil),
					With(incWrappers(fmt.Errorf("subpipe error"))),
				).
				Run(context.Background(), vals).Get()

		assert.Error(t, err)
		assert.Equal(t, IntWrappers{{2}, {1}, {2}, {1}, {2}, {1}}, result)
	})

	t.Run("проверка пайп each", func(t *testing.T) {
		t.Parallel()

		vals := IntWrappers{{1}, {2}, {3}, {4}, {5}, {6}}

		f := func(_ context.Context, ptr *IntWrapper) (_ *IntWrapper, err error) {
			ptr.int = (ptr.int * 2)
			return ptr, nil
		}

		result, err :=
			With(Each[IntWrappers](f)).
				Run(context.Background(), vals).Get()

		assert.NoError(t, err)
		assert.Equal(t, IntWrappers{{2}, {4}, {6}, {8}, {10}, {12}}, result)
	})

	t.Run("проверка пайп async each", func(t *testing.T) {
		t.Parallel()

		vals := []int{1, 2, 3, 4, 5, 6}

		f := func(_ context.Context, v int) (_ int, err error) {
			return v * 2, nil
		}

		result, err :=
			With(EachAsync[[]int](f, 0)).
				Run(context.Background(), vals).Get()

		assert.NoError(t, err)
		assert.Equal(t, []int{2, 4, 6, 8, 10, 12}, result)
	})

}

type IntWrapper struct {
	int
}

type IntWrappers []*IntWrapper

func incWrappers(err error) func(_ context.Context, value IntWrappers) (IntWrappers, error) {
	return func(_ context.Context, value IntWrappers) (IntWrappers, error) {
		lo.ForEach(value, func(item *IntWrapper, _ int) {
			item.int++
		})

		return value, err
	}
}

func filterEvenFunc(err error) func(_ context.Context, value IntWrappers) (context.Context, IntWrappers, bool, error) {
	return func(ctx context.Context, value IntWrappers) (context.Context, IntWrappers, bool, error) {
		var res IntWrappers
		for idx, v := range value {
			if idx%2 == 0 {
				res = append(res, v)
			}
		}

		return ctx, res, len(res) == 0, err
	}
}
