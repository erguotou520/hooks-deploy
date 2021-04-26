// eslint-disable-next-line import/no-commonjs
module.exports = {
  presets: [
    [
      '@babel/preset-env',
      {
        useBuiltIns: 'usage',
        corejs: '3.8'
      }
    ],
    ['@babel/preset-typescript']
  ],
  plugins: [['@babel/plugin-transform-runtime', {}]]
}
