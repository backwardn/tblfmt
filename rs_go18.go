// +build go1.8

package tblfmt

import (
	"io"
)

// ResultSet is the shared interface for a result set.
type ResultSet interface {
	Next() bool
	Scan(...interface{}) error
	Columns() ([]string, error)
	Close() error
	Err() error
	NextResultSet() bool
}

// EncodeAll encodes all result sets in rs to the writer using the encoder
// settings.
func (enc *TableEncoder) EncodeAll(w io.Writer) error {
	var err error
	if err = enc.Encode(w); err != nil {
		return err
	}
	for enc.resultSet.NextResultSet() {
		if err = enc.Encode(w); err != nil {
			return err
		}
	}
	return nil
}
