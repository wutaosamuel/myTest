package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

const fileChunk = 1 << 16 /* 64kb */

func main() {
	inputDir := "./input"
	outputDir := "./output"
	files := make([]*FileNames, 0);

	inputFiles, err := os.ReadDir(inputDir)
	if err != nil {
		PrintErr("input info error ", err)
	}
	fmt.Println("input files ->")
	for i, f := range inputFiles {
		fmt.Println(strconv.Itoa(i), "->", f.Name())
	}

	fmt.Println()
	fmt.Println(fileChunk)
	for _, f := range inputFiles {
		fileName := filepath.Join(inputDir, f.Name())
		file, err := os.Open(fileName)
		if err != nil {
			PrintErr(fileName, err)
		}
		defer file.Close()

		fileInfo, _ := file.Stat()
		fmt.Println(fileInfo.Name(), fileInfo.Size())

		fn := NewFileNames(inputDir, outputDir, f.Name())
		files = append(files, fn)
	}

	// .jpg
	fmt.Println()
	fileName := filepath.Join(inputDir, inputFiles[0].Name())
	file, err := os.Open(fileName)
	if err != nil {
		PrintErr(fileName, err)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fmt.Println(fileInfo.Name(), fileInfo.Size())
	parts := int(math.Ceil(float64(fileInfo.Size()) / float64(fileChunk)))
	for p := 0; p < parts-1; p++ {
		createPart(file, outputDir, p, fileChunk)
	}
	// last part
	restSize := fileInfo.Size() & (fileChunk - 1)
	fmt.Println(restSize)
	createPart(file, outputDir, parts-1, int(restSize))

	// output
	outputFiles, err := os.ReadDir(outputDir)
	if err != nil {
		PrintErr("input info error ", err)
	}
	fmt.Println()
	fmt.Println("output files ->")
	for i, f := range outputFiles {
		fInfo, err := f.Info()
		if err != nil {
			PrintErr("", err)
		}
		fmt.Println(strconv.Itoa(i), "->", f.Name(), fInfo.Size())
	}

	// merge
}

type FileNames struct {
	inputDir  string
	outputDir string

	FileName    string
	OutputNames []string
}

func NewFileNames(inputDir, outputDir, fileName string) *FileNames {
	return &FileNames{
		inputDir:    inputDir,
		outputDir:   outputDir,
		FileName:    fileName,
		OutputNames: make([]string, 0),
	}
}

// would be separated as 2 parts, return buffer and create part files
func createPart(file *os.File, outputDir string, chunkNo, bufferSize int) error {
	buf := make([]byte, bufferSize)
	fileInfo, _ := file.Stat()
	partName := fileInfo.Name() + "_" + strconv.Itoa(chunkNo)
	partName = filepath.Join(outputDir, partName)

	file.Read(buf)
	if _, err := os.Create(partName); err != nil {
		return err
	}
	return ioutil.WriteFile(partName, buf, 0664)
}

func CutExample() {
	filePath := ""
	output := ""

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	fmt.Printf("%s size is %d\n", filePath, fileSize)

	// NOTE: fixed chunk, can replace by random size
	const fileChunk = 4 * (1 << 20)
	totalParts := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("%s part no. is %d\n", filePath, totalParts)

	for i := uint64(0); i < totalParts; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		partName := filepath.Join(output, strconv.FormatUint(i, 10))
		_, err := os.Create(partName)
		if err != nil {
			fmt.Printf("part size %d\n", i)
			fmt.Println(err)
			os.Exit(1)
		}

		ioutil.WriteFile(partName, partBuffer, 0644)
	}
}

func PrintErr(str string, err error) {
	fmt.Print(str)
	fmt.Println(err)
	os.Exit(1)
}
