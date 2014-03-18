package yuicompressor

import (
	"testing"
)

func fixture_css() string {
	return(
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
}

func fixture_js() string {
	return(
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
}

func fixture_error_js() string {
	return "var x = {class: 'name'};"
}

func TestValidityCommand(t *testing.T) {
	yc := New()
	if yc.Command() != "/usr/bin/java -jar yuicompressor-2.4.8.jar" {
		t.Error("The base command is not valid: " + yc.Command())
	}
}

func TestUseJavaPath(t *testing.T) {
	yc := New()
	yc.UseJavaPath("/var/test/path/java")
	if yc.Command() != "/var/test/path/java -jar yuicompressor-2.4.8.jar" {
		t.Error("Impossible to set a new java_path: " + yc.Command())
	}
}

func TestUseJarPath(t *testing.T) {
	yc := New()
	yc.UseJarPath("/tmp/yui-jar/yui-compressor.jar")
	if yc.Command() != "/usr/bin/java -jar /tmp/yui-jar/yui-compressor.jar" {
		t.Error("Impossible to set a new jar_path: " + yc.Command())
	}
}

func TestUseJvmOptions(t *testing.T) {
	yc := New()
	yc.UseJvmOptions("-Xms64M -Xmx64M")
	if yc.Command() != "/usr/bin/java -Xms64M -Xmx64M -jar yuicompressor-2.4.8.jar" {
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
	_, err := yc.MinifyCss(data_uri_css)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestMinifyCss(t *testing.T) {
	yc := New()
	output, err := yc.MinifyCss(fixture_css())
	if err != nil {
		t.Error(err)
	}
	if output != "div.warning{display:none}div.error{background:red;color:white}@media screen and (max-device-width:640px){body{font-size:90%}}" {
		t.Error("The CSS should be compressed and it's not.")
	}
}

func TestMinifyJs(t *testing.T) {
	yc := New()
	output, err := yc.MinifyJs(fixture_js())
	if err != nil {
		t.Error(err)
	}
	if output != "var Foo={a:1};Foo.bar=(function(baz){if(false){doSomething()}else{for(var index=0;index<baz.length;index++){doSomething(baz[index])}}})(\"hello\");" {
		t.Error("The JS should be compressed and it's not.")
	}
}