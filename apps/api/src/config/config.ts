import "dotenv/config";
import * as z from "zod/v4";

const configSchema = z.object({
	PORT: z.coerce.number({ error: "PORT is missing" }).default(5000),
	DATABASE_URL: z
		.string({ error: "DATABASE_URL is missing" })
		.trim()
		.default("./microblog.db"),
	JWT_SECRET_KEY: z.string({ error: "JWT_SECRET_KEY is missing" }).nonempty(),
});

function validateConfig() {
	const { success, data, error } = configSchema.safeParse(process.env);
	if (!success) {
		const errors = Object.values(z.flattenError(error).fieldErrors).flat();
		throw Error(errors.at(0));
	}

	return data;
}

export const config = validateConfig();
