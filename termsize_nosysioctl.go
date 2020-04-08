// +build plan9 solaris appengine wasm aix

package flags

func getTerminalColumns() int {
	return 80
}
