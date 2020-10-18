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
	Filename     string
	Alternatives map[string]string
}

const Import = "com.adscale.core.ApplicationConfiguration"
const ImportAll = "com.adscale.core.*"

func (r *Repository) Init(filename string) error {
	fullPath, err := filepath.Abs(filename)
	r.Filename = fullPath
	return err
}

func (r *Repository) CheckProperties(c *config.Config) error {
	if err := r.setPropertyAlternatives(); err != nil {
		return err
	}

	return filepath.Walk(r.Filename, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".java" || strings.HasSuffix(path, "ApplicationConfiguration.java") {
			return nil
		}
		go checkFile(path, c, r.Alternatives)
		return nil
	})
}

func (r *Repository) setPropertyAlternatives() error {
	r.Alternatives = map[string]string{}

	file, err := os.Open(r.Filename + "/base/src/main/java/com/adscale/core/ApplicationConfiguration.java")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "static") && strings.Contains(t, "=") {
			t = t[:len(t)-1]
			t = strings.Replace(t, "public", "", -1)
			t = strings.Replace(t, "final", "", -1)
			t = strings.Replace(t, "static", "", -1)
			t = strings.Replace(t, "String", "", -1)
			split := strings.Split(t, "=")
			key := strings.TrimSpace(split[0])
			val := strings.TrimSpace(split[1])
			val = strings.Replace(val, "\"", "", -1)
			r.Alternatives[val] = key
		}
	}

	return nil
}

func checkFile(filepath string, c *config.Config, alternatives map[string]string) {
	if c.IsChecked() {
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	importBlock := true
	hasProperties := false
	fname := filepath[strings.LastIndex(filepath, "/")+1 : len(filepath)-5]

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		commented := strings.HasPrefix(strings.TrimSpace(t), "//")
		if commented {
			continue
		}

		if importBlock {
			if !hasProperties {
				if strings.Contains(t, "package com.adscale.core;") {
					hasProperties = true
					continue
				}
			}
			if strings.Contains(t, "class ") && strings.Contains(t, fname) {
				importBlock = false
				continue
			}
			if hasProperties {
				continue
			}
			if strings.Contains(t, Import) || strings.Contains(t, ImportAll) {
				hasProperties = true
			}
			continue
		}

		if !hasProperties {
			continue
		}

		for k, v := range c.Props {
			if v.Status {
				continue
			}

			if strings.Contains(t, fmt.Sprintf("\"%s\"", k)) {
				v.Status = true
				c.CheckedLength++
				break
			}
			alternative, ok := alternatives[k]
			if !ok {
				continue
			}
			if strings.Contains(t, fmt.Sprintf("ApplicationConfiguration.%s", alternative)) {
				v.Status = true
				c.CheckedLength++
			}
		}
	}
}
