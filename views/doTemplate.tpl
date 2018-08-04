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
</body>
</html>