import { defineConfig } from "drizzle-kit";

import { config } from "./src/config/config";

export default defineConfig({
  out: "./drizle",
  dialect: "sqlite",
  schema: "./src/db/schema.ts",
  breakpoints: true,
  casing: "snake_case",
  strict: true,
  verbose: true,
  dbCredentials: {
    // remove string class when zod validation is available
    url: String(config.DATABASE_URL),
  },
});
