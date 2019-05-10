package honeyfab

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func FindLogs(logDir string) ([]string, error) {
	logPaths := []string{}

	filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("accessing path %q: %v\n", path, err)
		}
		if !info.IsDir() && isStructuredLog(path) {
			logPaths = append(logPaths, path)
		}
		return nil
	})

	return logPaths, nil
}

func isStructuredLog(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("could not open file %q: %v\n", path, err)
	}

	reader := bufio.NewReader(f)

	contents, _, err := reader.ReadLine()
	if err != nil {
		log.Printf("could not read log %q: %v\n", path, err)
		return false
	}

	ignored := make(map[string]interface{})
	err = json.Unmarshal(contents, &ignored)
	if err != nil {
		log.Printf("log %q is not json: %v\n", path, err)
		return false
	}

	return true
}

func WriteHoneytailConf(baseHoneytailConfPath string, newHoneytailConfPath string, logPaths []string) error {
	baseHoneytailConf, err := ioutil.ReadFile(baseHoneytailConfPath)
	if err != nil {
		return fmt.Errorf("reading base honeytail.conf template: %v", err)
	}

	t, err := template.New("honeytail").Parse(string(baseHoneytailConf))
	if err != nil {
		return fmt.Errorf("parsing base template: %v", err)
	}

	newHoneytailConfFile, err := os.Create(newHoneytailConfPath)
	if err != nil {
		return fmt.Errorf("creating new honeytail.conf file: %v", err)
	}
	defer newHoneytailConfFile.Close()

	type HoneytailConf struct {
		LogFiles []string
	}

	err = t.Execute(newHoneytailConfFile, HoneytailConf{LogFiles: logPaths})
	if err != nil {
		return fmt.Errorf("executing template: %v", err)
	}

	return nil
}
