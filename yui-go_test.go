package yuicompressor

import (
	"testing"
)

const (
	fixture_css = (`div.warning {
			display: none;
		}

		div.error {
			background: red;
			color: white;
		}

		@media screen and (max-device-width: 640px) {
			body { font-size: 90%; }
		}`)

	fixture_js = (`// here's a comment
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

	fixture_css_minified = "div.warning{display:none}div.error{background:red;color:white}@media screen and (max-device-width:640px){body{font-size:90%}}"
	fixture_js_minified  = "var Foo={a:1};Foo.bar=(function(baz){if(false){doSomething()}else{for(var index=0;index<baz.length;index++){doSomething(baz[index])}}})(\"hello\");"
)

func TestOptions(t *testing.T) {
	yc := NewYuiCompressor()
	yc.Options(map[string]string{"javapath": "/usr/bin/java"})
}

func TestUseJarPath(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})
	if yc.Command() != "/usr/bin/java -jar ./yuicompressor-2.4.8.jar" {
		t.Error("Impossible to set a new jar_path: " + yc.Command())
	}
}

func TestUseJavaPath(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{
		"javapath": "/var/test/path/java",
		"jarpath":  "./yuicompressor-2.4.8.jar"})

	expected_command := "/var/test/path/java -jar ./yuicompressor-2.4.8.jar"
	if yc.Command() != expected_command {
		t.Error("Impossible to set a new java_path: " + yc.Command())
	}
}

func TestUseJvmOptions(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{
		"javapath":  "/var/test/path/java",
		"jvmparams": "-Xms64M -Xmx64M",
		"jarpath":   "./yuicompressor-2.4.8.jar"})

	expected_command := "/var/test/path/java -Xms64M -Xmx64M -jar ./yuicompressor-2.4.8.jar"
	if yc.Command() != expected_command {
		t.Error("Impossible to set jvm opts: " + yc.Command())
	}

}

func TestValidity(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})

	_, err := yc.MinifyJs().FromString(fixture_js).ToString()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestJSStringtoString(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})

	output, err := yc.MinifyJs().FromString(fixture_js).ToString()
	if err != nil {
		t.Error(err)
	}
	if output != fixture_js_minified {
		t.Error("The JS should be compressed to a String with a String and it's not.")
	}
}

func TestJSFiletoString(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})

	output, err := yc.MinifyJs().FromFile("assets_test/test1.js").ToString()
	if err != nil {
		t.Error(err)
	}
	if output != fixture_js_minified {
		t.Error("The JS should be compressed to a String with a File and it's not.")
	}
}

func TestCSSStringtoString(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})

	output, err := yc.MinifyCss().FromString(fixture_css).ToString()
	if err != nil {
		t.Error(err)
	}
	if output != fixture_css_minified {
		t.Error("The CSS should be compressed to a String with a String and it's not.")
	}
}

func TestCSSFiletoString(t *testing.T) {
	yc := NewYuiCompressor().Options(map[string]string{"jarpath": "./yuicompressor-2.4.8.jar"})

	output, err := yc.MinifyCss().FromFile("assets_test/test1.css").ToString()
	if err != nil {
		t.Error(err)
	}
	if output != fixture_css_minified {
		t.Error("The CSS should be compressed to a String with a File and it's not.")
	}
}
