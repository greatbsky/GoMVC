{{if .isStart}}
    {{if or (not .attrs.candisable) (contains .attrs.candisable `true`)}}
    <a id="btndisable" href="" class="easyui-linkbutton" data-options=""
            onclick="statusHandler(this, '{{if .attrs.urldisable}}{{.attrs.urldisable}}{{else}}{{.baseapi}}/api/{{.channel}}/{{.table}}/status/0{{end}}')"
    >{{if .attrs.namedisable}}{{.attrs.namedisable}}{{else}}冻结{{end}}</a>
    {{end}}
    {{if or (not .attrs.canenable) (contains .attrs.canenable `true`)}}
    <a id="btnenable" href="" class="easyui-linkbutton" data-options=""
       onclick="statusHandler(this, '{{if .attrs.urlenable}}{{.attrs.urlenable}}{{else}}{{.baseapi}}/api/{{.channel}}/{{.table}}/status/1{{end}}')"
    >{{if .attrs.nameenable}}{{.attrs.nameenable}}{{else}}解冻{{end}}</a>
    {{end}}
    <div id="dlg-status" class="easyui-dialog" style="{{if .attrs.style}}{{.attrs.style}}{{else}}width:500px; height:240px; padding:10px 20px;{{end}}" title="{0}选中的项" buttons="#dlg-status-btns" data-options="closed:true">
    <form id="statusform" method="post" action="#" onkeydown="if(event.keyCode==13){return false;}">
        <input type="hidden" set-key="uid" set-key2="oid" set-key3="id" name="ids" />
        <input type="hidden" name="adminuid" set-value="{{.adminuid}}" />
        <input type="hidden" set-key="status" name="statusfrom" />
    <table class="ftable">
{{else}}
    <tr>
        <td>操作密码</td>
        <td><input class="easyui-validatebox" type="password" set-value="" name="dopwd" data-options="required:true" /></td>
    </tr>
    <tr>
        <td>备注</td>
        <td><textarea class="easyui-validatebox" style="width:350px; height:60px;" set-value="" name="desc" data-options="required:true" ></textarea></td>
    </tr>
</table>
</form>
    <div id="dlg-status-btns">
        <a href="" class="easyui-linkbutton" iconCls="icon-ok" onclick="status_submitHandler(this)">确定</a>
        <a href="" class="easyui-linkbutton" iconCls="icon-cancel" onclick="statusdlg_cancelHandler(this)">取消</a>
    </div>
</div>
<script type="text/javascript">
    <!--
    $("html").bind("status-reset-event", function(){
        $("#statusform").form("clear");
    });
    //-->
</script>
{{end}}
