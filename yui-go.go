package yuicompressor

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"bytes"
)

type OutputFileError struct {
	filename string
}

func (e OutputFileError) Error() string {
	return "Can't create/open the output file: " + e.filename
}

type YuiCompressorInput struct {
	reader 		io.Reader
	filepath 	string
}

type YuiCompressorOutput struct {
	writer 		io.Writer
	filepath	string
	isFile		bool
}

type YuiCompressor struct {
	command   []string
	options   map[string]string
	fromJSCSS	bool
	input 		YuiCompressorInput
	output 		YuiCompressorOutput
}

func NewYuiCompressor() *YuiCompressor {
	return &YuiCompressor{options: make(map[string]string)}
}


func (yuicomp *YuiCompressor) MinifyCss() (*YuiCompressor) {
	yuicomp.fromJSCSS = false
	return yuicomp
}

func (yuicomp *YuiCompressor) MinifyJs() (*YuiCompressor) {
	yuicomp.fromJSCSS = true
	return yuicomp
}

func (yuicomp *YuiCompressor) FromFile(filename string) (*YuiCompressor) {
	file, _ := os.Open(filename)
	yuicomp.input = YuiCompressorInput{file, filename}
	return yuicomp	
}

//TODO: Instead of using a temporary file, use stdin
func (yuicomp *YuiCompressor) FromString(str string) (*YuiCompressor) {
	tmpfile := createTmpfile()
	readerAsFile(strings.NewReader(str), tmpfile)
	yuicomp.input = YuiCompressorInput{tmpfile, tmpfile.Name()}
	return yuicomp
}

//TODO: Instead of using a temporary file, use stdin
func (yuicomp *YuiCompressor) FromReader(reader io.Reader) (*YuiCompressor) {
	tmpfile := createTmpfile()
	readerAsFile(reader, tmpfile)
	yuicomp.input = YuiCompressorInput{tmpfile, tmpfile.Name()}
	return yuicomp
}

func (yuicomp *YuiCompressor) ToString() (string, error) {
	var output string
	bufferstr := bytes.NewBufferString(output)
	yuicomp.output = YuiCompressorOutput{bufferstr, "", false}
	yuicomp.minify()
	return bufferstr.String(), nil
}

func (yuicomp *YuiCompressor) ToFile(filepath string) (string, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return "", OutputFileError{filepath}
	}
	yuicomp.output = YuiCompressorOutput{nil, file.Name(), true}
	yuicomp.minify()
	return yuicomp.output.filepath, nil
}

func (yuicomp *YuiCompressor) minify() (*YuiCompressor) {
	if yuicomp.fromJSCSS == true {
		yuicomp.minifyJs(yuicomp.input, yuicomp.output)
	} else {
		yuicomp.minifyCss(yuicomp.input, yuicomp.output)
	}
	return yuicomp
}

func (yuicomp *YuiCompressor) minifyJs(input YuiCompressorInput, output YuiCompressorOutput) {
	yuicomp.generateFullCommand()
	if output.isFile == true {
		command_array := append(yuicomp.command, "--type", "js", "--nomunge", input.filepath, "-o", output.filepath)
		cmd := exec.Command(command_array[0], command_array[1:]...)
		cmd.Run() // Havn't catched yet
	} else {
		command_array := append(yuicomp.command, "--type", "js", "--nomunge", input.filepath)
		cmd := exec.Command(command_array[0], command_array[1:]...)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start() // Havn't catched yet
		io.Copy(output.writer, stdout)
	}
}

func (yuicomp *YuiCompressor) minifyCss(input YuiCompressorInput, output YuiCompressorOutput) {
	yuicomp.generateFullCommand()
	if output.isFile == true {
		command_array := append(yuicomp.command, "--type", "css", input.filepath, "-o", output.filepath)
		cmd := exec.Command(command_array[0], command_array[1:]...)
		cmd.Run() // Havn't catched yet
	} else {
		command_array := append(yuicomp.command, "--type", "css", input.filepath)
		cmd := exec.Command(command_array[0], command_array[1:]...)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start() // Havn't catched yet
		io.Copy(output.writer, stdout)
	}	
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
