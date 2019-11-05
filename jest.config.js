module.exports = {
  preset: "ts-jest",
  verbose: true,
  roots: ["<rootDir>/ui/src"],
  transform: {
    "^.+\\.tsx?$": "ts-jest"
  },
  setupFilesAfterEnv: ["<rootDir>/ui/src/setupTests.ts"],
  globals: {
    "ts-jest": {
      tsConfig: "ui/tsconfig.json"
    }
  }
};
