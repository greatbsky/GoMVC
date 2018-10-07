<!DOCTYPE html>
<html>
<head>
    {{template "base/head.tpl" .}}
    {{template "base/right.tpl" .}}
    {{.head}}
</head>
<body class="rightpanel">
    <div id="main" class="column">
        <h4 id="alert"></h4>
        {{.body}}
    </div>
    {{template "base/foot.tpl" .}}
    <div class="hd" id="lockscreen" style="width: 100%; height: 100%; position: absolute; left: 0; top: 0; text-align: center; overflow: hidden">
        <div style="position: absolute; left: 3%; top: 3%;">正在处理中……工作辛苦，喝杯茶，休息片刻 ^_^</div>
        <div style="width: 200%; height: 200%; left: -600px; background: #CCCCCC; position: absolute; opacity: 0.1;"></div>
        <a style="position: absolute; right: 3%; top: 3%; color: red; font-size: 18px; text-decoration: none;" href="" onclick="window.done()">×</a>
    </div>
<script>
    window.doing = function(closable) {
        setTimeout("show(\"#lockscreen\")", 500)
        if (typeof closable == "undefined") {
            closable = true
        }
        if (!closable)  {
            hide("#lockscreen a")
        } else {
            show("#lockscreen a")
        }
    }
    window.done = function() {
        hide("#lockscreen")
    }
</script>
</body>
</html>