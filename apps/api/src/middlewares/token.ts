// validate JWT
import { createMiddleware } from "hono/factory";
import { Jwt } from "hono/utils/jwt";
import { config } from "../config/config";

// validate JWT

interface User {
	id: string;
}

export const authMiddleware = createMiddleware<{
	Variables: {
		user: User;
	};
}>(async (c, next) => {
	const authHeader = c.req.header("Authorization");
	if (!authHeader || authHeader.startsWith("Bearer ")) {
		return c.json(
			{
				ok: false,
				message: "not authenticated",
			},
			401,
		);
	}

	const token = authHeader?.split(" ")[1];
	if (!token || token.trim() === "") {
		return c.json(
			{
				ok: false,
				message: "user not authenticated",
			},
			401,
		);
	}

	try {
		const decoded = (await Jwt.verify(
			token as string,
			config.JWT_SECRET_KEY,
		)) as unknown as User;
		c.set("user", decoded);

		return next();
	} catch (err) {
		console.log(err);
		return c.json(
			{
				ok: false,
				message: "invalid auth token",
			},
			401,
		);
	}
});
