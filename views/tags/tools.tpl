{{if .isStart}}
<div id="{{if .attrs.id}}{{.attrs.id}}{{else}}tools{{end}}">
    <a id="searchicon" href="" class="icon-filter" title="查询" onclick="searchShowHandler()"></a>
{{else}}
</div>
{{end}}