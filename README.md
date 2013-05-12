CSVMap
======

Map CSV records into easy-to-use data structures.

Install
-------

```
go get github.com/injekt/csvmap
```

Usage
-----

```go
func main() {
  records, err := csvmap.CSVToRecords("path/to/file.csv")
  if err != nil {
    return
  }
  for _, r := range records {
    // r is a csvmap.CSVRecord type (custom map[string]string)
    for header, value := range record {
      fmt.Printf("%s => %s\n", header, value)
    }
  }
}
```