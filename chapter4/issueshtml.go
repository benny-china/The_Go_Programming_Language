package main
import (
    "html/template"
    "ch4/github"
    "os"
    "log"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issue</h1>
<table>
<tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
</tr>
{{range.Items}}
<tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
    <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
    result,err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    if err := issueList.Execute(os.Stdout, result);err != nil {
        log.Fatal(err)
    }
}
