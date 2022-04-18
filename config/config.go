package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/yookoala/realpath"
	"gopkg.in/yaml.v2"

	"kubecm/utils"
)

var (
	config *Config
	mutex  sync.Mutex
)

type Config struct {
	Current string            `mapstructure:"current,omitempty" json:"current,omitempty" yaml:"current,omitempty"`
	Configs map[string]string `mapstructure:"configs,omitempty" json:"configs,omitempty" yaml:"configs,omitempty"`
}

func load() *Config {

	config := &Config{
		Configs: map[string]string{},
	}
	configPath, err := homedir.Expand(AppRC)
	if err != nil {
		log.Fatal(err)
	}

	d, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(d, config)
	if err != nil {
		fmt.Println("=================")
		log.Fatalf("error: %v", err)
	}

	return config
}

func GetConfig() *Config {

	if config != nil {
		return config
	}

	mutex.Lock()
	defer mutex.Unlock()

	// double check
	if config != nil {
		return config
	}

	// todo load config from file
	config = load()
	if config != nil {
		return config
	}

	return &Config{
		Configs: map[string]string{},
	}
}

// SetCurrent set current config
func (c *Config) SetCurrent(name string) *Config {

	c.Current = name

	return c
}

func copy(src string, dest string) {

	data, err := ioutil.ReadFile(src)
	utils.CheckErr(err)

	err = ioutil.WriteFile(dest, data, 0644)
	utils.CheckErr(err)
}

// Add 新增配置
func (c *Config) Add(name, path string, move bool) *Config {

	path, err := homedir.Expand(path)
	utils.CheckErr(errors.WithMessage(err, "展开文件路径"))

	path, err = realpath.Realpath(path)
	utils.CheckErr(errors.WithMessage(err, "获取 kubeconfig 文件绝对路径失败"))

	base := filepath.Base(path)

	dest, err := homedir.Expand(filepath.Join(KubeConfigVault, base))
	utils.CheckErr(err)

	c.Configs[name] = dest

	if !move {
		copy(path, dest)
	} else {
		err := os.Rename(path, dest)
		utils.CheckErr(errors.Wrap(err, fmt.Sprintf("移动文件失败: %s -> %s", path, dest)))
	}

	if c.Current == "" {
		c.Current = name
	}

	return c
}

// List config
func (c *Config) List() {

	if len(c.Configs) == 0 {
		fmt.Println(fmt.Sprintf("未找到受管的配置，请先通过 kubecm add 添加配置"))
	}

	current := c.Current
	greenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

	names := make([]string, 0, len(c.Configs))
	for k := range c.Configs {
		names = append(names, k)
	}
	sort.Strings(names)

	fmt.Println()
	fmt.Println(strings.Repeat("=", WordWrapColumn))

	for _, name := range names {
		config := c.Configs[name]
		star := strings.Repeat(" ", len("(current)"))
		if name == current {
			star = greenBold("(current)")
		}
		fmt.Println(fmt.Sprintf("%s %-24s %s", star, name, config))
	}

	fmt.Println(strings.Repeat("=", WordWrapColumn))
	fmt.Println()
}

// Del Delete specified config
func (c *Config) Del(name string) *Config {

	current := c.Current

	if current == name {
		c.Current = "-"
	}
	delete(c.Configs, name)

	return c
}

// Switch to specified config
func (c *Config) Switch(name string) *Config {

	if _, ok := c.Configs[name]; ok {
		c.Current = name
	} else {
		fmt.Println("别名有误，找不到指定名称的配置，操作被忽略")
		fmt.Println()
	}

	return c
}

// Rename old to new
func (c *Config) Rename(old, new string) *Config {

	current := c.Current

	if path, ok := c.Configs[old]; ok {
		c.Configs[new] = path
		delete(c.Configs, old)

		if current == old {
			c.Current = new
		}
	} else {
		fmt.Println("找不到指定别名的 config，操作被忽略")
		println()
	}

	return c
}

// Desc show content of specified config
func (c *Config) Desc(name string) {

	if path, ok := c.Configs[name]; ok {
		utils.Cat(path)
	} else {
		utils.CheckErr(errors.New("找不到指定别名的配置"))
	}
}

// Sync 根据最新的 kubecm 配置同步 symlink
func (c *Config) Sync() {

	d, err := yaml.Marshal(c)
	utils.CheckErr(errors.WithMessage(err, "数据 marshal 失败"))

	configPath, err := homedir.Expand(AppRC)
	utils.CheckErr(errors.WithMessagef(err, "展开路径 %s 失败", AppRC))

	defaultKubeConfig, err := homedir.Expand(DefaultKubeConfig)
	utils.CheckErr(errors.WithMessagef(err, "展开路径 %s 失败", DefaultKubeConfig))

	// 删掉当前 config 时，current = "-"
	// 首次添加 config 时，current = ""
	if c.Current != "-" && c.Current != "" {
		if _, err := os.Lstat(defaultKubeConfig); err == nil {
			err = os.Remove(defaultKubeConfig)
			utils.CheckErr(errors.WithMessagef(err, "移除软链 %s 失败", defaultKubeConfig))
		}

		err = os.Symlink(c.Configs[c.Current], defaultKubeConfig)
		utils.CheckErr(errors.WithMessagef(err, "创建软链 %s 失败", defaultKubeConfig))
	}

	// 删掉当前 config 时，current = "-"
	if c.Current == "-" {
		if _, err := os.Lstat(defaultKubeConfig); err == nil {
			err = os.Remove(defaultKubeConfig)
			utils.CheckErr(errors.WithMessagef(err, "移除软链 %s 失败", defaultKubeConfig))
		}
	}

	err = os.WriteFile(configPath, d, 0644)
	utils.CheckErr(errors.WithMessagef(err, "写入文件 %s 失败", configPath))
}
