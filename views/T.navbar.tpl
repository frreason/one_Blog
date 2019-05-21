{{define "navbar"}}
<!-- 导航栏 -->
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div>
            <a class="navbar-brand" href="#">One Blog</a>
            <ul class="nav navbar-nav">
                <li {{if .IsHome}}class="active" {{end}}><a href="/">首页</a></li>
                <li {{if .IsTopic}}class="active" {{end}}><a href="/topic">文章</a></li>
                <li class="dropdown" {{if .IsCategory}}class="active" {{end}}>
                    <a id="drop1" href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true"
                        aria-expanded="false">
                        分类
                        <span class="caret"></span>
                    </a>
                    <ul class="dropdown-menu" aria-labelledby="drop1">
                        <li><a href="#">C/C++</a></li>
                        <li><a href="#">Go</a></li>
                        <li><a href="#">Python</a></li>
                        <li role="separator" class="divider"></li>
                        <li><a href="/category">所有分类</a></li>
                    </ul>
                </li>
            </ul>
        </div>
        <div class="pull-right">
            <ul class="nav navbar-nav">
                {{if .IsLogin}}
                <li><a href="/login?exit=true">退出</a></li>
                {{else}}
                <li><a href="/login">管理员登陆</a></li>
                {{end}}
            </ul>
        </div>
    </div>


</div>
{{end}}