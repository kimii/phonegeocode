package phonegeocode

import (
	"testing"
)

func TestThingsThatAreFound(t *testing.T) {
	cases := []struct {
		number, country string
	}{
		{"+4479991113332", "GB"},
		{"+1807142342342", "CA"},
		{"+1781234555552", "US"},
		{"+3462233455552", "ES"},
		{"+3538523523455", "IE"},
		{"+19425242343333", "US"},
		{"+1786822211132", "US"},
		{"+34634343434443", "ES"},
		{"+190572354235235", "CA"},
	}

	p := New()

	for _, tc := range cases {
		cc, err := p.Country(tc.number)
		if err != nil {
			t.Errorf("Not expecting number '%s' to yield an error; got %v", tc.number, err)
		}
		if cc != tc.country {
			t.Errorf("Number '%s' did not match expected CC '%s'; got '%s'", tc.number, tc.country, cc)
		}
	}
}
