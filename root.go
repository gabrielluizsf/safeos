package safeos

import (
	"io"
	"io/fs"
	"os"
)

type Root struct {
	Dir string
}

func (r *Root) CreateDir(dirname string) error {
	root, err := r.init()
	if err != nil {
		return err
	}
	defer root.Close()
	return root.Mkdir(dirname, 0755)
}

func (r *Root) CreateFile(filePath string, data []byte) error {
	root, err := r.init()
	if err != nil {
		return err
	}
	defer root.Close()
	file, err := root.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func (r *Root) ReadFile(filePath string) ([]byte,  error){
	file, err := os.OpenInRoot(r.Dir, filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func (r *Root) DeleteFile(filePath string) error {
	root, err := r.init()
	if err != nil {
		return err
	}
	defer root.Close()
	return root.Remove(filePath)
}

func (r *Root) Delete() error {
	return os.RemoveAll(r.Dir)
}

func (r *Root) FS() fs.FS {
	root, err := r.init()
	if err != nil {
		return nil
	}
	return root.FS()
}


func (r *Root) Stat(filePath string) (os.FileInfo, error) {
	root, err := r.init()
	if err != nil {
		return nil, err
	}
	defer root.Close()
	return root.Stat(filePath)
}

func (r *Root) init() (*os.Root, error) {
	if _, err := os.Stat(r.Dir); os.IsNotExist(err) {
		_ = os.MkdirAll(r.Dir, 0755)
	}
	return os.OpenRoot(r.Dir)
}