module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  verbose: true,
  roots: ["<rootDir>/ui/src"],
  transform: {
    "^.+\\.tsx?$": "ts-jest"
  }
};
