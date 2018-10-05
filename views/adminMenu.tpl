<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    {{template "base/head.tpl" .}}
    <link rel="stylesheet" type="text/css" href="{{.basecss}}/css/admin/layout.css" media="screen" />
    <style type="text/css" media="screen">
        .toggle-icon{width:12px; height:10px; float:right; margin-top:5px; cursor:pointer;
            background:url('{{.baseimg}}/imgs/admin/adminicons.png') no-repeat -54px -55px;
        }
        .open {
            background:url('{{.baseimg}}/imgs/admin/adminicons.png') no-repeat -54px -76px;
        }
    </style>
    <script type="text/javascript">
        <!--
        $(function($) {
            $("#sidebar a .menu-name").append("<span class='toggle-icon open'></span>");
            $("#sidebar a .menu-name").click(function(){
                $(this).parent().next("ul.toggle:first").toggle(200);
                $(this).children().toggleClass("open");
            });
            $(".closed").click();

            $("ul").each(function() {
                if($(this).children().length == 0) {
                    $(this).prev("a").remove();
                }
            });
        });
        -->
    </script>
</head>
<body>
<aside id="sidebar" class="column" style="height: 100%;overflow-y: auto">
{{.body}}
    <footer>
        <hr>
        <p><strong>2018 {{.domain}}, All Right Reserved,Inc.</strong></p>
        <p>it@{{.domain}}</p>
    </footer>
</aside>
{{template "base/foot.tpl" .}}
</body>
</html>