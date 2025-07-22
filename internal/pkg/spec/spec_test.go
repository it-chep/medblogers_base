package spec

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

var (
	errStub1 = errors.New("1 - false")
	errStub2 = errors.New("2 - false")
	errStub3 = errors.New("3 - false")
)

func NewSpecificationStub(err error) Specification[any] {
	return NewSpecification[any](func(ctx context.Context, t any) (bool, error) {
		return err == nil, err
	})
}

func NewEmptySpecificationStub() Specification[any] {
	return NewSpecification[any](func(ctx context.Context, t any) (bool, error) {
		return false, nil
	})
}

type expected struct {
	ok      bool
	err     error
	twoErrs bool
}

type testCase struct {
	name     string
	spec     Specification[any]
	expected expected
}

func TestSpecification(t *testing.T) {
	t.Parallel()

	cases := []testCase{
		// common
		{
			name:     "обобщенная ошибка если не определена ошибка для правила",
			spec:     NewEmptySpecificationStub(),
			expected: expected{ok: false, err: fmt.Errorf("отсутсвует ошибка для правила валидации")},
		},
		// ? AND ?
		{
			name:     "Объединение правил через AND: true AND true",
			spec:     NewSpecificationStub(nil).And(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через AND: true AND false",
			spec:     NewSpecificationStub(nil).And(NewSpecificationStub(errStub2)),
			expected: expected{ok: false, err: errStub2},
		},
		{
			name:     "Объединение правил через AND: false AND true",
			spec:     NewSpecificationStub(errStub1).And(NewSpecificationStub(nil)),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name:     "Объединение правил через AND: false AND false",
			spec:     NewSpecificationStub(errStub1).And(NewSpecificationStub(errStub2)),
			expected: expected{ok: false, err: errStub1},
		},
		// ? OR ?
		{
			name:     "Объединение правил через OR: true OR true",
			spec:     NewSpecificationStub(nil).Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR: true OR false",
			spec:     NewSpecificationStub(nil).Or(NewSpecificationStub(errStub2)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR: false OR true",
			spec:     NewSpecificationStub(errStub1).Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR: false OR false",
			spec:     NewSpecificationStub(errStub1).Or(NewSpecificationStub(errStub2)),
			expected: expected{ok: false, err: fmt.Errorf("%s %s %s", errStub1, specConnector, errStub2), twoErrs: true},
		},
		// ? AND ? OR ?
		{
			name: "Объединение правил через AND OR: true AND true OR true",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(nil)).
				Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: false AND true OR true",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(nil)).
				Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: true AND false OR true",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(errStub2)).
				Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: true AND true OR false",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(nil)).
				Or(NewSpecificationStub(errStub3)),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: false AND false OR true",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(errStub2)).
				Or(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: true AND false OR false",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(errStub2)).
				Or(NewSpecificationStub(errStub3)),
			expected: expected{ok: false, err: fmt.Errorf("%s %s %s", errStub2, specConnector, errStub3), twoErrs: true},
		},
		{
			name: "Объединение правил через AND OR: false AND true OR false",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(nil)).
				Or(NewSpecificationStub(errStub3)),
			expected: expected{ok: false, err: fmt.Errorf("%s %s %s", errStub1, specConnector, errStub3), twoErrs: true},
		},
		{
			name: "Объединение правил через AND OR: false AND false OR false",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(errStub2)).
				Or(NewSpecificationStub(errStub3)),
			expected: expected{ok: false, err: fmt.Errorf("1 - false или 3 - false"), twoErrs: true},
		},
		// ? AND (? OR ?)
		{
			name: "Объединение правил через AND OR: true AND (true OR true)",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(nil).Or(NewSpecificationStub(nil))),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: false AND (true OR true)",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(nil).Or(NewSpecificationStub(nil))),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name: "Объединение правил через AND OR: true AND (false OR true)",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(errStub2).Or(NewSpecificationStub(nil))),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: true AND (true OR false)",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(nil).Or(NewSpecificationStub(errStub3))),
			expected: expected{ok: true, err: nil},
		},
		{
			name: "Объединение правил через AND OR: false AND (false OR true)",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(errStub2).Or(NewSpecificationStub(nil))),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name: "Объединение правил через AND OR: true AND (false OR false)",
			spec: NewSpecificationStub(nil).
				And(NewSpecificationStub(errStub2).Or(NewSpecificationStub(errStub3))),
			expected: expected{ok: false, err: fmt.Errorf("%s %s %s", errStub2, specConnector, errStub3), twoErrs: true},
		},
		{
			name: "Объединение правил через AND OR: false AND (true OR false)",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(nil).Or(NewSpecificationStub(errStub3))),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name: "Объединение правил через AND OR: false AND (false OR false)",
			spec: NewSpecificationStub(errStub1).
				And(NewSpecificationStub(errStub2).Or(NewSpecificationStub(errStub3))),
			expected: expected{ok: false, err: errStub1},
		},
		// ? AND NOT ?
		{
			name:     "Объединение правил через AND NOT: false AND NOT false",
			spec:     NewSpecificationStub(errStub1).AndNot(NewSpecificationStub(errStub2)),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name:     "Объединение правил через AND NOT: false AND NOT true",
			spec:     NewSpecificationStub(errStub1).AndNot(NewSpecificationStub(nil)),
			expected: expected{ok: false, err: errStub1},
		},
		{
			name:     "Объединение правил через AND NOT: true AND NOT false",
			spec:     NewSpecificationStub(nil).AndNot(NewSpecificationStub(errStub1)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через AND NOT: true AND NOT true",
			spec:     NewSpecificationStub(nil).AndNot(NewSpecificationStub(nil)),
			expected: expected{ok: false, err: fmt.Errorf("правило валидации, зарегистрированное с логическим отрицанием отработало без ошибки")},
		},
		// ? OR NOT ?
		{
			name:     "Объединение правил через OR NOT: false OR NOT false",
			spec:     NewSpecificationStub(errStub1).OrNot(NewSpecificationStub(errStub2)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR NOT: true OR NOT false",
			spec:     NewSpecificationStub(nil).OrNot(NewSpecificationStub(errStub1)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR NOT: true OR NOT true",
			spec:     NewSpecificationStub(nil).OrNot(NewSpecificationStub(nil)),
			expected: expected{ok: true, err: nil},
		},
		{
			name:     "Объединение правил через OR NOT: true OR NOT true",
			spec:     Not(NewSpecificationStub(nil)).And(NewSpecificationStub(nil)),
			expected: expected{ok: false, err: fmt.Errorf("правило валидации, зарегистрированное с логическим отрицанием отработало без ошибки")},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ok, err := tc.spec.IsSatisfied(context.Background(), nil)

			assert.Equal(t, tc.expected.ok, ok)
			if tc.expected.twoErrs {
				assert.True(t,
					(errors.Is(err, errStub1) && errors.Is(err, errStub2)) ||
						(errors.Is(err, errStub1) && errors.Is(err, errStub3)) ||
						(errors.Is(err, errStub2) && errors.Is(err, errStub3)),
				)
				assert.EqualError(t, err, tc.expected.err.Error())
			} else {
				assert.EqualValues(t, tc.expected.err, err)
			}
		})
	}
}
