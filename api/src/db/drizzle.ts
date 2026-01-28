import Database from "better-sqlite3";
import { drizzle } from "drizzle-orm/better-sqlite3";

import { config } from "../config/config";
import * as schema from "./schema";

const sqlite = new Database(config.DATABASE_URL);
sqlite.pragma("journal_mode = WAL");
sqlite.pragma("busy_timeout = 5000");
sqlite.pragma("synchronous = NORMAL");
sqlite.pragma("cache_size = 10000");
sqlite.pragma("foreign_keys = ON");
sqlite.pragma("temp_store = MEMORY");
sqlite.pragma("mmap_size = 268435456"); // 256MBfree
// sqlite.pragma("auto_vacuum = INCREMENTAL");

export const db = drizzle({ client: sqlite, schema, casing: "snake_case" });
