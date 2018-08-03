{{if .isStart}}
    <table class="easyui-datagrid" {{.attrFragment}} title="{{.title}}列表"
        {{if or (not .attrs.id) (equal .attrs.id "dg")}}
           id="dg"
           {{if not .attrs.url}}url="{{.base}}/api/{{.channel}}/{{.table}}/list"{{end}}
        {{end}}
           data-options="fitColumns:true,resizeHandle:'right',pagination:true,singleSelect:true,method:'get'"
        {{if or (not .attrs.id) (equal .attrs.id "dg")}}
           toolbar="#toolbar" tools="#tools"
        {{end}}
           rownumbers="false">
    <thead>
    <tr>
{{else}}
    </tr>
    </thead>
    </table>
{{end}}