# 4.6.文本和HTML模板
## 文本
```go 
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ)
if err != nil {
    log.Fatal(err)
}
```
## html生成
```go 
var data struct {
        A string        // untrusted plain text
        B template.HTML // trusted HTML
    }
```
A会转译失效
## [Golang template 小抄](https://colobu.com/2019/11/05/Golang-Templates-Cheatsheet/)
### 创建模版
- template.ParseFiles(filenames)可以解析一组模板，使用文件名作为模板的名字。
- template.ParseGlob(pattern)会根据pattern解析所有匹配的模板并保存。

### 执行模版
- 简单的模板tpl可以通过tpl.Execute(io.Writer, data)去执行
- tpl.ExecuteTemplate(io.Writer, name, data)和上面的简单模板类似，只不过传入了一个模板的名字，指定要渲染的模板(因为tpl可以包含多个模板)。

### 变量
传给模板的数据可以存在模板中的变量中，在整个模板中都能访问。 比如 { {$number := .}}, 我们使用$number作为变量，保存传入的数据，可以使用{ {$number}}来访问变量。
### 动作
#### if
```tmpl 
<h1>Hello, { {if .Name}} { {.Name}} { {- else}} Anonymous { {end}}!</h1>

```
告诉模板移除 .Name变量之间的空格。我们在end关键字中也加入减号。
#### range
```tmpl
{ {range .Items}}
  <div class="item">
    <h3 class="name">{ {.Name}}</h3>
    <span class="price">${ {.Price}}</span>
  </div>
{ {end}}

```
### 函数
1. 获取索引
```tmpl 
<body>
    <h1> { {index .FavNums 2 }}</h1>
</body>
```
2. and
```tmpl 
{ {if and .User .User.Admin}}
  You are an admin user!
{ {else}}
  Access denied!
{ {end}}
```
3. not 
```tmpl 
{ { if not .Authenticated}}
  Access Denied!
{ { end }}
```
4. 多个函数用"｜"分开
### 比较
eq: arg1 == arg2
ne: arg1 != arg2
lt: arg1 < arg2
le: arg1 <= arg2
gt: arg1 > arg2
ge: arg1 >= arg2
### 嵌套模版和布局
定义：
```tmpl 
{ {define "footer"}}
<footer> 
	<p>Here is the footer</p>
</footer>
{ {end}}
```
使用：
```tmpl 
{ {template "footer"}}
```
### 模版间传递变量
```tmpl 
{ {define "header"}}
	<h1>{ {.}}</h1>
{ {end}}
// Call template and pass a name parameter
{ {range .Items}}
  <div class="item">
    { {template "header" .Name}}
    <span class="price">${ {.Price}}</span>
  </div>
{ {end}}
```
### 创建布局
```go 
// Omitted imports & package
var LayoutDir string = "views/layouts"  
var bootstrap *template.Template
func main() {
	var err error
	bootstrap, err = template.ParseGlob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	bootstrap.ExecuteTemplate(w, "bootstrap", nil)
}

```
### 模版调用函数
```go 
func (u User) HasPermission(feature string) bool {  
  if feature == "feature-a" {
    return true
  } else {
    return false
  }
}
```
```tmpl
{ {if .User.HasPermission "feature-a"}}
  <div class="feature">
    <h3>Feature A</h3>
    <p>Some other stuff here...</p>
  </div>
{ {else}}
  <div class="feature disabled">
    <h3>Feature A</h3>
    <p>To enable Feature A please upgrade your plan</p>
  </div>
{ {end}}
```
```go 
// Example of creating a ViewData
vd := ViewData{
		User: User{
			ID:    1,
			Email: "curtis.vermeeren@gmail.com",
			// Create the HasPermission function
			HasPermission: func(feature string) bool {
				if feature == "feature-b" {
					return true
				}
				return false
			},
		},
	}

```
```tmpl 
{ {if (call .User.HasPermission "feature-b")}}
  <div class="feature">
    <h3>Feature B</h3>
    <p>Some other stuff here...</p>
  </div>
{ {else}}
  <div class="feature disabled">
    <h3>Feature B</h3>
    <p>To enable Feature B please upgrade your plan</p>
  </div>
{ {end}}
```
#### 自定义函数
```go 
// Creating a template with function hasPermission
testTemplate, err = template.New("hello.gohtml").Funcs(template.FuncMap{
    "hasPermission": func(user User, feature string) bool {
      if user.ID == 1 && feature == "feature-a" {
        return true
      }
      return false
    },
  }).ParseFiles("hello.gohtml")
```
