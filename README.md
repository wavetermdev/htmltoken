This is a fork of part of the golang.org/x/net/html package.

It is not a complete fork as we only want to modify and change https://pkg.go.dev/golang.org/x/net/html#Tokenizer.  So this is the minimal amount of code to get html.Tokenizer working.

The reason for the fork is to allow for returning of case-sensitive tag names and attribute names.  The current package normalizes the tag names and attribute names by calling (the equivalent of) strings.ToLower on them before returning them to the caller.  We made a very small two line change in token.go to remove those ToLower calls.  Other changes involve copying enough code from other files to get all the dependencies satisfied and get it compling again.

Why did we not fork the entire package?  Because the rest of the html package is a _validating_ html parser and is quite complicated.  As the HTML rules can change over time, it would need to be continually updated and synced with the upstream to keep it compliant.  As the actual syntax (tokenization rules) of HTML does not change often, this part of the package is likely much more stable.