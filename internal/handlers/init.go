package handlers

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/enorith/enocli/internal/pkg/helpers"
	"github.com/enorith/supports/file"
	"github.com/go-git/go-git/v5"
)

var EnorithMod = "github.com/enorith/enorith"

func InitCommand(dir, module string) error {
	if ok, _ := file.PathExists(dir); ok {
		return fmt.Errorf("dir [%s] exists", dir)
	}

	_, e := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://%s", EnorithMod),
		Progress: os.Stdout,
	})
	if e != nil {
		return e
	}

	fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			realpath := filepath.Join(dir, path)
			helpers.FileReplaceContent(realpath, []byte(EnorithMod), []byte(module))
		}
		return nil
	})

	modPath := filepath.Join(dir, "go.mod")
	helpers.FileReplaceContent(modPath, []byte(EnorithMod), []byte(module))
	dotGit := filepath.Join(dir, ".git")
	os.RemoveAll(dotGit)
	return nil
}
