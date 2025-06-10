package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	configDir := flag.String("dir", "configs/", "配置文件目录")
	outputFile := flag.String("output", "merged_config.yaml", "输出文件路径")
	flag.Parse()

	files, err := getConfigFiles(*configDir)
	if err != nil {
		log.Fatalf("获取配置文件失败: %v", err)
	}

	mergedConfig, err := mergeConfigs(files)
	if err != nil {
		log.Fatalf("合并配置失败: %v", err)
	}

	if err := writeConfig(*outputFile, mergedConfig); err != nil {
		log.Fatalf("写入配置失败: %v", err)
	}

	fmt.Printf("配置已成功合并到 %s\n", *outputFile)
}

func getConfigFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func mergeConfigs(files []string) (*yaml.Node, error) {
	var merged *yaml.Node

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("读取文件 %s 失败: %v", file, err)
		}

		var config yaml.Node
		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("解析文件 %s 失败: %v", file, err)
		}

		// 调试输出：打印根节点类型
		fmt.Printf("文件 %s 的根节点类型: %d\n", file, config.Kind)

		// 处理文档节点
		var root *yaml.Node
		if config.Kind == yaml.DocumentNode {
			if len(config.Content) == 0 {
				continue // 空文档，跳过
			}
			root = config.Content[0]
		} else {
			root = &config
		}

		// 确保根内容是映射类型
		if root.Kind != yaml.MappingNode {
			return nil, fmt.Errorf("文件 %s 的根内容不是映射类型 (实际类型: %d)", file, root.Kind)
		}

		if merged == nil {
			merged = root
			continue
		}

		merged = mergeNodes(merged, root)
	}

	return merged, nil
}

func mergeNodes(dst, src *yaml.Node) *yaml.Node {
	// 创建一个映射来跟踪键的位置
	keyPositions := make(map[string]int)
	for i := 0; i < len(dst.Content); i += 2 {
		key := dst.Content[i].Value
		keyPositions[key] = i
	}

	// 处理源节点中的每个键值对
	for i := 0; i < len(src.Content); i += 2 {
		keyNode := src.Content[i]
		valueNode := src.Content[i+1]
		key := keyNode.Value

		// 检查目标节点中是否已存在此键
		if pos, exists := keyPositions[key]; exists {
			// 存在相同键，处理值的合并
			dstValue := dst.Content[pos+1]
			if dstValue.Kind == yaml.MappingNode && valueNode.Kind == yaml.MappingNode {
				// 递归合并嵌套映射
				mergeNodes(dstValue, valueNode)
			} else {
				// 非映射类型直接替换
				dst.Content[pos+1] = valueNode
			}
		} else {
			// 不存在则添加到末尾
			dst.Content = append(dst.Content, keyNode, valueNode)
		}
	}

	return dst
}

func writeConfig(path string, config *yaml.Node) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(path, data, 0644)
}
