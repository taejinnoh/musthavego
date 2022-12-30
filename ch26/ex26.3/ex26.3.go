package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FileInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}
	word := os.Args[1]
	fmt.Println("찾으려는 단어:", word)
	files := os.Args[2:]
	fileInfos := []FileInfo{}
	for _, path := range files {
		fileInfos = append(fileInfos, FindWordInAllFiles(word, path)...)
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.filename)
		fmt.Println("------------------------------")
		for _, lineInfo := range fileInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("------------------------------\n")
	}
}

func FindWordInAllFiles(word, path string) []FileInfo {
	var fileInfos []FileInfo

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return fileInfos
	}
	for _, filename := range filelist {
		fileInfos = append(fileInfos, FindWordInFile(word, filename))
	}
	return fileInfos
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInFile(word, filename string) FileInfo {
	fileInfo := FileInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return fileInfo
	}
	defer file.Close()

	lineNo := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if strings.Contains(line, word) {
			fileInfo.lines = append(fileInfo.lines, LineInfo{lineNo, line})
		}
	}
	return fileInfo
}
