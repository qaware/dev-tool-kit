const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');

let sourceDir = path.resolve(__dirname, 'src');
let buildDir = path.resolve(__dirname, 'build');

module.exports = {
    entry: {
        index: path.resolve(sourceDir, 'main.js')
    },
    output: {
        path: buildDir,
        filename: 'main.js'
    },
    optimization: {
        splitChunks: false
    },
    mode: 'production',
    module: {
        rules: [
            {
                test: /\.(eot|svg|ttf|woff|woff2)$/i,
                loader: 'url-loader'
            },
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader']
            }
        ]
    },
    plugins: [
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: path.resolve(sourceDir, 'main.css'),
                    to: path.resolve(buildDir, 'main.css')
                }
            ]
        })
    ]
};
