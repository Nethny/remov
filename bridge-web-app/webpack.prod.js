const {merge} = require("webpack-merge");
const common = require("./webpack.common");
const TerserPlugin = require('terser-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = merge(common, {
    mode: 'production',
    devtool: 'source-map',
    module: {
        rules: [
            {
                test: /\.(css)$/,
                use: [MiniCssExtractPlugin.loader,'css-loader']
            }
        ]
    },
    plugins: [new MiniCssExtractPlugin()],
    optimization: {
        minimizer: [
            new TerserPlugin({
                // Use multi-process parallel running to improve the build speed
                // Default number of concurrent runs: os.cpus().length - 1
                parallel: true,
            }),
        ],
    },
});
