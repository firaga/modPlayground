package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// mergeConfigs 合并多个 YAML 配置文件，保留字段顺序
func mergeConfigs(filePaths []string) (*yaml.Node, error) {
	var finalConfig yaml.Node

	for _, filePath := range filePaths {
		// 读取 YAML 文件内容
		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
		}

		// 解析 YAML 文件为 Node
		var config yaml.Node
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal yaml from file %s: %v", filePath, err)
		}

		// 合并配置到 finalConfig
		err = mergeNodes(&finalConfig, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to merge nodes: %v", err)
		}
	}

	return &finalConfig, nil
}

// mergeNodes 将 src 合并到 dst，支持多种节点类型
func mergeNodes(dst, src *yaml.Node) error {
	// 如果目标节点为空文档节点，则直接复制源节点
	if dst.Kind == yaml.DocumentNode && len(dst.Content) == 0 {
		*dst = *src
		return nil
	}

	switch src.Kind {
	case yaml.MappingNode:
		// 源节点是映射节点（键值对）
		if dst.Kind != yaml.MappingNode {
			return fmt.Errorf("destination node is not a mapping node")
		}

		// 构建目标节点的键值映射
		dstMap := make(map[string]*yaml.Node)
		for i := 0; i < len(dst.Content); i += 2 {
			keyNode := dst.Content[i]
			valueNode := dst.Content[i+1]
			if keyNode.Kind == yaml.ScalarNode {
				dstMap[keyNode.Value] = valueNode
			}
		}

		// 遍历源节点并合并
		for i := 0; i < len(src.Content); i += 2 {
			keyNode := src.Content[i]
			valueNode := src.Content[i+1]

			if keyNode.Kind == yaml.ScalarNode {
				if existingValue, exists := dstMap[keyNode.Value]; exists {
					// 如果目标中已存在该键，则递归合并
					err := mergeNodes(existingValue, valueNode)
					if err != nil {
						return err
					}
				} else {
					// 如果目标中不存在该键，则添加到目标节点
					dst.Content = append(dst.Content, keyNode, valueNode)
				}
			}
		}

	case yaml.SequenceNode:
		// 源节点是数组节点
		if dst.Kind != yaml.SequenceNode {
			return fmt.Errorf("destination node is not a sequence node")
		}

		// 将源节点的所有元素追加到目标节点
		dst.Content = append(dst.Content, src.Content...)

	default:
		// 源节点是标量节点或其他类型
		*dst = *src
	}

	return nil
}

// writeConfigToFile 将合并后的配置写入文件
func writeConfigToFile(config *yaml.Node, outputPath string) error {
	// 将配置序列化为 YAML 格式
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to yaml: %v", err)
	}

	// 写入文件
	err = ioutil.WriteFile(outputPath, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config to file %s: %v", outputPath, err)
	}

	return nil
}

// getConfigFilePaths 获取 configs/ 目录下所有的 YAML 文件路径
func getConfigFilePaths(dirPath string) ([]string, error) {
	var filePaths []string

	// 遍历目录
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是 YAML 文件
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %v", dirPath, err)
	}

	return filePaths, nil
}

func main() {
	// 定义配置文件目录路径和输出文件路径
	configDirPath := "configs/"
	outputPath := "merged_config.yaml"

	// 获取 configs/ 目录下的所有 YAML 文件路径
	filePaths, err := getConfigFilePaths(configDirPath)
	if err != nil {
		fmt.Printf("Error getting config file paths: %v\n", err)
		return
	}

	if len(filePaths) == 0 {
		fmt.Println("No YAML files found in the configs/ directory.")
		return
	}

	// 合并配置文件
	finalConfig, err := mergeConfigs(filePaths)
	if err != nil {
		fmt.Printf("Error merging configs: %v\n", err)
		return
	}

	// 将合并后的配置写入文件
	err = writeConfigToFile(finalConfig, outputPath)
	if err != nil {
		fmt.Printf("Error writing merged config to file: %v\n", err)
		return
	}

	fmt.Printf("Configs merged successfully and written to %s\n", outputPath)
}
