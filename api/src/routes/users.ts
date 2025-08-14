// generate route to add new user in honojs
import bcrypt from "bcryptjs";
import { eq } from "drizzle-orm";
import { Hono } from "hono";
import * as jwt from "hono/jwt";
import { z } from "zod";

import { config } from "../config/config";
import { db } from "../db/drizzle";
import { usersTable } from "../db/schema";

const app = new Hono().basePath("/users");

const createUserSchema = z.object({
  username: z.string().min(2),
  email: z.email(),
  password: z.string().min(6),
});

app.post("/", async (c) => {
  const body = await c.req.json();

  // Check if user already exists
  const existingUser = await db.query.usersTable.findFirst({
    where: eq(usersTable.username, body.username),
  });

  if (existingUser?.username || existingUser?.email) {
    return c.json({ error: "username or email already exists" }, 400);
  }

  try {
    // Hash password
    const hashedPassword = await bcrypt.hash(body.password, 12);

    // Create user
    const newUser = await db
      .insert(usersTable)
      .values({
        username: body.username,
        email: body.email,
        password: hashedPassword,
      })
      .returning();

    // Generate token for automatic login
    const token = await jwt.sign(
      { id: newUser.id, exp: 24 * 60 * 60 * 1 },
      config.JWT_SECRET_KEY,
    );

    return c.json({ data: { token } }, 201);
  } catch (_e: unknown) {
    console.log(_e);
    return c.json({ error: "Failed to create user" }, 500);
  }
});

export default app;
