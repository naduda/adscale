package config

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ConfigurationProperty struct {
	Line    int    `json:"line"`
	Value   string `json:"value"`
	Status  bool   `json:"status"`
	Enabled bool   `json:"enabled"`
}

type Config struct {
	Filename      string
	Props         map[string]*ConfigurationProperty
	CheckedLength int
}

func (c *Config) Init(filename string) error {
	c.Filename = filename
	file, err := os.Open(c.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	line := 0
	c.Props = map[string]*ConfigurationProperty{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())
		arr := strings.Split(text, "=")
		if len(arr) > 1 {
			if strings.HasPrefix(text, "#") && !strings.Contains(text, "=") {
				continue
			}

			key := strings.TrimSpace(arr[0])
			value := strings.TrimSpace(arr[1])
			enabled := !strings.HasPrefix(text, "#")
			if !enabled {
				key = strings.TrimSpace(key[1:])
				if strings.Contains(key, " ") {
					continue
				}
			}
			c.Props[key] = &ConfigurationProperty{line, value, false, enabled}
		}
	}

	return scanner.Err()
}

func (c *Config) Format() error {
	input, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	for i := 1; i < len(lines); i++ {
		l := len(strings.TrimSpace(lines[i]))
		if l > 0 {
			continue
		}
		prev := len(strings.TrimSpace(lines[i-1]))
		for prev == 0 {
			lines[i] = "###"
			i++
			if i == len(lines) {
				break
			}
			prev = len(strings.TrimSpace(lines[i]))
		}
	}

	lines = filter(lines, func(s string) bool {
		return s != "###"
	})

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(c.Filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func (c *Config) IsChecked() bool {
	return c.CheckedLength == len(c.Props)
}
