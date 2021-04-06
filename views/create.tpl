<form class="layui-form" action="/create"  method="post" style="margin:20px">

    <div class="layui-form-item">
        <label class="layui-form-label">姓名：</label>
        <div class="layui-input-block">
            <input type="hidden" name="id" value="{{.money.Id}}" > 
            <input type="text" name="username" required value="{{.money.Username}}" lay-verify="required" placeholder="请输入姓名" autocomplete="off" class="layui-input">
        </div>
    </div>

    <div class="layui-form-item">
        <label class="layui-form-label">金额：</label>
        <div class="layui-input-block">
            <input type="text" name="money" required value="{{.money.Money}}" lay-verify="required" placeholder="请输入金额" autocomplete="off" class="layui-input">
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">单选框</label>
        <div class="layui-input-block">
            <input type="radio" name="status" value="1" title="启用" {{if .money.Status}} checked {{end}}>
            <input type="radio" name="status" value="0" title="禁用" {{if .money.Status}} {{else}} checked {{end}}>
        </div>
    </div>
    <div class="layui-form-item">
        <label class="layui-form-label">地址：</label>
        <div class="layui-input-block">
            <input type="text" name="address" required value="{{.money.Address}}" lay-verify="required" placeholder="请输入地址" autocomplete="off" class="layui-input">
        </div>
    </div>

    <div class="layui-form-item">
        <div class="layui-input-block">
            <button class="layui-btn" lay-submit lay-filter="formDemo">提交</button>
            <a class="layui-btn layui-btn-primary" href="/">返回</a>
        </div>
    </div>
</form>
<script>
    layui.use(['form'],function(){
        var form = layui.form;
        
        //刷新界面 所有元素
    
        form.render();
    
    });
</script>
