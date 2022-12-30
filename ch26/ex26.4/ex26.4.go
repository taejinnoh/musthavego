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
	//word := "brother"
	fmt.Println("찾으려는 단어:", word)

	files := os.Args[2:]
	//files := []string{"*.txt"}
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
	fileInfos := []FileInfo{}

	filelist, err := filepath.Glob(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return fileInfos
	}

	ch := make(chan FileInfo)
	cnt := len(filelist)
	receiveCnt := 0

	for _, filename := range filelist {
		fmt.Println("goroutine start")
		go FindWordInFile(word, filename, ch)
	}

	if len(filelist) == 0 {
		fmt.Println("파일이 없습니다.")
		close(ch)
	}

	for fileInfo := range ch {
		fmt.Println("channel output")
		fileInfos = append(fileInfos, fileInfo)
		receiveCnt++
		if receiveCnt == cnt {
			break
		}
	}
	return fileInfos
}

func FindWordInFile(word, filename string, ch chan FileInfo) {
	fileInfo := FileInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- fileInfo
		return
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
	ch <- fileInfo
	fmt.Println("channel input")
}
