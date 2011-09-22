package ssgen

import (
	"os"
	"io"
	"io/ioutil"
	"path/filepath"
)

func Convert(dirPath, outputDir string, convertors []Converter) os.Error {
	// get temp dir to write to first.
	tempDir, err := ioutil.TempDir("", "bwg")
	if err != nil {
		return err
	}
	err = convert(dirPath, tempDir, convertors)
	if err != nil {
		return err
	}

	//TODO: move the temp dir to actual dir.
	// move? or copy contents? Prolly move (backup old one to dir.prev or
	// somethign, just in case.
	if exists(outputDir) {
		//if outputDir exists, create a bk, just in case
		err = os.Rename(outputDir, outputDir+".bwgbk")
		if err != nil {
			return err
		}
	}
	return os.Rename(tempDir, outputDir)
}

func convert(directory, targetDirectory string, convertors []Converter) os.Error {
	// make sure directory exists.
	err := os.MkdirAll(targetDirectory, 0777)
	if err != nil {
		return err
	}

	f, err := os.Open(directory)
	if err != nil {
		return err
	}
	filenames, err := f.Readdirnames(-1)
	f.Close()

	for _, name := range filenames {
		from := filepath.Join(directory, name)
		to := filepath.Join(targetDirectory, name)

		switch info, err := os.Stat(from); {
		case err != nil:
			return err
		case info.IsDirectory():
			er := convert(from, to, convertors)
			if er != nil {
				return er
			}
		case info.IsRegular():
			er := convertFile(from, to, convertors)
			if er != nil {
				return er
			}
		}
	}

	return nil
}

func convertFile(from, to string, convertors []Converter) os.Error {
	for _, c := range convertors {
		if c.Supports(from) {
			new_filename := c.GetNewFilename(to)
			// get io.Reader for input file
			r, err := os.Open(from)
			if err != nil {
				return err
			}
			// get io.Writer for output file
			w, err := os.Create(new_filename)
			if err != nil {
				return nil
			}
			// convert
			_, err = c.Convert(r, w)
			return err
		}
	}

	// didn't find any convertors that support this file
	//  so, just copy it over.
	return copyFile(from, to)
}

func copyFile(from, to string) os.Error {
	f, err := os.Open(from)
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := os.Create(to)
	if err != nil {
		return err
	}
	defer t.Close()

	_, err = io.Copy(f, t)
	return err
}

func exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return false
}
