# TheFool
TheFool is a simple command line csv combination program that can be used to merge, subtract and subset csv files. 
## Instalation 
building the fool only requires a working go tool chain and can be done by running
```
git clone https://github.com/Grimmr/TheFool.git
cd TheFool
go build .
```
## Usage
TheFool takes a single expression to evaluate and prints the esulting csv table to stdout. 
### Example usage
```
TheFool dogs.csv less myDogs.csv > notMyDogs.csv
TheFool dogs.csv or cats.csv
TheFool animals.csv less (birds.csv or dogs.csv)
```
### Operators
#### and
selects only the rows where all fields common between both tables are the same. The resultant table will have only the common fields.
example:
```
|h1, h2|     |h1, h3|   |h1|
|a,  b | and |a , e | = |a |
|c,  d |     |f,  g |
```

### or
combines both inputs into a single table by concatanating one onto the other. If the two input tables have a diferent set of field names the resultant table will have all fields from both inputs, any new fields added to a row will be left blank. 
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
