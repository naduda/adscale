package repo

import (
	"adscale-tools/config"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Repository struct {
	Filename string
}

func (r *Repository) Init(filename string) error {
	fullPath, err := filepath.Abs(filename)
	r.Filename = fullPath
	return err
}

func (r *Repository) CheckProperties(c *config.Config) error {
	return filepath.Walk(r.Filename, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".java" {
			return nil
		}
		checkFile(path, c)
		return nil
	})
}

func checkFile(filename string, c *config.Config) {
	if c.IsChecked() {
		fmt.Println("Done!")
		return
	}

	c.Filename = filename
	file, err := os.Open(c.Filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		for k, v := range c.Props {
			if v.Status {
				continue
			}
			if strings.Contains(t, k) {
				v.Status = true
				c.CheckedLength++
				break
			}
		}
	}
}
