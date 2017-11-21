package main

import (
	"testing"
)

func TestShouldReturnString(t *testing.T) {
	for name, testIO := range map[string]struct {
		in     int
		expect string
	}{
		"0": {
			in: 0,
			expect: `
 _
| |
|_|
`,
		},
		"1": {
			in: 1,
			expect: `

  |
  |
`,
		},
		"2": {
			in: 2,
			expect: `
 _
 _|
|_
`,
		},
		"3": {
			in: 3,
			expect: `
 _
 _|
 _|
`,
		},
		"4": {
			in: 4,
			expect: `

|_|
  |
`,
		},
		"5": {
			in: 5,
			expect: `
 _
|_
 _|
`,
		},
		"6": {
			in: 6,
			expect: `
 _
|_
|_|
`,
		},
		"7": {
			in: 7,
			expect: `
 _
  |
  |
`,
		},
		"8": {
			in: 8,
			expect: `
 _
|_|
|_|
`,
		},
		"9": {
			in: 9,
			expect: `
 _
|_|
 _|
`,
		},
	} {
		result := ToLCD(testIO.in)
		if result != testIO.expect {
			t.Errorf("%s expected %s but got %s", name, testIO.expect, result)
		}
	}
}

func TestShouldReturn7DigitCode (T *testing.T)
{

}