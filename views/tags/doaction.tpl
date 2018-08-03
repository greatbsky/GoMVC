{{if eq .tag "doaction"}}
    {{template "doaction" .}}
{{end}}
{{/* doaction template */}}
{{define "doaction"}}
    {{if .isStart}}
    <script type="text/javascript">
        <!--
        $(function(){
            $("html").bind("datagrid-onselect-event", {{.attrs.id4code}}StateHandler);
            $("html").bind("datagrid-onunselect-event", {{.attrs.id4code}}StateHandler);
            $("html").bind("datagrid-onloadsuccess-event", {{.attrs.id4code}}StateHandler);
        });
        function {{.attrs.id4code}}StateHandler(event, j) {
        {{if or (not .attrs.needselection) (contains .attrs.needselection `true`)}}
            var list = getSelections(j.dg);
            if (list.length == 0) {
                $("#{{.attrs.id4code}}").linkbutton("disable");
            } else {
                $("#{{.attrs.id4code}}").linkbutton("enable");
            }
        {{end}}
        }
        -->
    </script>
        <a {{.attrFragment}} id="{{.attrs.id}}" href="" class="easyui-linkbutton"
                             {{if .attrs.disable}}disabled="true"{{end}} data-options="{{.attrs.options}}"
        {{if .attrs.onclick}}
                             onclick="{{.attrs.onclick}}"
        {{else}}
                             onclick="doactionHandler(this, '{{.attrs.action}}', {{if .attrs.confirm}}true{{else}}false{{end}}, {{if .attrs.needselection}}true{{else}}false{{end}})"
        {{end}}
        >{{.attrs.name}}
    {{else}}
        </a>
    {{end}}
{{end}}
