package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mattn/go-runewidth"
)

var config = make(map[string]interface{})

func parseConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return err
	}

	return nil
}

func getProblemsWithLabel(dirPath string) ([]Problem, error) {

	var problems []Problem
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println("Error reading directory:", dirPath, err)
		return nil, err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		// fmt.Println(fullPath)
		if entry.IsDir() {
			labelPath := filepath.Join(fullPath, "label.txt")
			if data, err := os.ReadFile(labelPath); err != nil {
				// 忽略没有 label.txt 的目录
				continue
			} else {
				problems = append(problems, *newProblem(entry.Name(), fullPath, string(data)))
			}

		}
	}
	sortProblem(problems)
	return problems, nil

}
func generateMd(problemsMap map[string]([]Problem)) error {

	outputPath := filepath.Join(config["output_dir"].(string), config["output_file"].(string))
	file, err := os.Create(outputPath)
	if err != nil {
		log.Println("Error create file", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	// 保证顺序
	keys := make([]string, 0, len(problemsMap))
	for k := range problemsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		problems := problemsMap[key]

		fmt.Fprintf(writer, "# "+key+"\n\n")

		header := []string{"catagory", "status", "recommend", "difficulty", "name", "alias", "lastupdate"}

		// 定义每列宽度
		columnWidths := []int{12, 10, 12, 12, 50, 30, 15}

		// 打印表头
		for i, col := range header {
			fmt.Fprint(writer, "| "+padRight(col, columnWidths[i]))
		}
		fmt.Fprintln(writer, "|")

		// 打印分隔符
		for _, width := range columnWidths {
			fmt.Fprint(writer, "| "+strings.Repeat("-", width))
		}
		fmt.Fprintln(writer, "|")

		// 打印数据
		for _, row := range problems {

			row.ColorRender()

			line_show := fmt.Sprintf("|%s|%s|%s|%s|%s|%s|%s|", padRight(row.catagory, columnWidths[0]), padRight(row.status, columnWidths[1]), padRight(row.recommend, columnWidths[2]), padRight(row.difficulty, columnWidths[3]),
				padRight(row.name, columnWidths[4]), padRight(row.alias, columnWidths[5]), padRight(row.lastupdate, columnWidths[6]))
			fmt.Println(line_show)
			md_address := "[" + row.name + "](" + row.path + ")"
			line_md := fmt.Sprintf("|%s|%s|%s|%s|%s|%s|%s|", padRight(row.catagory, columnWidths[0]), padRight(row.status, columnWidths[1]), padRight(row.recommend, columnWidths[2]), padRight(row.difficulty, columnWidths[3]),
				md_address, padRight(row.alias, columnWidths[5]), padRight(row.lastupdate, columnWidths[6]))
			fmt.Fprintln(writer, line_md)
		}
	}
	return nil
}

// 根据指定宽度填充空格，左对齐
func padRight(input string, totalWidth int) string {
	currentWidth := runewidth.StringWidth(input)
	// fmt.Println(currentWidth)
	if currentWidth >= totalWidth {
		return input
	}
	return input + strings.Repeat(" ", totalWidth-currentWidth)
}

func main() {
	log.Println("start mark!")

	err := parseConfig()
	if err != nil {
		panic("Error parsing config")
	}

	structProblems, err := getProblemsWithLabel("../../datastruct")
	if err != nil {
		panic("Error finding code dir")
	}
	algoProblems, err := getProblemsWithLabel("../../algorithm")
	if err != nil {
		panic("Error finding code dir")
	}
	specProblems, err := getProblemsWithLabel("../../spical")
	if err != nil {
		panic("Error finding code dir")
	}
	codeProblems, err := getProblemsWithLabel("../../leetcode")
	if err != nil {
		panic("Error finding code dir")
	}

	generateMd(map[string]([]Problem){
		"01数据结构": structProblems,
		"02算法":   algoProblems,
		"03专题":   specProblems,
		"04题目":   codeProblems})
}
