# Learning the GO language 

## WC GO - a wordcount implementation in GO

I'm aware that there are many implementations of WC in GO and in other languages as well. I wrote this to learn more about the GO programming language. Initially I had some problems with getting the wrong amount of characters returned from the character function. I'm not sure if it was a testing issue rather than an issue with the function. I wrote unit tests and they all passed, I ran the tests with *The Odysse* and *The Complete Works of William Shakespeare* and it corresponded to the numbers from the original WC.

It was a fun project and I learned a lot.


### Useage 
Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified. A word is a non-zero-length sequence of printable characters delimited by white space.

The options below may be used to select which counts are printed, always in the following order: newline, word, character, byte, maximum line length.


``` bash
-l Print the newline counts

-c Print the byte counts

-m Print the character counts

-w Print the word counts

-L Print the maximum display width

```
