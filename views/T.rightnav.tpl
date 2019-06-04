{{define "rightnav"}}

<div class="col-md-3">
    <h3>文章分类</h3>
    <ul class="nav nav-tabs nav-stacked">
        {{range .Categroies}}
        <li>
            <a href='http://127.0.0.1:8658/topic/category/{{.Id}}' style="color: black">
                {{.Title}}
                <span class="badge pull-right">{{.Total}}</span>
            </a> 
        </li>
        {{end}}
    </ul>
    <h3>浏览最多</h3>
    <ul class="nav nav-tabs nav-stacked">
        {{range .ViewsMaxTopic}}
        <li>
            <a href='http://127.0.0.1:8658/topic/view/{{.Id}}' style="color: black">
                {{.Title}}
                <span class="text-muted pull-right" style="font-size: 14px;">浏览数：{{.Views}}</span>
            </a>
        </li>
        {{end}}
    </ul>
    <h3>最新评论</h3>
    <ul class="nav nav-tabs nav-stacked">
        {{range .LastestComments}}
        <li>
            <a href='http://127.0.0.1:8658/topic/view/{{.Tid}}' style="color: black">
                {{.Writer}}: {{.Content}}
            </a>
        </li>
        {{end}}
    </ul>
</div>
</div>


{{end}}