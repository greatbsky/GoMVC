<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    <style type="text/css" media="screen">
        tr {line-height:30px;}
        .mlfttop60 {margin-left:60px; margin-top: 60px;}
    </style>
    <script type="text/javascript">
        <!--
        -->
    </script>
</head>
<body>
<section>
    <div class="easyui-panel mlfttop60" title="当前管理员信息" style="width:300px;">
        <table cellpadding="6">
            <tr>
                <td>姓名:</td>
                <td>${(admin.name)!}</td>
            </tr>
            <tr>
                <td>登录ID:</td>
                <td>${(admin.adminid)!}</td>
            </tr>
            <tr>
                <td>手机号码:</td>
                <td>${(admin.mobile)!}</td>
            </tr>
            <tr>
                <td>电子邮件:</td>
                <td>${(admin.email)!}</td>
            </tr>
        </table>
    </div>
</section>
</body>
</html>