package wiring

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	wiringapi "go.githedgehog.com/fabric/api/wiring/v1alpha2"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilyaml "k8s.io/apimachinery/pkg/util/yaml"
)

func init() {
	scheme := runtime.NewScheme()
	if err := wiringapi.AddToScheme(scheme); err != nil {
		log.Fatalf("error adding fabricv1alpha1 to the scheme: %#v", err)
	}

	decoder = serializer.NewCodecFactory(scheme).UniversalDeserializer()
}

var decoder runtime.Decoder

func LoadDataFrom(from string) (*Data, error) {
	fromFile := "."

	data, err := New()
	if err != nil {
		return nil, err
	}

	if info, err := os.Stat(from); err == nil && !info.IsDir() {
		fromFile = filepath.Base(from)
		from = filepath.Dir(from)
	}

	log.Println("Loading data from directory (recursively)", from)
	f := os.DirFS(from)
	err = LoadDir(f, fromFile, data)
	if err != nil {
		return nil, errors.Wrap(err, "error loading dir")
	}

	return data, nil
}

func LoadDir(f fs.FS, root string, data *Data) error {
	err := fs.WalkDir(f, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			log.Println("Walking into", path)

			return nil
		}

		if strings.Contains(path, "kustom") || filepath.Ext(path) != ".yaml" {
			log.Println("Skipping file", path)

			return nil
		}

		log.Println("Loading data from", path)

		err = LoadFile(f, path, data)
		if err != nil {
			return errors.Wrapf(err, "error loading file %s", path)
		}

		return nil
	})

	return err
}

func LoadFile(f fs.FS, path string, data *Data) error {
	yamlFile, err := fs.ReadFile(f, path)
	if err != nil {
		return errors.Wrapf(err, "error reading file %s", path)
	}

	multidocReader := utilyaml.NewYAMLReader(bufio.NewReader(bytes.NewReader(yamlFile)))

	for {
		buf, err := multidocReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.Wrapf(err, "error multidoc-parsing file %s", path)
		}

		obj, _, err := decoder.Decode(buf, nil, nil)
		if err != nil {
			return errors.Wrapf(err, "error decoding object from file %s", path)
		}

		switch typed := obj.(type) {
		case *wiringapi.Rack:
			if err := data.Add(typed); err != nil {
				return err
			}
		case *wiringapi.Switch:
			if err := data.Add(typed); err != nil {
				return err
			}
		case *wiringapi.SwitchPort:
			if err := data.Add(typed); err != nil {
				return err
			}
		}
	}

	return nil
}
