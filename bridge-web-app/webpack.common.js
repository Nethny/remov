const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const { ProvidePlugin } = require("webpack");
// const { ModuleFederationPlugin } = require("webpack").container;
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
            // Styles and style modules
            {
                test: /\.(scss)$/,
                use: [{
                    // inject CSS to page
                    loader: 'style-loader'
                }, {
                    // translates CSS into CommonJS modules
                    loader: 'css-loader'
                }, {
                    // Run postcss actions
                    loader: 'postcss-loader',
                    options: {
                        // `postcssOptions` is needed for postcss 8.x;
                        // if you use postcss 7.x skip the key
                        postcssOptions: {
                            // postcss plugins, can be exported to postcss.config.js
                            plugins: function () {
                                return [
                                    require('autoprefixer')
                                ];
                            }
                        }
                    }
                }, {
                    // compiles Sass to CSS
                    loader: 'sass-loader'
                },
                ]
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
        /*
        new ProvidePlugin({
            process: "process/browser",
        }),
        new DotEnv({
            path: "./config/env/.env", // load this now instead of the ones in '.env'
            safe: false, // load '.env.example' to verify the '.env' variables are all set. Can also be a string to a different file.
            systemvars: false, // load all the predefined 'process.env' variables which will trump anything local per dotenv specs.
            silent: false, // hide any errors
            defaults: false, // load '.env.defaults' as the default values if empty.
        }),*/
    ],
}