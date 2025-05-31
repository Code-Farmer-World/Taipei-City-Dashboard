const PORT = process.env.VUE_APP_PORT

module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? '/neihu_traffic/'
    : '/',
  devServer: {
    port: PORT? PORT: 3000
  },
  productionSourceMap: false
}
