This is a fork of part of the golang.org/x/net/html package.

## v0.2.0

For v0.2.0 we made a more radical change to the [tokenizer](https://pkg.go.dev/golang.org/x/net/html#Tokenizer) package.

We added a new syntax to allow attributes to be set with '{}' syntax.
Any valid JSON expression is allowed within the curly brackets (this more
closely matches JSX syntax).

```
<div data-num={5}></div>
```

To support proper decoding in the client, attributes now have a an `IsJson bool` field
which is set to true if an attribute was parsed with the new {} syntax.

If you only need the case-sensitive tokenization for tags/attributes it is
recommended to use v0.1.0 and not v0.2.0.

## v0.1.0

It is not a complete fork as we only want to modify and change https://pkg.go.dev/golang.org/x/net/html#Tokenizer.  So this is the minimal amount of code to get html.Tokenizer working.

The reason for the fork is to allow for returning of case-sensitive tag names and attribute names.  The current package normalizes the tag names and attribute names by calling (the equivalent of) strings.ToLower on them before returning them to the caller.  We made a very small two line change in token.go to remove those ToLower calls.  Other changes involve copying enough code from other files to get all the dependencies satisfied and get it compling again.

Why did we not fork the entire package?  Because the rest of the html package is a _validating_ html parser and is quite complicated.  As the HTML rules can change over time, it would need to be continually updated and synced with the upstream to keep it compliant.  As the actual syntax (tokenization rules) of HTML does not change often, this part of the package is likely much more stable.