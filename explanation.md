# Explanation

The proposed design conforms to the principle of keeping things as simple as possible for solving the problem at hand, without premature abstraction or optimisation.

## Command line argument parsing

The program starts execution in the main method by checking arguments and setting the appropriate input source - an opened file for reading in the case where there is an argument, stdin when there is no argument and exits with an error if the wrong number of arguments is provided.

## Link collection

Link collection then starts, using the input source determined from the command line argument's presence or absence.

The link collection routine is separated into a function (named collect) so that it can be tested (main_test.go).

The collect function gets passed a reader - that way it is generic and anything that implements io.Read can be collected from.  
The collect function returns a list of links.

## Link checking

Now that the list of links had been collected from the input, link checking can proceed.
Link checking needs to happen concurrently to speed up the app.  It is the only part that need to be concurrent, as this will be the only significant performance bottleneck.

Once we have the links from calling the collect method, for each link we fire up a goroutine (the check function), to check the link.

The check function does an http.Get request and construct a result - a struct type, which it sends on a result channel.


## Results reporting

The results then get consumed from the results channel, generating the report.

A problem to overcome is how to know when to stop polling for results.  The simplest way is by only doing as many polls as there are links.  As an alternative, one could use sync.WaitGroup, incrementing (wg.Add) when a new check starts and decrementing (wg.Done) at the end of each goroutine that does the checking, firing off the polling loop in a separate goroutine and waiting (in the main routine) until they are all finished (wg.Wait).

The main routine then reads from the result channel and creates a report from the data.  

The results for the links won't arrive in the same order as the link appear in the document, which may or may not be an issue, depending on requirements.  If it is an issue, then it can be remedied by mapping the results to their links, so that the original links list can be iterated through and the appropriate result be found by looking it up in the map (keyed off the link).

The list can be iterated through twice, once for reporting the working links and once for reporting the error links afterwards.

 





