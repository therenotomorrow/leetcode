import js from '@eslint/js'
import prettier from 'eslint-plugin-prettier'
import prettierConfig from 'eslint-config-prettier'
import jest from 'eslint-plugin-jest'

export default [
  js.configs.recommended,
  {
    files: ['**/*.js'],
    languageOptions: {
      sourceType: 'module',
      ecmaVersion: 'latest',
      globals: {
        describe: 'readonly',
        test: 'readonly',
        expect: 'readonly',
      },
    },
    plugins: {
      prettier,
      jest,
    },
    rules: {
      'prettier/prettier': 'error',
      eqeqeq: ['error', 'always'],
      'no-console': 'warn',
      'no-debugger': 'error',
    },
  },
  prettierConfig,
]
