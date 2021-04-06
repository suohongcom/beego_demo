<div class="main layui-clear">
    <form action="/login" method="post">
        <div class="fly-panel fly-panel-user" pad20>
            <div class="layui-tab layui-tab-brief">
                <ul class="layui-tab-title">
                    <li class="layui-this">欢迎登录后台系统</li>
                </ul>
                <div class="layui-form layui-tab-content" id="LAY_ucm" style="padding-top: 35px; padding-left: 45px;  ">
                    <div class="layui-form layui-form-pane">
                        <div class="layui-form-item">
                            <label class="layui-form-label">用户名</label>
                            <div class="layui-input-inline">
                                <input type="text" name="username" required lay-verify="username" placeholder="用户名"
                                       autocomplete="off" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">密码</label>
                            <div class="layui-input-inline">
                                <input type="password" name="password" required lay-verify="password"
                                       placeholder="密码" autocomplete="off" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">验证码</label>
                            <div class="layui-input-inline" style="width: 95px;">
                                <input type="text" name="captcha" required lay-verify="captcha"
                                       placeholder="密码" autocomplete="off" class="layui-input">               
                            </div>
                            <div class="layui-word-aux">{{create_captcha}}</div>
                        </div>
                        <div class="layui-form-item" style="float: right; margin-right: 42px;">
                            <input type="checkbox" name="is_top" {{if .post.IsTop}} checked {{end}} value="1"
                                   title="记住密码">

                        </div>
                        <div class="layui-form-item">
                            <button lay-submit class="layui-btn btn-submit" style="width: 300px; border-radius:3px"
                                    lay-submit=""
                                    lay-filter="sub">立即登录
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </form>
</div>