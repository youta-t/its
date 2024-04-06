package parser

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ParseContext interface {
	Import(importPath string) (*Package, error)
	ImportDir(pkgDir string) (*Package, error)
}

type parseContext struct {
	root string

	// import path => source dir
	require map[string]string
	stdDir  string

	packageCache map[string]*Package
}

func New() (ParseContext, error) {
	gomodPath := os.Getenv("GOMOD")
	ctxt := &parseContext{
		root:         filepath.Dir(gomodPath),
		require:      map[string]string{},
		packageCache: map[string]*Package{},
	}

	goroot := runtime.GOROOT()

	ctxt.stdDir = filepath.Join(goroot, "src")
	cmd := exec.Command(
		filepath.Join(goroot, "bin", "go"),
		"list", "-f", "{{.Path}} @ {{.Dir}}", "-m", "all",
	)

	stdout := new(strings.Builder)
	stderr := new(strings.Builder)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%w: %s", err, stderr)
	}

	for _, line := range strings.Split(stdout.String(), "\n") {
		importPath, dir, ok := strings.Cut(line, " @ ")
		if !ok {
			continue
		}
		ctxt.require[importPath] = dir
	}

	return ctxt, nil
}

type importTarget struct {
	Dir        string
	ImportPath string
}

func (pc *parseContext) detectImportPath(srcDir string) (importTarget, error) {
	targ := importTarget{
		Dir: srcDir,
	}

	// Is it non standard package?
	for module, dir := range pc.require {
		rel, err := filepath.Rel(dir, srcDir)
		if err != nil || strings.HasPrefix(rel, "..") {
			continue
		}
		targ.ImportPath = path.Join(module, filepath.ToSlash(rel))
		return targ, nil
	}
	// Is it in the standard lib?
	if rel, err := filepath.Rel(pc.stdDir, srcDir); err == nil && !strings.HasPrefix(rel, "..") {
		targ.ImportPath = filepath.ToSlash(rel)
		return targ, nil
	}

	return targ, fmt.Errorf("%s is not found in any modules/standard lib", srcDir)
}

func (pc *parseContext) detectSourceDir(importPath string) (importTarget, error) {
	targ := importTarget{
		ImportPath: importPath,
	}

	importPathSlash := importPath + "/"

	for module, dir := range pc.require {
		m := module + "/"
		if !strings.HasPrefix(importPathSlash, m) {
			continue
		}
		rel := filepath.FromSlash(importPath[len(m):])
		targ.Dir = filepath.Join(dir, rel)

		return targ, nil
	}

	targ.Dir = filepath.Join(pc.stdDir, filepath.FromSlash(importPath))
	return targ, nil
}

func (bc *parseContext) Import(importPath string) (*Package, error) {
	if p, ok := bc.packageCache[importPath]; ok {
		return p, nil
	}

	imp, err := bc.detectSourceDir(importPath)
	if err != nil {
		return nil, err
	}
	return bc.doImport(imp)
}

func (bc *parseContext) ImportDir(dir string) (*Package, error) {
	imp, err := bc.detectImportPath(dir)
	if err != nil {
		return nil, err
	}
	if p, ok := bc.packageCache[imp.ImportPath]; ok {
		return p, nil
	}
	return bc.doImport(imp)
}

func (bc *parseContext) doImport(pkg importTarget) (*Package, error) {

	parsed, err := parsePackage(bc, pkg)
	if err != nil {
		return nil, err
	}
	if bc.packageCache == nil {
		bc.packageCache = map[string]*Package{}
	}
	bc.packageCache[parsed.Path] = parsed
	return parsed, nil
}
