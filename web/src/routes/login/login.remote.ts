import { PUBLIC_API_BASE_URL } from '$env/static/public';
import { type Actions, fail } from '@sveltejs/kit';

export const actions = {
	default: async ({ request, fetch }) => {
		const form = Object.fromEntries(await request.formData());
		// TODO: do a simple zod validation on the form object

		console.log({ form });
		try {
			const res = await fetch(`${PUBLIC_API_BASE_URL}/api/v1/auth/token`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(form)
			});
			const { error, data } = await res.json();
			if (!res.ok) {
				// TODO: show a toast notification for error
				console.log(error);
				return fail(400, { message: error });
			}
		} catch (_e: any) {
			console.log(_e);
			return fail(400, { message: _e.message });
		}

		// TODO: set cookies and all other stuff.
	}
} satisfies Actions;
