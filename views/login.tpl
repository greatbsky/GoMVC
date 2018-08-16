<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
{{template "base/head.tpl" .}}
    <link rel="stylesheet" type="text/css" href="{{.basecss}}/css/admin/pages/login.css" media="screen" />
    <script type="text/javascript">

        (function logout() {
            $.get("{{.base}}/api/admin/logout");
            $.get("{{.baseapi}}{{.apiLogout}}");
        })();
        $(function() {
            var userAgent = navigator.userAgent; //取得浏览器的userAgent字符串
            var isOpera = userAgent.indexOf("Opera") > -1; //判断是否Opera浏览器
            var isIE = userAgent.indexOf("Trident/7.0") > -1 || userAgent.indexOf("MSIE") > -1 && !isOpera; //判断是否IE浏览器
            var isEdge = userAgent.indexOf("Edge") > -1 && !isIE; //判断是否IE的Edge浏览器
            var isFF = userAgent.indexOf("Firefox") > -1; //判断是否Firefox浏览器
            var isSafari = userAgent.indexOf("Safari") > -1 && userAgent.indexOf("Chrome") == -1; //判断是否Safari浏览器
            var isChrome = userAgent.indexOf("Chrome") > -1 && userAgent.indexOf("Safari") > -1 && !isEdge; //判断Chrome浏览器
            if (isFF || isChrome) {
                return;
            } else {
                alert("本管理后台仅适用于chrome、firefox浏览器，使用其他浏览器访问时可能会导致某些功能不能正常使用。");
            }
        });
        forbiddenIframe();
    </script>
</head>
<body bgproperties="fixed">
<div class="container">
    <div class="topbar">
        <a href="">
            <strong></strong>http://www.{{.domain}}
        </a>
        <span class="right">
	        <a href="">
	            <strong></strong>
	        </a>
	    </span>
        <div class="clr"></div>
    </div>
    <header>
        <br /><br />
        <h1>Please Login <span>Admin Web Console</span></h1>
        <h2><br /></h2>
    </header>
    <section>
        <div class="content">
            <div id="wrapper">
                <div id="login" class="animate form easyui-panel">
                    <form id="f_login" action="" method="post">
                        <h1>Log in</h1>
                        <p>
                            <label for="txtid" class="uname" data-icon="u">管理员ID</label>
                            <input id="txtid" name="id" value="" required="required" placeholder="请输入管理员ID" />
                        </p>
                        <p>
                            <label for="password" class="youpasswd" data-icon="p">密码</label>
                            <input id="txtpwd" name="password" type="password" required="required"  placeholder="请输入密码" />
                        </p>
                        <p class="login button">
                            <span id="errmsg" class="red left"></span><input type="button" value="登录" onclick="submitLoginForm()"  />
                        </p>
                    </form>
                    <script>
                        function submitLoginForm() {
                            $.post("{{if .apiLogin}}{{.apiLogin}}{{else}}{{.baseapi}}/api/admin/login{{end}}", $("#f_login").serialize()).done(function(data) {
                                data = parseResult(data)
                                if (data.success) {
                                    Cookies.set('adminuid', data.adminuid);
                                    window.location.assign("{{.base}}/admin/index")
                                } else {
                                    console.log(data);
                                    $("#errmsg").text(data.msg)
                                }
                            }).fail(function(xhr, status, error) {
                                console.log(data);
                                $("#errmsg").text("服务端返回结果不正确")
                            })
                        }
                    </script>
                </div>
            </div>
        </div>
    </section>
</div>
{{template "base/foot.tpl" .}}
</body>
</html>