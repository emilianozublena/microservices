package internal

import (
	"fmt"
	"os"
	"testing"
)

// @todo write tests for internal package
func TestGetEnv(t *testing.T) {
	envKey := "MYENV"
	envValue := "VALUE"
	envDefault := "DEFAULT"
	t.Run("Fallback to default OS Env", func(t *testing.T) {
		//Given we have a non existing OS env
		//When we try to get it
		myEnv := GetEnv(envKey, envDefault)
		//Then we assert we got it
		if myEnv != envDefault {
			t.Error("Expected default", envDefault, "but got", myEnv)
		}
	})

	t.Run("Existing OS env", func(t *testing.T) {
		//Given we have a valid OS env
		os.Setenv(envKey, envValue)
		//When we try to get it
		myEnv := GetEnv(envKey, envDefault)
		//Then we assert we got it
		if myEnv != envValue {
			t.Error("Expected", envValue, "Got", myEnv)
		}
	})
}

func BenchmarkGetEnv(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetEnv("SomeKey", "SomeDefault")
	}
}

func ExampleGetEnv() {
	myEnv := GetEnv("SomeKey", "DefaultValue")
	fmt.Println(myEnv)
	// Output: DefaultValue
}
