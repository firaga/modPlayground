package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// mergeConfigs 合并多个 YAML 配置文件，相同的键以后出现的值为准
func mergeConfigs(filePaths []string) (map[string]interface{}, error) {
	finalConfig := make(map[string]interface{})

	for _, filePath := range filePaths {
		// 读取 YAML 文件内容
		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", filePath, err)
		}

		// 解析 YAML 文件
		var config map[string]interface{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal yaml from file %s: %v", filePath, err)
		}

		// 合并配置到 finalConfig
		mergeMaps(finalConfig, config)
	}

	return finalConfig, nil
}

// mergeMaps 将 src 合并到 dst，支持嵌套结构和字段数量不一致的情况
func mergeMaps(dst, src map[string]interface{}) {
	for key, value := range src {
		if srcMap, ok := value.(map[string]interface{}); ok {
			// 如果值是 map，则递归合并
			if dstMap, ok := dst[key].(map[string]interface{}); ok {
				mergeMaps(dstMap, srcMap)
			} else {
				// 如果 dst 中不存在该键，或者该键不是 map，则直接覆盖
				dst[key] = srcMap
			}
		} else {
			// 如果不是 map，则直接覆盖
			dst[key] = value
		}
	}
}

// writeConfigToFile 将合并后的配置写入文件
func writeConfigToFile(config map[string]interface{}, outputPath string) error {
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
