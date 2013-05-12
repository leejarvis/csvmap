package csvmap

import (
	"encoding/csv"
	"io"
	"os"
)

// CSVRecord is a custom map type for storing header=>value strings
// from CSV rows.
type CSVRecord map[string]string

// CSVRecords implements a container for CSVRecord types.
type CSVRecords []CSVRecord

// Headers returns a slice of headers associated with this record list.
func (r CSVRecords) Headers() []string {
	keys := make([]string, 0, len(r[0]))
	for k := range r[0] {
		keys = append(keys, k)
	}
	return keys
}

// Implement a custom Reader for handling old-school Windows CR line
// breaks. Thanks Dom!
type CSVReader struct {
	r io.Reader
}

func (or CSVReader) Read(dst []byte) (n int, err error) {
	n, err = or.r.Read(dst)
	for i := range dst[:n] {
		if dst[i] == '\r' {
			dst[i] = '\n'
		}
	}
	return
}

// CSVToRecords opens and parses the CSV file found at filepath.
// It constructs and returns a slice of CSVRecord maps with
// headers mapped to values.
func CSVToRecords(filepath string) (CSVRecords, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csv := csv.NewReader(CSVReader{file})
	csv.TrailingComma = true
	csv.TrimLeadingSpace = true

	rows, err := csv.ReadAll()
	if err != nil {
		return nil, err
	}
	headers, rows := rows[0], rows[1:]
	records := make([]CSVRecord, len(rows))
	for num, row := range rows {
		rec := make(CSVRecord, len(row))
		for i, header := range headers {
			rec[header] = row[i]
		}
		records[num] = rec
	}

	return records, nil
}
