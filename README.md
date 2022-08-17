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
 | Op | Usage | Header Semantics | Row Semantics | Row Multiplicity |
 |----|-------|------------------|---------------|------------------|
 | and | \<table\> and \<table\> | Only headers that exist in both input tables are selected | Only common fields are considered. Only rows that exist in both input tables are selected. | either 1 or 0, duplicate rows are discarded |
 | or | \<table\> or \<table\> | All headers from both inputs tables are selected | All rows from both tables are selected. Any fields added to a row are left empty. | eitehr 1 or 0, duplicate rows are discarded |
 | less | \<table\> less \<table\> | Only headers from the left hand input table are selected | Only common fields are considered. All rows that apear in the left hand input table less the rows that apear in the right hand table are slected. | Selected rows have the same multiplicity as in the left hand input table. |   
 | % | \<table\> % \<number\> | All headers from the input table are selected | Rows are selected randomly up to either \<number\> or the length of the input table whichever is smaller. | Multiplicity of a row is no more then it's multiplicity in the input table. |
 | + | \<table\> + \<number\> | All headers from both inputs tables are selected | All rows from both tables are selected. Any fields added to a row are left empty. | The multiplicity of a selected row is exactly the sum of its multiplities in the input tables |
