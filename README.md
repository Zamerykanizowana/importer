# importer
Funny importer. Tool to count customers from one domain.

## Expected input
`.csv` file with `email` column.

## Expected output
`.csv` file with columns: `domain` and `number of customers`.


## How to build
```
go build
```

## How to run 

### Without build 
```
go run importer.go -path <path_to_input_file>
```

Default output file is `outfile.csv`. For different output file name use:
```
go run importer.go -path <path_to_input_file> -output <path_to_output_file>
```

### With build
```
./importer -path <path_to_inport_file>
```
or
```
./importer -path <path_to_input_file> -output <path_to_output_file>
```


## Test

### Run tests
```
cd customerimporter/
go test
```
