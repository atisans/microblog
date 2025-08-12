import { form } from '$app/server';
import { PUBLIC_API_BASE_URL } from '$env/static/public';

export const login = form(async (formData) => {
	const form = Object.fromEntries(formData);

	try {
		const res = await fetch(`${PUBLIC_API_BASE_URL}/api/v1/auth/token`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(form)
		});
		const { error } = await res.json();
		if (!res.ok) return { error };
	} catch (_e: unknown) {
		console.log(_e);
		// @ts-ignore
		return { error: _e.message };
	}
});
