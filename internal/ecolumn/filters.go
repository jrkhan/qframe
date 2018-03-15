package ecolumn

import (
	"github.com/tobgu/qframe/errors"
	"github.com/tobgu/qframe/filter"
	"github.com/tobgu/qframe/internal/index"
	qfstrings "github.com/tobgu/qframe/internal/strings"
)

var filterFuncs = map[string]func(index.Int, []enumVal, enumVal, index.Bool){
	filter.Gt:  gt,
	filter.Gte: gte,
	filter.Lt:  lt,
	filter.Lte: lte,
	filter.Eq:  eq,
	filter.Neq: neq,
}

var filterFuncs2 = map[string]func(index.Int, []enumVal, []enumVal, index.Bool){
	filter.Gt:  gt2,
	filter.Gte: gte2,
	filter.Lt:  lt2,
	filter.Lte: lte2,
	filter.Eq:  eq2,
	filter.Neq: neq2,
}

var multiFilterFuncs = map[string]func(comparatee string, values []string) (*bitset, error){
	"like":  like,
	"ilike": ilike,
}

var multiInputFilterFuncs = map[string]func(comparatee qfstrings.StringSet, values []string) *bitset{
	"in": in,
}

func like(comp string, values []string) (*bitset, error) {
	return filterLike(comp, values, true)
}

func ilike(comp string, values []string) (*bitset, error) {
	return filterLike(comp, values, false)
}

func filterLike(comp string, values []string, caseSensitive bool) (*bitset, error) {
	matcher, err := qfstrings.NewMatcher(comp, caseSensitive)
	if err != nil {
		return nil, errors.Propagate("enum like", err)
	}

	bset := &bitset{}
	for i, v := range values {
		if matcher.Matches(v) {
			bset.set(enumVal(i))
		}
	}

	return bset, nil
}

func in(comp qfstrings.StringSet, values []string) *bitset {
	bset := &bitset{}
	for i, v := range values {
		if comp.Contains(v) {
			bset.set(enumVal(i))
		}
	}

	return bset
}