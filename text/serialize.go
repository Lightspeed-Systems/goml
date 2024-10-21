package text

import (
	"encoding/json"
	"golang.org/x/text/transform"
	"io"
)

// SaveNaiveBayes saves a NaiveBayes model to the filesystem
func SaveNaiveBayes(fh io.Writer, model *NaiveBayes) error {
	enc := json.NewEncoder(fh)
	err := enc.Encode(model)
	if err != nil {
		return err
	}
	return nil
}

// LoadNaiveBayes loads a NaiveBayes model from the filesystem
// The model is already trained - required fields to load are: Words, Count, and Probabilities
// Additionally, the model should define a `sanitize` function and a tokenizer
func LoadNaiveBayes(fh io.Reader, sanitize func(rune) bool) (*NaiveBayes, error) {
	dec := json.NewDecoder(fh)
	model := new(NaiveBayes)
	err := dec.Decode(model)
	if err != nil {
		return nil, err
	}
	// todo: update this to use the new transform package
	model.sanitize = transform.RemoveFunc(sanitize)

	return model, nil
}
