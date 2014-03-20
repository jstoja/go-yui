Go-Yui
======

  Go-Yui is a golang **interface** to Yui-Compressor, for minifying Javascript and CSS assets.

**YOU MUST HAVE JAVA 1.4 (AT MIN) INSTALLED SINCE IT USES Yui-Compressor**

[Yui-Compressor GitHub page](http://yui.github.io/yuicompressor/)

#Getting started

After installing Go and setting up your GOPATH, create a golang file.
##To minify some CSS

	package main
	
	import (
		"fmt"

		"github.com/jstoja/go-yui"
	)
	
	func main() {
		yc := NewYuiCompressor().Options(map[string]string{
			"javapath":  "/var/test/path/java",
			"jvmparams": "-Xms64M -Xmx64M"})
	
		output, err := yc.MinifyCssFile("file.css")
		if err != nil {
			panic(err)
		}
		fmt.Println(output)		
	}
	
##To minify some JS

	package main
	
	import (
		"fmt"

		"github.com/jstoja/go-yui"
	)

	func main() {
		yc := NewYuiCompressor().Options(map[string]string{
			"javapath":  "/var/test/path/java",
			"jvmparams": "-Xms64M -Xmx64M"})

		output, err := yc.MinifyJsFile("file.js")
		if err != nil {
			panic(err)
		}
		fmt.Println(output)		
	}
	

Then install the yui-go package using the command:

	go get github.com/jstoja/go-yui

You just have to run it:

	go run example.go

#Features

##Now

It:

* Grabs the yui-compressor jar according to the package directory
* Takes JAVA/JAR/JVM parameters to fit your environment
* Takes a CSS string/stream/file and process it
* Takes a JS string/stream/file and process it
* Uses the yui-compressor jar

##In the near future

It'll:

* Be able to use all options of yui-compressor
* Integrated as a [Martini](https://github.com/codegangsta/martini/) plugin since I think this go-package is **super-cool**

##In the great future

I'd like it:

* Rewritten using pure concurrent Go instead of using the lib
* To be used as both as a package and a cli-tool

	
#Notes about this project

**This is my first go-package and since I'm new to Golang, I would really appreciate all advices about this work.**

#Licence

##Go-Yui code (MIT license)

	The MIT License (MIT)

	Copyright (c) 2014 Julien Bordellier

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

##YUI Compressor (BSD license)

	Copyright (c) 2009, Yahoo! Inc.
	All rights reserved.

	Redistribution and use of this software in source and binary forms,
	with or without modification, are permitted provided that the following
	conditions are met:

	* Redistributions of source code must retain the above
	  copyright notice, this list of conditions and the
	  following disclaimer.

	* Redistributions in binary form must reproduce the above
	  copyright notice, this list of conditions and the
	  following disclaimer in the documentation and/or other
	  materials provided with the distribution.

	* Neither the name of Yahoo! Inc. nor the names of its
	  contributors may be used to endorse or promote products
	  derived from this software without specific prior
	  written permission of Yahoo! Inc.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
	AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
	IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE
	FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
	DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
	SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
	CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
	OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

	This software also requires access to software from the following sources:

	The Jarg Library v 1.0 ( http://jargs.sourceforge.net/ ) is available
	under a BSD License. Copyright (c) 2001-2003 Steve Purcell,
	Copyright (c) 2002 Vidar Holen, Copyright (c) 2002 Michal Ceresna and
	Copyright (c) 2005 Ewan Mellor.

	The Rhino Library ( http://www.mozilla.org/rhino/ ) is dually available
	under an MPL 1.1/GPL 2.0 license, with portions subject to a BSD license.

	Additionally, this software contains modified versions of the following
	component files from the Rhino Library:

	[org/mozilla/javascript/Decompiler.java]
	[org/mozilla/javascript/Parser.java]
	[org/mozilla/javascript/Token.java]
	[org/mozilla/javascript/TokenStream.java]

	The modified versions of these files are distributed under the MPL v 1.1
	( http://www.mozilla.org/MPL/MPL-1.1.html )

