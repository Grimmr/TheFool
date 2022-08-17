# TheFool
TheFool is a simple command line csv combination program that can be used to merge, subtract and subset csv files. 
## Installation
### From Release
Pre-built binaries for Windows and linux can be found on the releases page 
### From Source
building TheFool only requires a working go tool chain and can be done by running
```
git clone https://github.com/Grimmr/TheFool.git
cd TheFool
go build .
```
## Usage
TheFool takes a single expression to evaluate and prints the resulting csv table to stdout. 
### Example usage
```
TheFool dogs.csv less myDogs.csv > notMyDogs.csv
TheFool dogs.csv or cats.csv
TheFool animals.csv less (birds.csv or dogs.csv)
```
### Operators
#### and
selects only the rows where all fields common between both tables are the same. The resultant table will have only the common fields. Output will be made set like (ie: no duplicate rows).
example:
```
|h1, h2|     |h1, h3|   |h1|
|a,  b | and |a , e | = |a |
|c,  d |     |f,  g |
```

### or
combines both inputs into a single table by concatenating one onto the other. If the two input tables have a different set of field names the resultant table will have all fields from both inputs, any new fields added to a row will be left blank. Output will be made set like (ie: no duplicate rows).
example:
```
|h1, h2|     |h1, h3|   |h1, h2, h3|
|a,  b | or  |a , e | = |a,  b,    |
                        |a,  ,   e |
```

### less
removes entries in the right hand input from the left hand input. Only considers common fields.
example:
```
|h1, h2|      |h1|   |h1, h2|
|a,  b | less |a | = |c,  d |
|c,  d |     
```

### % \<number\>
creates a random subset of the input table with length \<number\>. all fields are preserved 

### +
like 'or', combines both inputs into a single table by concatenating one onto the other. If the two input tables have a different set of field names the resultant table will have all fields from both inputs, any new fields added to a row will be left blank. unlike or this does not remove duplicate rows from the output
```
|h1, h2|     |h2, h2|   |h1, h2|
|a,  b | or  |a , b | = |a,  b |
                        |a,  b |
```