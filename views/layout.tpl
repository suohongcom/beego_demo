<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>后台登录</title>
    <link rel="stylesheet" href="/static/plug/layui/css/layui.css">
    <style>
        .main {
            margin: 0 auto;
            width: 400px;
            border: 1px solid;
            border-color: #eeeeee;
            border-radius: 5px;
            margin-top: 100px;
        }
    </style>
</head>
<script type="text/javascript" src="/static/plug/layui/layui.js"></script>

<body ontouchstart>
    {{if ne "login" .Header }}
    <!-- 布局容器 -->
    <div class="layui-layout layui-layout-admin">
        <!-- 头部 -->
        <div class="layui-header">
            <div class="layui-main">
                <!-- logo -->
                <a href="/" style="color: #c2c2c2; font-size: 18px; line-height: 60px;">后台管理系统</a>
                <!-- 水平导航 -->
                <ul class="layui-nav" style="position: absolute; top: 0; right: 0; background: none;">
                    <li class="layui-nav-item">
                        <a href="javascript:;">

                        </a>
                    </li>
                    <li class="layui-nav-item">
                        <a href="javascript:;">
                            admin,欢迎你！
                        </a>
                        <dl class="layui-nav-child">
                            <dd>
                                <a href="/logout">
                                    退出系统
                                </a>
                            </dd>
                        </dl>
                    </li>
                </ul>
            </div>
        </div>
    </div>
    
    {{end}}

    {{.LayoutContent}}
</body>
<script type="text/javascript" src="/static/plug/layui/lay/lib/jquery.js"></script>
<script type="text/javascript" src="/static/plug/layui/layui.js"></script>

</html>