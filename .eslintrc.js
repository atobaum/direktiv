module.exports = {
  extends: [
    // By extending from a plugin config, we can get recommended rules without having to add them manually.
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:import/recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react-hooks/recommended",
    // This disables the formatting rules in ESLint that Prettier is going to be responsible for handling.
    // Make sure it's always the last config, so it gets the chance to override other configs.
    "eslint-config-prettier",
  ],
  settings: {
    react: {
      // Tells eslint-plugin-react to automatically detect the version of React to use.
      version: "detect",
    },
    // Tells eslint how to resolve imports
    "import/resolver": {
      node: {
        paths: ["src"],
        extensions: [".js", ".jsx", ".ts", ".tsx"],
      },
    },
  },
  env: {
    browser: true,
  },
  ignorePatterns: [
    "node_modules/",
    "dist/",
    ".eslintrc.js",
    "env.d.ts",
    /**
     * TODO_HOOKS_TESTS
     * The Problem: All hooks tests are currently not working and have some linting issues.
     * We must decide:
     *  - Do we keep the hooks as they are?
     *    - when we migrate to typescript, we could potentially refactor them. There are a lot
     *      of repeating patterns.
     *  - if we want to keep the hooks, do we keep and update the tests
     *    - PROS: they are already written and could be useful
     *    - CONS: they are an impplementation detail and maybe
     *      should not be tested
     *    - if we delete them, we need to test the data fetching layer but on a
     *      much higher level. We might use PlayWright for end to end tests.
     *
     */
    // 🚧 search for TODO_HOOKS_TESTS to find all places that needs some action 🚧
    "src/hooks/*test.js",
  ],
  rules: {
    // PLEASE ALWAYS PROVIDE A REASON FOR DISABLING/OVERWRITING A RULE
    // IT'S HARD TO EVALUATE THIS SECTION AT A LATER POINT IN TIME

    // simpe rule to avoid unnecessary curly braces
    "react/jsx-curly-brace-presence": "error",

    // It's save to import React when using vite
    "react/react-in-jsx-scope": "off",

    // console logs are fine in development, but eslint can help us
    // remember to remove them. console.error and console.warn are
    // allowed for now. They should only be placed in code that should
    // not be reached and can provide a helpful hint to the developer.
    "no-console": ["error", { allow: ["error", "warn"] }],

    // REMOVE WHEN 100 % TYPESCRIPT IS ACHIEVED

    // we will use TypeScript's types for component props instead)
    "react/prop-types": "off",
    // this is the default in typescript and we want to enforce it in JavaScript as well
    "prefer-const": "error",
  },
};
