module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './monitrix/tsconfig.json',
    tsconfigRootDir: __dirname,
    sourceType: 'module',
  },
  plugins: ['@typescript-eslint'],
  extends: ['plugin:@typescript-eslint/recommended'],
  rules: {
    // optional rules here
  },
};
