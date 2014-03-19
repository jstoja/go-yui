package yuicompressor

import (
	"os"
	"testing"
)

const	fixture_css = (
`div.warning {
	display: none;
}

div.error {
	background: red;
	color: white;
}

@media screen and (max-device-width: 640px) {
	body { font-size: 90%; }
}`)

const fixture_js = (
`// here's a comment
var Foo = { "a": 1 };
Foo["bar"] = (function(baz) {
	/* here's a
	multiline comment */
	if (false) {
		doSomething();
	} else {
		for (var index = 0; index < baz.length; index++) {
			doSomething(baz[index]);
		}
	}
})("hello");`)

func TestUseJarPath(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	if yc.Command() != "/usr/bin/java -jar ./yuicompressor-2.4.8.jar" {
		t.Error("Impossible to set a new jar_path: " + yc.Command())
	}
}

func TestUseJavaPath(t *testing.T) {
	yc := New()
	yc.UseJavaPath("/var/test/path/java")
	yc.UseJarPath("./yuicompressor-2.4.8.jar")

	expected_command := "/var/test/path/java -jar ./yuicompressor-2.4.8.jar"

	if yc.Command() != expected_command {
		t.Error("Impossible to set a new java_path: " + yc.Command())
	}
}

func TestUseJvmOptions(t *testing.T) {
	yc := New()
	yc.UseJavaPath("/usr/bin/java")
	yc.UseJvmOptions("-Xms64M -Xmx64M")
	yc.UseJarPath("./yuicompressor-2.4.8.jar")

	expected_command := "/usr/bin/java -Xms64M -Xmx64M -jar ./yuicompressor-2.4.8.jar"
	if yc.Command() != expected_command {
		t.Error("Impossible to set jvm opts: " + yc.Command())
	}

}

func TestValidity(t *testing.T) {
	data_uri_css := `div {
	background: white url(\'data:image/png;base64,iVBORw0KGgoAAAANSUhEU
	gAAABAAAAAQAQMAAAAlPW0iAAAABlBMVEUAAAD///+l2Z/dAAAAM0lEQVR4nGP4/5/h
	/1+G/58ZDrAz3D/McH8yw83NDDeNGe4Ug9C9zwz3gVLMDA/A6P9/AFGGFyjOXZtQAAA
	AAElFTkSuQmCC\') no-repeat scroll left top;}`
	
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	_, err := yc.MinifyCssString(data_uri_css)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestMinifyCss(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	output, err := yc.MinifyCssString(fixture_css)
	if err != nil {
		t.Error(err)
	}
	if output != "div.warning{display:none}div.error{background:red;color:white}@media screen and (max-device-width:640px){body{font-size:90%}}" {
		t.Error("The CSS should be compressed and it's not.")
	}
}

func TestMinifyCssReader(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	fd, err := os.Open("assets_test/test1.css")
	output, err := yc.MinifyCssReader(fd)
	if err != nil {
		t.Error(err)
	}
	if output != "div.warning{display:none}div.error{background:red;color:white}@media screen and (max-device-width:640px){body{font-size:90%}}" {
		t.Error("The JS should be compressed with a stream and it's not.")
	}
}

func TestMinifyCssFile(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	output, err := yc.MinifyCssFile("assets_test/test1.css")
	if err != nil {
		t.Error(err)
	}
	if output != "div.warning{display:none}div.error{background:red;color:white}@media screen and (max-device-width:640px){body{font-size:90%}}" {
		t.Error("The JS should be compressed with a stream and it's not.")
	}
}

func TestMinifyJs(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	output, err := yc.MinifyJsString(fixture_js)
	if err != nil {
		t.Error(err)
	}
	if output != "var Foo={a:1};Foo.bar=(function(baz){if(false){doSomething()}else{for(var index=0;index<baz.length;index++){doSomething(baz[index])}}})(\"hello\");" {
		t.Error("The JS should be compressed and it's not.")
	}
}

func TestMinifyJsReader(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	fd, err := os.Open("assets_test/test1.js")
	output, err := yc.MinifyJsReader(fd)
	if err != nil {
		t.Error(err)
	}
	if output != "var Foo={a:1};Foo.bar=(function(baz){if(false){doSomething()}else{for(var index=0;index<baz.length;index++){doSomething(baz[index])}}})(\"hello\");" {
		t.Error("The JS should be compressed with a stream and it's not.")
	}
}

func TestMinifyJsFile(t *testing.T) {
	yc := New()
	yc.UseJarPath("./yuicompressor-2.4.8.jar")
	output, err := yc.MinifyJsFile("assets_test/test1.js")
	if err != nil {
		t.Error(err)
	}
	if output != "var Foo={a:1};Foo.bar=(function(baz){if(false){doSomething()}else{for(var index=0;index<baz.length;index++){doSomething(baz[index])}}})(\"hello\");" {
		t.Error("The JS should be compressed with a stream and it's not.")
	}
}