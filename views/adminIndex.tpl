<!DOCTYPE html>
<html>
<head>
    <title>后台管理中心——{{.title}}</title>
    {{template "base/head.tpl" .}}
    <script type="text/javascript">
        <!--
        forbiddenIframe();
        -->
    </script>
</head>
<frameset rows="90,*" border="0" framespacing="0" frameborder="0">
    <frame id="topbar" name="topbar" src="{{.base}}/admin/top" marginheight="0" marginwidth="0" noresize="noresize" scrolling="no" />
    <frameset cols="23%,*" border="0" framespacing="0" frameborder="0">
        <frame id="leftmenu" name="leftmenu" src="{{.base}}/admin/menu" marginheight="0" marginwidth="0" noresize="noresize" scrolling="no" />
        <frame id="rightpanel" name="rightpanel" src="{{.base}}/admin/home" marginheight="0" marginwidth="0" noresize="noresize" scrolling="auto" />
        <noframes>对不起，您的浏览器不支持框架元素，请升级!</noframes>
    </frameset>
    <noframes>对不起，您的浏览器不支持框架元素，请升级!</noframes>
</frameset>
</html>
{{template "base/foot.tpl" .}}