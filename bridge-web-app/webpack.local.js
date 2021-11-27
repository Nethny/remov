const { merge } = require("webpack-merge");
const common = require("./webpack.common");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const DotEnv = require("dotenv-webpack");

module.exports = merge(common, {
    mode: "development",
    devtool: "inline-source-map",
    devServer: {
        historyApiFallback: true,
        hot: true,
        port: 8088,
    },
    plugins: [
        new MiniCssExtractPlugin(),
        new DotEnv({
            path: `./.env`,
            safe: false, // load '.env.example' to verify the '.env' variables are all set. Can also be a string to a different file.
            systemvars: false, // load all the predefined 'process.env' variables which will trump anything local per dotenv specs.
            silent: true, // hide any errors
            defaults: false, // load '.env.defaults' as the default values if empty.
        }),
    ],
});