// vue.config.js
module.exports = {
    // 选项...
    devServer: {
        disableHostCheck: true,
        proxy: {
            '/go/': {
                target: 'http://localhost:8080',
                secure: false,
                changeOrigin: true,
                pathRewrite: {
                    "^/go/": "/"
                }
            }
        }

    },
    //配置本地与现实build环境
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
};

