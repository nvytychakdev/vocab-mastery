// @ts-check
const eslint = require('@eslint/js');
const tseslint = require('typescript-eslint');
const angular = require('angular-eslint');
const importPlugin = require('eslint-plugin-import');

module.exports = tseslint.config(
  {
    files: ['**/*.ts'],
    extends: [
      eslint.configs.recommended,
      ...tseslint.configs.recommended,
      ...tseslint.configs.stylistic,
      ...angular.configs.tsRecommended,
      // importPlugin.flatConfigs.recommended,
      // importPlugin.flatConfigs.typescript,
    ],
    processor: angular.processInlineTemplates,
    plugins: {
      import: importPlugin,
    },
    rules: {
      '@typescript-eslint/consistent-type-definitions': ['error', 'type'],
      '@angular-eslint/directive-selector': [
        'error',
        {
          type: 'attribute',
          prefix: 'app',
          style: 'camelCase',
        },
      ],
      '@angular-eslint/component-selector': [
        'error',
        {
          type: 'element',
          prefix: 'app',
          style: 'kebab-case',
        },
      ],
    },
  },
  {
    files: ['**/*.html'],
    extends: [...angular.configs.templateRecommended, ...angular.configs.templateAccessibility],
    rules: {},
  },

  // restriction rules for import
  {
    files: ['src/app/core/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@domain/*', '@feature/*'] }],
    },
  },
  {
    files: ['src/app/domaints/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@feature/*'] }],
    },
  },
  {
    files: ['src/app/domains/auth/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@domain/*'] }],
    },
  },
  {
    files: ['src/app/domains/dictionary/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@domain/word/*', '@domain/translation/*'] }],
    },
  },
  {
    files: ['src/app/domains/word/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@domain/dictionary/*'] }],
    },
  },
  {
    files: ['src/app/domains/translation/**/*.ts'],
    rules: {
      'no-restricted-imports': ['error', { patterns: ['@domain/*'] }],
    },
  }
);
