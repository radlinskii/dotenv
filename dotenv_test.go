package dotenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func ExampleSetEnv() {
	fileContent := []byte(`
		TEST1 = 1
		TEST2 = 2
	`)

	filePath := ".env"

	if err := ioutil.WriteFile(filePath, fileContent, 0644); err != nil {
		fmt.Println(err)
	}
	defer os.Remove(filePath)

	SetEnv()

	fmt.Println(os.Getenv("TEST1"))
	fmt.Println(os.Getenv("TEST2"))

	if err := os.Unsetenv("TEST1"); err != nil {
		fmt.Println(err)
	}
	if err := os.Unsetenv("TEST2"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	// 2
}

func ExampleSetEnvFromPath() {
	fileContent := []byte(`TEST1 = 1
# TEST2 = 2
# environment variable #3
TEST3 = 3
`)

	filePath := "testdata/.env"

	if err := ioutil.WriteFile(filePath, fileContent, 0644); err != nil {
		fmt.Println(err)
	}
	defer os.Remove(filePath)

	SetEnvFromPath(filePath)

	fmt.Println(os.Getenv("TEST1"))
	fmt.Println(os.Getenv("TEST2"))
	fmt.Println(os.Getenv("TEST3"))

	if err := os.Unsetenv("TEST1"); err != nil {
		fmt.Println(err)
	}
	if err := os.Unsetenv("TEST2"); err != nil {
		fmt.Println(err)
	}
	if err := os.Unsetenv("TEST3"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// 1
	//
	// 3
}

func TestSetEnvFromPath0(t *testing.T) {
	ok := t.Run("SetEnvFromPath should not return an error if .env file does not exist", func(t *testing.T) {
		path := filepath.Join("testdata", "0.env")

		if err := SetEnvFromPath(path); err != nil {
			t.Error(err)
		}
	})

	if !ok {
		t.Fail()
	}
}

func TestSetEnvFromPath1(t *testing.T) {
	ok := t.Run("SetEnvFromPath should parse empty .env file", func(t *testing.T) {
		path := filepath.Join("testdata", "1.env")

		if err := SetEnvFromPath(path); err != nil {
			t.Error(err)
		}
	})

	if !ok {
		t.Fail()
	}
}

func TestSetEnvFromPath2(t *testing.T) {
	ok := t.Run("SetEnvFromPath should parse variables from .env file", func(t *testing.T) {
		path := filepath.Join("testdata", "2.env")

		if err := SetEnvFromPath(path); err != nil {
			t.Error(err)
		}

		testVariableValue(t, "1", "TEST1")
		testVariableValue(t, "2", "TEST2")
	})

	if !ok {
		t.Fail()
	}

	unsetVariables(t, "TEST1", "TEST2")
}

func TestSetEnvFromPath3(t *testing.T) {
	ok := t.Run("SetEnvFromPath should omit commented lines", func(t *testing.T) {
		path := filepath.Join("testdata", "3.env")

		if err := SetEnvFromPath(path); err != nil {
			t.Error(err)
		}

		testVariableValue(t, "1", "TEST1")
		testVariableValue(t, "", "TEST2")
		testVariableValue(t, "", "TEST3")
	})

	if !ok {
		t.Fail()
	}

	unsetVariables(t, "TEST1", "TEST2", "TEST3")
}

func TestSetEnvFromPath4(t *testing.T) {
	ok := t.Run("SetEnvFromPath should not overwrite exported variables", func(t *testing.T) {
		// exporting TEST1 variable before reading the .env file
		if err := os.Setenv("TEST1", "99"); err != nil {
			t.Error(err)
		}

		path := filepath.Join("testdata", "4.env")
		if err := SetEnvFromPath(path); err != nil {
			t.Error(err)
		}

		testVariableValue(t, "99", "TEST1")
	})

	if !ok {
		t.Fail()
	}

	unsetVariables(t, "TEST1")
}

func TestSetEnvFromPath5(t *testing.T) {
	ok := t.Run("SetEnvFromPath should return parsing error on unformatted file", func(t *testing.T) {
		path := filepath.Join("testdata", "5.env")

		if err := SetEnvFromPath(path); err != nil {
			if err.Error() != "Error parsing "+path+" file" {
				t.Error(err)
			}
		}
	})

	if !ok {
		t.Fail()
	}
}

func TestSetEnvFromPath6(t *testing.T) {
	ok := t.Run("SetEnvFromPath should return os.Setenv error", func(t *testing.T) {
		path := filepath.Join("testdata", "6.env")

		if err := SetEnvFromPath(path); err != nil {
			if expected, got := os.NewSyscallError("setenv", syscall.EINVAL).Error(), err.Error(); got != expected {
				t.Errorf("Wrong error! expected: %q, got: %q", expected, got)
			}
		}
	})

	if !ok {
		t.Fail()
	}
}

func TestSetEnvFromPath7(t *testing.T) {
	path := filepath.Join("testdata", "7.env")
	fileInfo, err := os.Stat(path)
	if err != nil {
		t.Error(err)
	}
	// save the current mode of the file
	originalMode := fileInfo.Mode()

	ok := t.Run("SetEnvFromPath should return ioutil.Readfile error", func(t *testing.T) {
		// remove permission to read the specified file so that permission error can be returned
		if err := os.Chmod(path, 0333); err != nil {
			t.Fatal(err)
		}

		if err := SetEnvFromPath(path); err != nil {
			pathErr := &os.PathError{Op: "open", Path: path, Err: os.ErrPermission}
			if expected, got := pathErr.Error(), err.(*os.PathError).Error(); expected != got {
				t.Errorf("Wrong error! expected: %q, got: %q", expected, got)
			}
		}
	})

	if !ok {
		t.Fail()
	}

	// change readed file permission back to default
	if err := os.Chmod(path, originalMode); err != nil {
		t.Fatal(err)
	}
}

func testVariableValue(t *testing.T, expected, name string) {
	if got, ok := os.LookupEnv(name); ok {
		if got != expected {
			t.Errorf("Wrong env variable %q value! expected: %q, got: %q", name, expected, got)
		}
	}
}

func unsetVariables(t *testing.T, variables ...string) {
	for _, v := range variables {
		if err := os.Unsetenv(v); err != nil {
			t.Error(err)
		}
	}
}
