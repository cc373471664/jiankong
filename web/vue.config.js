// vue.config.js
module.exports = {
    // 选项...
    devServer: {
        disableHostCheck: true,
        proxy: {
            '/go/': {
                target: 'http://localhost:8071',
                secure: false,
                ws: true, // 是否启用websockets
                changeOrigin: true,
                pathRewrite: {
                    "^/go/": "/"
                }
            },
            '/socket': {
                target: 'ws://localhost:8071',//后端目标接口地址
                changeOrigin: true,//是否允许跨域
                pathRewrite: {
                    '^/socket': '',//重写,
                },
                ws: true //开启ws, 如果是http代理此处可以不用设置
            }
        }

    },
    //配置本地与现实build环境
    publicPath: process.env.NODE_ENV === 'production' ? './' : '/',
};

