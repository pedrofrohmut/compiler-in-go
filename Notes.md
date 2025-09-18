# Notes

## First TODO: Make the executable work for debug

  For the executable to make it usefull for both the debug and REPL without the flag
in the program, I intend to add the flag --input or -i where you provide the string
to be used and than it will trigger the debugger and when there is no --input flag
then just open the REPL.

- Example:

```console
# Opens the monkey repl for parsing
$ monkey parser

# Runs the parser with the input and does not open the repl
$ monkey parser --input 'let foo ="bar";'
```
