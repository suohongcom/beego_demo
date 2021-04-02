<form class="layui-form" action="/create"  method="post" style="margin:20px">

    <div class="layui-form-item">
        <label class="layui-form-label">姓名：</label>
        <div class="layui-input-block">
            <input type="hidden" name="id" value="{{.post.id}}" > 
            <input type="text" name="username" required value="{{.post.username}}" lay-verify="required" placeholder="请输入姓名" autocomplete="off" class="layui-input">
        </div>
    </div>

    <div class="layui-form-item">
        <label class="layui-form-label">金额：</label>
        <div class="layui-input-block">
            <input type="text" name="money" required value="{{.post.money}}" lay-verify="required" placeholder="请输入金额" autocomplete="off" class="layui-input">
        </div>
    </div>

    <div class="layui-form-item">
        <label class="layui-form-label">地址：</label>
        <div class="layui-input-block">
            <input type="text" name="address" required value="{{.post.address}}" lay-verify="required" placeholder="请输入地址" autocomplete="off" class="layui-input">
        </div>
    </div>

    <div class="layui-form-item">
        <div class="layui-input-block">
            <button class="layui-btn" lay-submit lay-filter="formDemo">提交</button>
            <button class="layui-btn layui-btn-primary" onclick="history.go(-1)">返回</button>
        </div>
    </div>
</form>
<script>
    layui.use(['element','form'], function () {
    });
</script>
