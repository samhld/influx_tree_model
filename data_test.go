package main

import (
	"testing"
)

type StubAPI struct {
	measurement string
	tags        []Tag
	tagKeys     []string
	fieldKeys   []string
	keyValMap   map[string][]string
	fields      []*Field
}

// func (s *StubAPI) Query(ctx context.Context, query string) (*QueryTableResult, error) {
// 	return nil, nil
// }

// func (s *StubAPI) QueryRaw(ctx context.Context, query string, dialect *domain.Dialect) (string, error) {
// 	return nil, nil
// }

// implement dataGetter
//________________________
func (s *StubAPI) getTagKeys() []string {
	return s.tagKeys
}

func (s *StubAPI) getTagKeyValues(key string) []string {
	return s.keyValMap[key]
}

func (s *StubAPI) getTagKeyValueCounts() map[string]int64 {
	keyValCountMap := make(map[string]int64)
	for key, vals := range s.keyValMap {
		keyValCountMap[key] = int64(len(vals))
	}
	return keyValCountMap
}

func (s *StubAPI) getFieldKeys() []string {
	return s.fieldKeys
}

func (s *StubAPI) getMeasurement() string {
	return s.measurement
}

func (s *StubAPI) getFields() []*Field {
	return s.fields
}

//______________________
func TestMakeTiers(t *testing.T) {
	rule := "MEASUREMENT>region>app>FIELD"
	// use raw StubAPI to test creation of the struct
	stub := &StubAPI{
		"test",
		[]Tag{
			{Key{"region", 1, nil, nil}, Value{"us-west", 2, nil, nil}},
			{Key{"region", 1, nil, nil}, Value{"us-east", 2, nil, nil}},
			{Key{"app", 3, nil, nil}, Value{"cart", 4, nil, nil}},
			{Key{"app", 3, nil, nil}, Value{"home", 4, nil, nil}},
			{Key{"app", 3, nil, nil}, Value{"login", 4, nil, nil}},
		},
		[]string{"region", "app"},
		[]string{"value"},
		map[string][]string{
			"region": {"us-west", "us-east"},
			"app":    {"cart", "home", "login"},
		},
		[]*Field{{"value", 5}},
	}
	tokenizedRule := NewRuleTokenizer().Tokenize(rule)

	gotTiers := MapTokensToData(stub, tokenizedRule)
	wantTiers := Tiers{
		0: &Measurement{"test", 0},
		1: &Key{"region", 1, nil, nil},
		2: []*Value{
			{"us-west", 2, nil, nil},
			{"us-east", 2, nil, nil},
		},
		3: &Key{"app", 3, nil, nil},
		4: []*Value{
			{"cart", 4, nil, nil},
			{"home", 4, nil, nil},
			{"login", 4, nil, nil},
		},
		5: []*Field{
			{"value", 5},
		},
	}

	assertEqual(t, gotTiers, wantTiers)
}
