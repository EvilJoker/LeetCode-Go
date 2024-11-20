package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var config = make(map[string]interface{})
var problems []Problem

const TABLE_WITH = 20

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

func getProblemsWithLabel() error {
	dirPath := config["leetcode_dir"].(string)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println("Error reading directory:", dirPath, err)
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			labelPath := filepath.Join(fullPath, "label.txt")
			if data, err := os.ReadFile(labelPath); err != nil {
				// 忽略没有 label.txt 的目录
				continue
			} else {
				problems = append(problems, *newProblem(entry.Name(), string(data)))
			}

		}
	}

	return nil

}

func generateMd() error {
	outputPath := filepath.Join(config["output_dir"].(string), config["output_file"].(string))
	file, err := os.Create(outputPath)
	if err != nil {
		log.Println("Error create file", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString("# 进展\n\n")
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	// 表格标题和数据
	fmt_str := fmt.Sprintf("|%%-%ds|%%-%ds|%%-%ds|%%-%ds|%%-%ds|%%-%ds|", TABLE_WITH, TABLE_WITH, TABLE_WITH, TABLE_WITH, TABLE_WITH, TABLE_WITH)

	header := fmt.Sprintf(fmt_str, "catagory", "status", "recommend", "difficulty", "name", "lastupdate")
	tmp := strings.Repeat("-", TABLE_WITH)
	separator := fmt.Sprintf("|%s|%s|%s|%s|%s|%s|", tmp, tmp, tmp, tmp, tmp, tmp)

	// 写入表格头部
	_, err = writer.WriteString(header + "\n")
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	// 写入表格分隔符
	_, err = writer.WriteString(separator + "\n")
	if err != nil {
		log.Println("Error writing to file:", err)
		return err
	}

	// 写入表格数据
	for _, problem := range problems {
		_, err = fmt.Fprintf(writer, fmt_str+"\n", problem.catagory, problem.status, problem.recommend, problem.difficulty, problem.name, problem.lastupdate)
		if err != nil {
			log.Println("Error writing to file:", err)
			return err
		}
	}

	log.Printf("Markdown file %s generated successfully!\n", outputPath)

	return nil

}

func main() {
	log.Println("start mark!")

	err := parseConfig()
	if err != nil {
		panic("Error parsing config")
	}

	err = getProblemsWithLabel()
	if err != nil {
		panic("Error finding code dir")
	}

	sortProblem(problems)

	fmt.Println(problems)

	generateMd()
}
