package types

type Command struct {
	Cmd		string // The literal string to invoke this command. "ping", "help", "todo", etc
	Version	string // Arbitrary version. Mostly for dev/debug purposes

	Func	interface{}
	Flags	FuncFlags
	Usage	string // The usage info to be printed for !help
}

// Usage should have format of
// `
// <cmdname> <args>
//		- Information on usage
// <cmdname> <otherarg> <moreargs>
//		- Some information
//		  provided on multiple
//
//		  lines like so.
// <cmdname> <others>
//		- asdf

// The help function will automatically prepend the function prefix to each <cmdname> but the help must be written in the above format for this to work.
// What actually matters is that <cmdname> is on the first col of the line, ie match /^<cmdname>/
