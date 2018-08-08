<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    {{template "base/head.tpl" .}}
    <link rel="stylesheet" type="text/css" href="{{.basecss}}/css/admin/layout.css" media="screen" />
    <link rel="stylesheet" type="text/css" href="{{.basecss}}/css/admin/right.css" media="screen" />
    <style type="text/css" media="screen">
        .breadcrumbs_container {position:relative;}
        #messagebar {position:absolute; right:30px; margin:4px 0; border-radius:15px; padding:3px 0; padding-right:9px; background-position:9px 3px;}
    </style>
    <script type="text/javascript">
        <!--
        -->
    </script>
</head>
<body>
<header id="header">
    <hgroup>
        <h1 class="site_title"><a href="#">{{.title}}</a></h1>
        <h2 class="section_title"></h2><div class="btn_view_site"><a href="{{.base}}/admin/index" target="_top">主页</a></div>
    </hgroup>
</header>
<section id="secondary_bar">
    <div class="user">
        <p>超级管理员</p>
        <a class="logout_user" href="{{.base}}/login" title="退出">退出</a>
    </div>
    <div class="breadcrumbs_container">
        <article class="breadcrumbs"><a href="">Website Admin</a> <div class="breadcrumb_divider"></div> <a class="current">Dashboard</a></article>
        <div id="messagebar"></div>
    </div>
</section>
{{template "base/foot.tpl" .}}
</body>
</html>