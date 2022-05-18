package string_sum

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	res, err := StringSum("   ")
	if res != "" {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", "<empty string>", res)
	}
	if err != nil && err.Error() != errorEmptyInput.Error() {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", errorEmptyInput, err)
	}

	res, err = StringSum("1")
	if res != "" {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", "<empty string>", res)
	}
	if err != nil && err.Error() != errorNotTwoOperands.Error() {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", errorNotTwoOperands, err)
	}

	res, err = StringSum("1+2+3")
	if res != "" {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", "<empty string>", res)
	}
	if err != nil && err.Error() != errorNotTwoOperands.Error() {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", errorNotTwoOperands, err)
	}

	res, err = StringSum("3+5")
	if res != "8" || err != nil {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", "8", res)
	}

	res, err = StringSum("-3+5")
	if res != "2" || err != nil {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", "2", res)
	}

	res, err = StringSum("-3-5")
	if res != "-8" || err != nil {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", "-8", res)
	}

	res, err = StringSum("3-5")
	if res != "-2" || err != nil {
		t.Errorf("Unexpected error:\n\tExpected: %q\n\tGot: %q", "-2", res)
	}

}
