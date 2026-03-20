package notes

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO use either sqlite or just the file system
// TODO rename DB -> FS

type FS struct {
	root string
}

func Open(root string) (*FS, error) {
	var abs, err = filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	// TODO check it's a valid directory
	return &FS{
		root: abs,
	}, nil
}

func (f *FS) Write(path string, contents string) {
	log.Printf("FS.Write(%s, %s)", path, contents)
	var absPath = filepath.Join(f.root, path)
	_, err := os.Stat(absPath)
	if err == nil {
		fmt.Fprintf(os.Stderr, "About to overwrite %s\n", path)
	} else if errors.Is(err, os.ErrNotExist) {
		// TODO we must create parents

		var findFirstExistingParent func(string) = nil
		findFirstExistingParent = func (path string) {
			if path == "/" {
				panic("Uh oh, traversed too far up!")
			}
			var parent = filepath.Dir(path)
			fmt.Printf("[DEBUG] checking %s\n", parent)
			_, err := os.Stat(parent)
			if err == nil {
				// We found it!
				fmt.Printf("[DEBUG] Found nearest parent!\n")
				return
			} else if errors.Is(err, os.ErrNotExist) {
				findFirstExistingParent(parent)
				err = os.Mkdir(parent, 0700)
				if err != nil {
					panic(err)
				}
			}
		}
		findFirstExistingParent(absPath)
	} else {
		panic(err)
	}
	if err = os.WriteFile(absPath, []byte(contents), 0600); err != nil {
		panic(err)
	}
}

func (f FS) GetAllPaths() ([]string, error) {
	var dirEntries, err = os.ReadDir(f.root)
	if err != nil {
		return nil, err
	}

	var paths = []string{}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			log.Printf("Found file %s\n", dirEntry.Name())
			paths = append(paths, dirEntry.Name())
		}
	}

	return paths, nil
}
