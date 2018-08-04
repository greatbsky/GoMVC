{{if .isStart}}
    <section id="{{if .attrs.id}}{{.attrs.id}}{{else}}editsection{{end}}" class="hd">
    <form id="{{if .attrs.editformid}}{{.attrs.editformid}}{{else}}editform{{end}}" method="post"
          {{if not .attrs.action}}action="{{.baseapi}}/api/{{.channel}}/{{.table}}/edit"{{end}}
          enctype="multipart/form-data">
    <table class="ftable">
{{else}}
    </table>
        <div class="formbtns">
            <a href="" class="easyui-linkbutton" data-options="iconCls:'icon-ok'" onclick="submitHandler(this)">提交</a>
            <a href="" class="easyui-linkbutton" data-options="iconCls:'icon-back'" data-crumbreturn="">返回</a>
        </div>
    </form>
    </section>
{{end}}