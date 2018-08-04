{{if .isStart}}
    <div id="toolbar">
    {{if or (not .attrs.canadd) (contains .attrs.canadd `true`)}}
    <a id="btnadd" href="" class="easyui-linkbutton" data-options="iconCls:'icon-add'"
       data-crumbshow="{selector:'{{if .attrs.addsection}}{{.attrs.addsection}}{{else}}#addsection{{end}}', title:'{{if .attrs.nameadd}}{{.attrs.nameadd}}{{else}}添加{{end}}'}"
       onclick="{{if .attrs.onclickadd}}{{.attrs.onclickadd}}{{else}}addHandler(this){{end}}"
    >{{if .attrs.nameadd}}{{.attrs.nameadd}}{{else}}添加{{end}}</a>
    {{end}}
    {{if (or (not .attrs.canedit) (contains .attrs.canedit "true"))}}
    <a id="btnedit" href="" class="easyui-linkbutton" data-options="iconCls:'icon-edit'"
       data-crumbshow="{selector:'{{if .attrs.editsection}}{{.attrs.editsection}}{{else}}#editsection{{end}}', title:'{{if .attrs.titleedit}}{{.attrs.titleedit}}{{else}}编辑{{end}}'}"
       onclick="{{if .attrs.onclickedit}}{{.attrs.onclickedit}}{{else}}editHandler(this){{end}}"
    >{{if .attrs.titleedit}}{{.attrs.titleedit}}{{else}}编辑{{end}}</a>
    {{end}}
    {{if (or (not .attrs.candel) (contains .attrs.candel "true"))}}
    <a id="btndel" href="" class="easyui-linkbutton" data-options="iconCls:'icon-remove'"
       onclick="{{if .attrs.onclickdel}}{{.attrs.onclickdel}}{{else}}deleteHandler('{{if .attrs.urldel}}{{.attrs.urldel}}{{else}}{{.baseapi}}/api/{{.channel}}/{{.table}}/delete{{end}}'){{end}}"
    >{{if .attrs.titledel}}{{.attrs.titledel}}{{else}}删除{{end}}</a>
    {{end}}
    {{if (or (not .attrs.canview) (contains .attrs.canview `true`))}}
    <a id="btnview" href="" class="easyui-linkbutton" data-options="iconCls:'icon-tip'"
       data-crumbshow="{selector:'{{if .attrs.viewsection}}{{.attrs.viewsection}}{{else}}#viewsection{{end}}', title:'{{if .attrs.titleview}}{{.attrs.titleview}}{{else}}查看详情{{end}}'}"
       onclick="{{if .attrs.onclickview}}{{.attrs.onclickview}}{{else}}viewHandler(this){{end}}"
    >{{if .attrs.titleview}}{{.attrs.titleview}}{{else}}查看详情{{end}}</a>
    {{end}}
{{else}}
    </div>
    <script type="text/javascript">
        <!--
        var topTo = 0;
        $(function() {
            topTo = $("#toolbar").offset().top;
        });
        $(window).scroll(function() {
            var top = $(this).scrollTop();
            if(top > topTo){
                $("#toolbar").css("position", "fixed");
                $("#toolbar").css("z-index", 1);
                $("#toolbar").css("top", 0);
            }else{
                $("#toolbar").css("position", "static");
                $("#toolbar").css("top", topTo);
            }
        });
        -->
    </script>
{{end}}