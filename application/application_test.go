package application

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestApplication_Run(t *testing.T) {
	for i := 1; i <= 2; i++{
		inFile, _ := os.Open(fmt.Sprintf("tests/%v/%v.txt", i,i))
		defer inFile.Close()

		expectedFile, _ := os.Open(fmt.Sprintf("tests/%v/%va.txt", i, i))
		defer expectedFile.Close()

		expectedErrFile, _ := os.Open(fmt.Sprintf("tests/%v/%vb.txt", i,i))
		defer expectedErrFile.Close()


		expectedOutputBytes, _ := io.ReadAll(bufio.NewReader(expectedFile))
		expectedOutput := string(expectedOutputBytes)

		expectedErrBytes, _ := io.ReadAll(bufio.NewReader(expectedErrFile))
		expectedErr := string(expectedErrBytes)


		tempFileRes, _ := os.CreateTemp(os.TempDir(), "output-res.txt")
		defer os.Remove(tempFileRes.Name()) 

		tempFileErr, _ := os.CreateTemp(os.TempDir(), "output-err.txt")
		defer os.Remove(tempFileErr.Name()) 


		outputBuf := bufio.NewWriter(tempFileRes)
		outputBufErr := bufio.NewWriter(tempFileErr)
		app := NewApplication(bufio.NewReader(inFile), outputBuf, outputBufErr)

		app.Run()
		outputBufErr.Flush()

		actualOutput, _ := os.ReadFile(tempFileRes.Name())
		actualOutputString := strings.TrimSpace(string(actualOutput))
		lines := strings.Split(actualOutputString, "\n")
		if len(lines) > 0 {
			actualOutputString = strings.Join(lines[:len(lines)-1], "\n")
		}

		if actualOutputString != expectedOutput {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutputString)
		}


		actualErr, _ := os.ReadFile(tempFileErr.Name())
		actualErrString := strings.TrimSpace(string(actualErr))

		if actualErrString != expectedErr {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expectedErr, actualErrString)
		}
	}
}