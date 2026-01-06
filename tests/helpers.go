package data_parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileExists checks if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadCSV reads a CSV file and returns its contents as a 2D slice of strings
func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteCSV writes a 2D slice of strings to a CSV file
func WriteCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, record := range data {
		if err := csvWriter.Write(record); err != nil {
			return err
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}

// GetFileSize returns the size of a file in bytes
func GetFileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// StringToFloat converts a string to a float64
func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// StringToInt converts a string to an int64
func StringToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// RemoveEmptyStrings removes empty strings from a slice of strings
func RemoveEmptyStrings(s []string) []string {
	var result []string
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// LogError logs an error with a message
func LogError(msg string, err error) {
	log.Printf("%s: %v", msg, err)
}

// CheckErr checks an error and logs it if it's not nil
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}