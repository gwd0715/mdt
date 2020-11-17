package util

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gwd0715/opencc"
)

// translate
func Translate(filename string, to string) {
	//open original file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create target file
	targetFilename, err := nameTarget(filename, to)
	if err != nil {
		fmt.Println(err)
		return
	}
	targetFile, err := os.Create(targetFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer targetFile.Close()

	// translate
	if err := ccTranslate(file, targetFile, to); err != nil {
		fmt.Println(err)
		return
	}
}

// opencc translate
func ccTranslate(in io.Reader, out io.Writer, t string) error {
	if t == "ch" {
		t = "tw2sp"
	} else if t == "tw" {
		t = "s2twp"
	} else {
		return errors.New("Wrong convert config")
	}
	cc, err := opencc.NewOpenCC(t)
	if err != nil {
		return err
	}
	if err := cc.ConvertFile(in, out); err != nil {
		return err
	}
	return nil
}

// generate target file name
func nameTarget(name string, to string) (string, error) {
	ns := strings.Split(name, ".")
	l := len(ns)
	if ns[l-1] != "md" {
		return "", errors.New("This is not markdown file")
	}

	switch to {
	case "ch":
		if ns[l-2] == "zh-tw" {
			ns[l-2] = "zh-ch"
		} else {
			ns[l-1] = "zh-ch"
			ns = append(ns, "md")
		}
	case "tw":
		if ns[l-2] == "zh-ch" {
			ns[l-2] = "zh-tw"
		} else {
			ns[l-1] = "zh-tw"
			ns = append(ns, "md")
		}
	}
	return strings.Join(ns, "."), nil
}
