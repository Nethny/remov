const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const DotEnv = require("dotenv-webpack");

module.exports = {
    devtool: 'eval-source-map',
    entry: {
        app: path.join(__dirname, "src", "index.tsx"),
    },
    output: {
        path: path.resolve(__dirname, "build"),
        filename: "[name].bundle.js",
        publicPath: "auto",
        clean: true,
    },
    resolve: {
        extensions: [".tsx", ".ts", ".js", ".css", ".scss"],
        alias: {
            process: "process/browser",
        },
        fallback: {
            "crypto": require.resolve("crypto-browserify"),
            "stream": require.resolve("stream-browserify")
        }
    },
    module: {
        rules: [
            // Vanilla JS
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ["babel-loader"],
            },
            // Typescript
            {
                test: /\.(ts|tsx)$/,
                exclude: /node_modules/,
                use: ["ts-loader"],
            },
            // Images
            {
                test: /\.(?:ico|gif|png|jpg|jpeg)$/i,
                exclude: /node_modules/,
                type: "asset/resource"
            },
            // Fonts and SVGs
            {
                test: /\.(woff(2)?|eot|ttf|otf|svg|)$/,
                exclude: /node_modules/,
                type: "asset/inline",
                use: "svgo-loader"
            },
            // HTML
            {
                test: /\.html$/,
                loader: 'html-loader'
            }
        ],
    },

    // https://webpack.js.org/guides/code-splitting/#splitchunksplugin
    optimization: {
        splitChunks: {
            chunks: 'all',
        },
    },
    performance: {
        maxEntrypointSize: 512000,
        maxAssetSize: 512000
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: path.join(__dirname, "public", "index.html"),
        }),
    ],
}