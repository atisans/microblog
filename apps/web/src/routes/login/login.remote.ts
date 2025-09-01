import { form } from "$app/server";
import { API_ENDPOINT } from "$env/static/private";

export const login = form(async (formData) => {
	const form = Object.fromEntries(formData);

	try {
		const res = await fetch(`${API_ENDPOINT}/api/v1/auth/token`, {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(form),
		});
		const { error } = await res.json();
		if (!res.ok) return { error };
	} catch (_e: unknown) {
		console.log(_e);
		// @ts-ignore
		return { error: _e.message };
	}
});
