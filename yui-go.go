package yuicompressor

import (
	"os"
	"strings"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Compressor struct {
	command		[]string
	jvmParams string
	jarPath		string
	javaPath 	string
}

func New() (*Compressor) {
	return &Compressor{}
}

func (yuicomp Compressor) MinifyJs(jsStr string) (string, error) {
	tmpfile := createTmpfile()

	tmpfile.WriteString(jsStr)
	tmpfile.Sync()

	yuicomp.generateFullCommand()
	command_array := append(yuicomp.command, "--type", "js", "--nomunge", tmpfile.Name())
	output, err := exec.Command(command_array[0], command_array[1:]...).Output()
	outputStr := string(output[:])

	tmpfile.Close()

	return outputStr, err
}

func (yuicomp Compressor) MinifyCss(cssStr string) (string, error) {
	tmpfile := createTmpfile()

	tmpfile.WriteString(cssStr)
	tmpfile.Sync()

	yuicomp.generateFullCommand()
	command_array := append(yuicomp.command, "--type", "css", tmpfile.Name())
	output, err := exec.Command(command_array[0], command_array[1:]...).Output()
	outputStr := string(output[:])

	tmpfile.Close()

	return outputStr, err
}

func (yuicomp *Compressor) Command() string {
	yuicomp.generateFullCommand()
	return strings.Join(yuicomp.command, " ")
}

func (yuicomp *Compressor) UseJavaPath(javaPath string) {
	yuicomp.javaPath = javaPath
}

func (yuicomp *Compressor) UseJarPath(jarPath string) {
	yuicomp.jarPath = jarPath
}

func (yuicomp *Compressor) UseJvmOptions(jvmParams string) {
	yuicomp.jvmParams = jvmParams
}

func (yuicomp *Compressor) generateFullCommand() {
	yuicomp.command = []string{yuicomp.getJavaPath()}
	if yuicomp.jvmParams != "" {
		yuicomp.command = append(yuicomp.command, yuicomp.jvmParams)
	}

	yuicomp.command = append(yuicomp.command, "-jar", yuicomp.getJarPath())
}

func (yuicomp Compressor) getJavaPath() string {
	if yuicomp.javaPath != "" {
		return yuicomp.javaPath
	}
	return getDefaultJavaPath()
}

func (yuicomp Compressor) getJarPath() string {
	if yuicomp.jarPath != "" {
		return yuicomp.jarPath
	}
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

func createTmpfile() (*os.File) {
	tmpfile, err := ioutil.TempFile("/tmp", "yui_compress")
	if err != nil {
		panic("Impossible to create a temporary file in /tmp.")
	}
	return tmpfile
}