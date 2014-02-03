/*
	envreader provides a small wrapper for os.GetEnv for retrieving a collection of environment variables into a map.

	Simple example:

		myEnv := envreader.Read("PATH")
		usePath(myEnv["PATH"])
*/
package envreader

import (
	"os"
)

/*
	Reads and returns all environment variables found in keys
*/
func Read (keys ...string) map[string]string {
	vals := make(map[string]string)
	for _, key := range keys {
		vals[key] = os.Getenv(key)
	}
	return vals
}
