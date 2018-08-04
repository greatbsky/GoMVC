{{if .isStart}}
    <form class="searchform" method="get" {{.attrFragment}} {{if not .attrs.action}}action="{{.baseapi}}/api/{{.channel}}/{{.table}}/list"{{end}}>
    <table class="normal">
{{else}}
    </table>
        <div class="formbtns">
            <a href="" class="easyui-linkbutton" data-options="iconCls:'icon-search'" onclick="searchHandler(this, '{{.attrs.datagrid}}')">查询</a>
            <a href="" class="easyui-linkbutton" data-options="iconCls:'icon-undo'" onclick="resetForm(this)">重置</a>
        </div>
    </form>
    <script type="text/javascript">
        <!--
        $(":input[class='search']").keydown(function(){
            if(event.keyCode==13){
                $(".searchform .formbtns").children("a").eq(0).click();
                return false;
            }
        });
        -->
    </script>
{{end}}