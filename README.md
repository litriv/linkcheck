# linkcheck

Checks the image links in a (more or less) 300 page manual in docx format and report on the results.

## Running the program

The program optionally takes an argument - the name of the file to check.  If the program is invoked without an argument, it reads from stdin, instead of the file.

Examples of ways of running the program:

* Without arguments

		linkcheck <filename
or
		cat filename | linkcheck

* With arguments

linkcheck filename

## Output and errors

Output goes to stdout, with fatal errors to stderr.  The program returns with exit status 1 on a fatal error.