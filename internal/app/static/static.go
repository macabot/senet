package static

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag"
	"github.com/macabot/senet/internal/app/component/page"
)

func GeneratePages(outputDir string) error {
	pages := map[string]func() *hypp.VNode{
		"index.html": func() *hypp.VNode {
			return page.Home()
		},
		"rules.html": func() *hypp.VNode {
			return page.Rules()
		},
	}

	for path, nodeFunc := range pages {
		if err := generatePage(outputDir, path, nodeFunc); err != nil {
			return err
		}
	}

	return nil
}

const doctype = "<!DOCTYPE html>"

func generatePage(outputDir string, path string, nodeFunc func() *hypp.VNode) error {
	dir := filepath.Dir(path)
	pageOutputDir := filepath.Join(outputDir, dir)
	if err := os.MkdirAll(pageOutputDir, os.FileMode(0775)); err != nil {
		return err
	}

	pageOutputPath := filepath.Join(outputDir, path)
	fileHandle, err := os.Create(pageOutputPath)
	if err != nil {
		return fmt.Errorf("could not create file for page '%s': %w", path, err)
	}
	defer fileHandle.Close()

	node := nodeFunc()
	if node == nil {
		return fmt.Errorf("nodeFunc for page '%s' returned nil node", path)
	}
	if _, err := fileHandle.WriteString(doctype); err != nil {
		return fmt.Errorf("failed to render DOCTYPE for page '%s': %w", path, err)
	}
	if err := tag.Render(fileHandle, node); err != nil {
		return fmt.Errorf("failed to render page '%s': %w", path, err)
	}

	return nil
}
