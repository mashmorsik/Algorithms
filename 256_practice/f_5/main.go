package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeEmpty(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, val := range v {
			if val != nil {
				// Рекурсивно удаляем пустые значения
				if removed := removeEmpty(val); removed != nil {
					result[key] = removed
				}
			}
		}
		// Возвращаем nil, если словарь пустой
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		result := make([]interface{}, 0)
		for _, val := range v {
			if val != nil {
				// Рекурсивно удаляем пустые значения
				if removed := removeEmpty(val); removed != nil {
					result = append(result, removed)
				}
			}
		}
		// Возвращаем nil, если список пустой
		if len(result) == 0 {
			return nil
		}
		return result
	default:
		return v
	}
}

func removeEmptyStr(rawStr string) string {
	var jsonData interface{}
	err := json.Unmarshal([]byte(rawStr), &jsonData)
	if err != nil {
		log.Println(err)
		return ""
	}

	result := removeEmpty(jsonData)
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.ReplaceAll(string(jsonResult), "\t", "")
}

func inputDataIterator(raw string) string {
	var result strings.Builder
	lines := strings.Split(raw, "\n")
	i := 1

	for i < len(lines) {
		// Читаем количество следующих строк
		count, err := strconv.Atoi(strings.TrimSpace(lines[i]))
		if err != nil {
			fmt.Printf("ошибка преобразования числа в строке %d: %v", i+1, err)
			return ""
		}
		i++

		var tempResult strings.Builder
		for j := 0; j < count; j++ {
			tempResult.WriteString(strings.TrimSpace(lines[i]))
			i++
		}

		result.WriteString(tempResult.String() + ",")
	}

	return strings.TrimRight(result.String(), ",")
}

func combiner(str string) string {
	idata := inputDataIterator(str)
	return removeEmptyStr("[" + idata + "]")
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//var numSetsSets int
	var jsonString string
	jsonString, err := in.Read(in)
	if err != nil {
		return
	}

	dataStrs := combiner(jsonString)
	fmt.Fprintln(out, dataStrs)

	//jsonString += strconv.Itoa(numSets)
	//
	//for i := 0; i < numSets; i++ {
	//	var line string
	//
	//	for {
	//		line, _ = in.ReadString('\n')
	//		if line == "\n" {
	//			break
	//		}
	//		jsonString += line
	//	}
	//
	//	fmt.Println(jsonString)
	//
	//	dataStrs := combiner(jsonString)
	//	fmt.Fprintln(out, dataStrs)
	//}
}
