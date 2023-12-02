package main

import "testing"

func TestProcessInput(t *testing.T) {
	tests := []struct{
		input string
		want uint64
	}{
		{
			input: "1abc2",
			want: 12,
		},

		{
			input: "pqr3stu8vwx",
			want: 38,
		},

		{
			input: "a1b2c3d4e5f",
			want: 15,
		},

		{
			input: "treb7uchet",
			want: 77,
		},
		{
			input: "7b",
			want: 77,
		},
		{
			input: "test",
			want: 0,
		},
	}

	for _, test := range tests {
		got, err := proccessInput([]byte(test.input)); 
		if err != nil {
			t.Fatalf("processLine(%q) returned error", err)		
		}

		if got != test.want {
			t.Errorf("processInput(%q) = %d want %d", test.input, got, test.want)
		}
	}
}
