

    <div class="layui-form" action="" style="margin: 5px; border: 1px silver">

        <div class="layui-inline">
            <label class="layui-form-label">关键字</label>
            <div class="layui-input-block">
                <input type="text" name="title" placeholder="请输入关键字" autocomplete="on" class="layui-input">
            </div>
        </div>

        <div class=" layui-inline">
            <button class="layui-btn layui-btn-normal  ">
                搜索
            </button>
        </div>

        <div class=" layui-inline" style="float: right">
            <button class="layui-btn " onclick="add()">
                <i class="layui-icon">&#xe608;</i> 添加
            </button>
        </div>

    </div>

    <table class="layui-table" lay-skin="row">
        <thead>
            <tr>
                <th>标题</th>
                <th>时间</th>
                <td>置顶</td>
                <th>点击量</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>

        </tbody>
    </table>

    <div class="layui-box layui-laypage layui-laypage-default" id="layui-laypage-0">

    </div>


<script type="text/javascript">
    /**
     * 对layui进行全局配置
     */
    layui.config({
        base: '/static/js/'
    });

    var add = function () {
        location.href = "/create"
    }

    layui.use('form', function () {
        var $ = layui.jquery,
            form = layui.form();
    });

    layui.use('element', function () {
    });
</script>