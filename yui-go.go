package yuicompressor

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type YuiCompressor struct {
	command   []string
	options   map[string]string
	jvmParams string
	jarPath   string
	javaPath  string
}

func NewYuiCompressor() *YuiCompressor {
	return &YuiCompressor{options: make(map[string]string)}
}

func (yuicomp *YuiCompressor) Options(options map[string]string) *YuiCompressor {
	if options != nil {
		yuicomp.useParameters(options, "javapath")
		yuicomp.useParameters(options, "jarpath")
		yuicomp.useParameters(options, "jvmparams")
	}
	return yuicomp
}

func (yuicomp *YuiCompressor) useParameters(options map[string]string, field string) {
	_, err := options[field]
	if err == true {
		yuicomp.options[field] = options[field]
	}
}

func (yuicomp YuiCompressor) MinifyJsReader(reader io.Reader) (string, error) {
	tmpfile := createTmpfile()
	readerAsFile(reader, tmpfile)
	outputStr, err := yuicomp.MinifyJsFile(tmpfile.Name())
	tmpfile.Close()
	return outputStr, err
}

func (yuicomp YuiCompressor) MinifyJsString(jsStr string) (string, error) {
	reader := strings.NewReader(jsStr)
	return yuicomp.MinifyJsReader(reader)
}

func (yuicomp YuiCompressor) MinifyJsFile(filename string) (string, error) {
	yuicomp.generateFullCommand()
	command_array := append(yuicomp.command, "--type", "js", "--nomunge", filename)
	output, err := exec.Command(command_array[0], command_array[1:]...).Output()
	outputStr := string(output[:])
	return outputStr, err
}

func (yuicomp YuiCompressor) MinifyCssString(cssStr string) (string, error) {
	reader := strings.NewReader(cssStr)
	return yuicomp.MinifyCssReader(reader)
}

func (yuicomp YuiCompressor) MinifyCssReader(reader io.Reader) (string, error) {
	tmpfile := createTmpfile()
	readerAsFile(reader, tmpfile)
	outputStr, err := yuicomp.MinifyCssFile(tmpfile.Name())
	tmpfile.Close()
	return outputStr, err
}

func (yuicomp YuiCompressor) MinifyCssFile(filename string) (string, error) {
	yuicomp.generateFullCommand()
	command_array := append(yuicomp.command, "--type", "css", filename)
	output, err := exec.Command(command_array[0], command_array[1:]...).Output()
	outputStr := string(output[:])
	return outputStr, err
}

func (yuicomp *YuiCompressor) Command() string {
	yuicomp.generateFullCommand()
	return strings.Join(yuicomp.command, " ")
}

func (yuicomp *YuiCompressor) generateFullCommand() {
	yuicomp.command = []string{yuicomp.getOptionField("javapath", getDefaultJavaPath)}
	_, e := yuicomp.options["jvmparams"]
	if e == true {
		yuicomp.command = append(yuicomp.command, yuicomp.options["jvmparams"])
	}

	yuicomp.command = append(yuicomp.command, "-jar", yuicomp.getOptionField("jarpath", getDefaultJarPath))
}

type OptionFallback func() string

func (yuicomp YuiCompressor) getOptionField(field string, fb OptionFallback) string {
	v, e := yuicomp.options[field]
	if e == true {
		return v
	} else {
		return fb()
	}
}

func getDefaultJarPath() string {
	// The only trick I found to get the package directory where the jarfile is stored
	_, filename, _, _ := runtime.Caller(1)
	jarPath := filepath.Dir(filename) + "/yuicompressor-2.4.8.jar"
	return jarPath
}

func getDefaultJavaPath() string {
	java_path, err := exec.LookPath("java")
	if err != nil {
		panic("Please install Java on your system.")
	}
	return java_path
}

func createTmpfile() *os.File {
	tmpfile, err := ioutil.TempFile("/tmp", "yui_compress")
	if err != nil {
		panic("Impossible to create a temporary file in /tmp.")
	}
	return tmpfile
}

func readerAsFile(reader io.Reader, tmpfile *os.File) {
	tmpstring := make([]byte, 1024)
	n, err := reader.Read(tmpstring)
	tmpfile.Write(tmpstring[:n])
	for err != nil {
		n, err = reader.Read(tmpstring)
		if err != nil {
			tmpfile.Write(tmpstring[:n])
		}
	}
}
