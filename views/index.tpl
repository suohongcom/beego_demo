

    <div class="layui-form" action="" style="margin: 5px; border: 1px silver">
        <form class="layui-form" action="/"  method="get"  style="margin-top:30px">
            <div class="layui-inline">
                <label class="layui-form-label">关键字</label>
                <div class="layui-input-block">
                    <input type="text" name="keyword" placeholder="请输入关键字" autocomplete="on" value="{{.keyword}}" class="layui-input">
                </div>
            </div>

            <div class=" layui-inline">
                <button class="layui-btn layui-btn-normal  ">
                    搜索
                </button>
            </div>
        </form>

        <div class=" layui-inline" style="float: right">
            <button class="layui-btn " onclick="add()">
                <i class="layui-icon">&#xe608;</i> 添加
            </button>
        </div>

    </div>

    <table class="layui-table" lay-skin="row">
        <thead>
                <th>序号</th>
                <th>姓名</th>
                <th>金额</th>
                <td>状态</td>
                <td>地址</td>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
        {{range .list}}
            <tr>
                <td>{{.Id}}</td>
                <td>{{.Username}}</td>
                <td>{{.Money}}</td>
                <td>{{.Status}}</td>
                <td>{{.Address}}</td>
                <td>
                    <a href="/create?id={{.Id}}" >
                        <i class="layui-icon">&#xe642;</i> 修改
                    </a>
                    <a href="javascript:void(0)" onclick="del({{.Id}})"  >
                        <i class="layui-icon">&#x1006;</i> 删除
                    </a>
                </td>
            </tr>
        {{end}}
        </tbody>
    </table>

    <div class="layui-box layui-laypage layui-laypage-default" id="layui-laypage-0">
    {{str2html .pagebar}}

    </div>


<script type="text/javascript">
    /**
     * 对layui进行全局配置
     */
    layui.config({
        base: '/static/js/'
    });

    var add = function () {
        location.href = "/create?id=0"
    }

    layui.use('form', function () {
        var $ = layui.jquery,
            form = layui.form();
    });

    layui.use('element', function () {
    });
</script>