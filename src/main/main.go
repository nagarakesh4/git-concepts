package main

import (
	"net/http"
	"io/ioutil"
	"text/template"  //to create used template
)

func main() {
	//handling routes through Handle - it implements the HTTP Handler interface
	http.Handle("/handle", new(MyHandler))
	
	//handling routes through handlefunc
	http.HandleFunc("/handlefunc", func(w http.ResponseWriter, request *http.Request){
			w.Write([]byte("Hello World - Venkata Go Web Prog"))
	})
	
	//handling route for static HTML content
	http.HandleFunc("/static/html", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		template, err := template.New("helloWorld").Parse(staticDoc);
		if err == nil{
			//send the template browser using response write with no data parameters
			template.Execute(w, nil)
		}
	})
	
	//handling route for dynamic HTML content
	http.HandleFunc("/dynamic/html", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		template, err := template.New("helloWorld").Parse(dynamicDoc);
		if err == nil{
			//send the template browser using response write with data parameters
			//template.Execute(w, request.URL.Path)
			template.Execute(w, "Venkata Buddhiraju from San Jose!")
		}
	})
	
	//handling route for dynamic HTML content with data objects passed
	http.HandleFunc("/dynamic/html/obj", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		template, err := template.New("helloWorld").Parse(dynamicObjDoc);
		if err == nil{
			
			//initialize the Context struct
			context := Context{"Venkata", "San Jose"}
			
			//send the template browser using response write with data parameters
			//template.Execute(w, request.URL.Path)
			template.Execute(w, context)
		}
	})
	
	//handling route for branching templates logic
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		template, err := template.New("helloWorld").Parse(branchingLogicDoc);
		if err == nil{
			//send the template browser using response write with data parameters
			template.Execute(w, request.URL.Path)
		}
	})
	
	//handling route for loop templates logic
	http.HandleFunc("/loop", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		template, err := template.New("helloWorld").Parse(loopDoc);
		if err == nil{
			
			//initialize the loopContext struct
			loopContext := loopContext{
				[3]string{"apple", "banana", "orange"},
				"Golang Fruits App", //here , is required for two reasons
				//easy when adding more elements to the initializer
				//auto semi colon insertion by Golang if no comma so insert comma
			}
			
			
			//send the template browser using response write with data parameters
			template.Execute(w, loopContext)
		}
	})
	
	//handling route for sub-templates logic
	http.HandleFunc("/subTemplate", func(w http.ResponseWriter, request *http.Request){
		w.Header().Add("Content-Type", "text/html")
		
		//create a new template using New() and pass in the name of the template (can be any name) and invoke the parse
		//method and provide the below doc constant
		//template, err := template.New("helloWorld").Parse(loopDoc);
		templates := template.New("template")
		templates.New("bodyTemplate").Parse(subTemplateDoc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		
		//initialize the loopContext struct
		loopContext := loopContext{
			[3]string{"apple", "banana", "orange"},
			"Golang Fruits App", //here , is required for two reasons
			//easy when adding more elements to the initializer
			//auto semi colon insertion by Golang if no comma so insert comma
		}
		
		//send the template browser using response write with data parameters
		templates.Lookup("bodyTemplate").Execute(w, loopContext)
	})
	
	http.ListenAndServe(":8200", nil)
}

//create a custom struct that composes (extends) the http.Handler interface
type MyHandler struct{
	
	//use GO's composition system to make it honor the http.Handler interface
	http.Handler
	
	//we don't need to explicitly state this interface, since GO accepts it as an HTTP Handler by just 
	//implementing the interface
	
}

//implement the above interface - http.Handler
//func (this *MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request){
func (MyHandler) ServeHTTP(w http.ResponseWriter, request *http.Request){
	
	path := "public"+request.URL.Path
	println(path)
	
	//copy the contents of the file to data variable
	data, err := ioutil.ReadFile(string(path))
	
	if err == nil{
		w.Header().Add("Content-Type", "text/plain")
		w.Header().Add("Authorization", "Basic Auth")
		w.Write(data)
	}else{
		//write the status code in header
		w.WriteHeader(404)
		
		//mention the 404 message on the browser
		w.Write([]byte("404 "+ http.StatusText(404)))
	}
	
}

//set a mutli line string containing the HTML text
const staticDoc = `
<!DOCTYPE html>
<html>
	<head><title>Golang Web Application</title></head>
	<body>
		<h1>Hello Venkata Buddhiraju!</h1>
		<h2>Hello World - Golang</h2><h3>Web application created with Golang static template</h3>
	</body>
</html>
`

const dynamicDoc = `
<!DOCTYPE html>
<html>
	<head><title>Golang Web Application</title></head>
	<body>
		<h1>Hello {{.}}</h1>
		<h2>Hello World - Golang</h2><h3>Web application created with Golang dynamic template</h3>
	</body>
</html>`

const dynamicObjDoc = `
<!DOCTYPE html>
<html>
	<head><title>Golang Web Application</title></head>
	<body>
		<h1>Hello {{.Name}} from {{.City}}</h1>
		<h2>Hello World - Golang</h2><h3>Web application created with Golang dynamic template</h3>
	</body>
</html>`
 
//create a type context to initialize the arbitrary objects
type Context struct{
	Name string
	City string
}

const branchingLogicDoc = `
<!DOCTYPE html>
<html>
	<head><title>Golang Web Application</title></head>
	<body>
		{{if eq . "/venkata"}}
			<h2> This logic is implemented through branches, you are now at /venkata URL </h2>
		{{else}}
			<h2> This logic is implemented through branches, you are now at {{.}} URL </h2>
		{{end}}	
	</body>
</html>`

const loopDoc = `
<!DOCTYPE html>
<html>
	<head><title>{{.Title}}</title></head>
	<body>
		<ul>
			{{range .Fruit}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>`

const subTemplateDoc = `
{{template "header" .Title}}
	<body>
		<ul>
			{{range .Fruit}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
	<head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type loopContext struct{
	Fruit [3]string
	Title string
}
