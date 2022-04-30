package strings

/**
You’re working with an intern that keeps coming to you with JavaScript code that won’t run because the braces, brackets,
and parentheses are off. To save you both some time, you decide to write a braces/brackets/parentheses validator.
Let’s say:

‘(‘, ‘{‘, ‘[‘ are called “openers.”
‘)’, ‘}’, ‘]’ are called “closers.”
Write an efficient function that tells us whether an input string’s openers and closers are properly nested.

Examples:

“{ [] ( ) }” should return 1
“{ [(] ) }” should return 0
“{ [ }” should return 0


Breakdown

We can use a greedy approach to walk through our string character by character,
making sure the string validates “so far” until we reach the end.

What do we do when we find an opener or closer?

Well, we’ll need to keep track of our openers so that we can confirm they get closed properly.
What data structure should we use to store them? When choosing a data structure, we should start by
deciding on the properties we want. In this case, we should figure out how we will want to retrieve
our openers from the data structure! So next we need to know: what will we do when we find a closer?

Suppose we’re in the middle of walking through our string, and we find our first closer:

[{ ( )] . . . . ^

How do we know whether or not that closer in that position is valid?

A closer is valid if and only if it’s the closer for the most recently seen, unclosed opener.
In this case, ‘(‘ was seen most recently, so we know our closing ‘)’ is valid.

So we want to store our openers in such a way that we can get the most recently added one quickly ,
and we can remove the most recently added one quickly (when it gets closed). Does this sound familiar?

What we need is a stack!

Solution

We iterate through our string, making sure that:

each closer corresponds to the most recently seen, unclosed opener
every opener and closer is in a pair
We use a stack to keep track of the most recently seen, unclosed opener. And if the stack is ever empty when we come to a closer,
we know that closer doesn’t have an opener.

So as we iterate:

If we see an opener, we push it onto the stack.
If we see a closer, we check to see if it is the closer for the opener at the top of the stack.
If it is, we pop from the stack. If it isn’t, or if the stack is empty, we return 0.
If we finish iterating and our stack is empty, we know every opener was properly closed.

*/

func Isvalid(s string) bool {
	open := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var cmp []rune
	// if length of string is not even do not even try to validate
	if len(s)%2 != 0 {
		return false
	}
	for _, v := range s {
		//if the current character is in the open map,  put its closer into the stack and continue
		if closer, ok := open[v]; ok {
			cmp = append(cmp, closer)
			continue
		}
		// otherwise, we're dealing with a closer
		// check to make sure the stack isn't empty
		// and whether the top of the stack matches
		// the current character
		l := len(cmp) - 1
		if l < 0 || v != cmp[l] {
			return false
		}
		// take the last element off the stack
		cmp = cmp[:l]
	}
	// if the stack is empty, return true, otherwise false
	return len(cmp) == 0

}
