import { serve } from "@hono/node-server";

import app from "./src/app.ts";
import { config } from "./src/config/config.ts";

serve({ fetch: app.fetch, port: config.PORT }, (info) => {
	console.log(`Server is running on port :${info.port}`);
});
