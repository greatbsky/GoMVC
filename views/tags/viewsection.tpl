{{if .isStart}}
    <section class="hd" id="{{if .attrs.id}}{{.attrs.id}}{{else}}viewsection{{end}}">
    <form id="{{if .attrs.viewformid}}{{.attrs.viewformid}}{{else}}viewform{{end}}">
        <table class="ftable">
{{else}}
        </table>
        <div class="formbtns">
            <a href="" class="easyui-linkbutton" data-options="iconCls:'icon-back'" data-crumbreturn="">返回</a>
        </div>
    </form>
    </section>
{{end}}