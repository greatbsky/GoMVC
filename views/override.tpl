{{if eq "js" .ext}}
/**
* 解析返回结果，每个项目规范不一致
* @param data
*/
function parseResult(data) {
    return data
}
{{end}}
{{if eq "css" .ext}}
{{end}}
{{.body}}