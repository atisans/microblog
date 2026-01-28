<script>
import { login } from "./login.remote";
</script>

<main
	class="flex min-h-screen items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 px-4 py-12 sm:px-6 lg:px-8"
>
	<section class="w-full max-w-md space-y-8">
		<header class="text-center">
			<h1 class="mb-2 text-3xl font-bold text-gray-900">Welcome back</h1>
			<p class="text-sm text-gray-600">Sign in to your account</p>
		</header>

		<form
			{...login}
			class="mt-8 space-y-6 rounded-xl border border-gray-200 bg-white p-8 shadow-lg"
		>
			<fieldset class="space-y-5">
				<div>
					<label for="username" class="label mb-2 text-gray-700"> Username </label>
					<input
						id="username"
						name="username"
						type="text"
						required
						class="input text-gray-900 placeholder-gray-500 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Enter your username"
					/>
				</div>

				<div>
					<label for="password" class="label mb-2 text-sm text-gray-700"> Password </label>
					<input
						id="password"
						name="password"
						type="password"
						required
						class="input text-gray-900 placeholder-gray-500 transition-colors duration-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Enter your password"
					/>
				</div>

				<div class="flex items-center">
					<input
						id="remember"
						name="remember"
						type="checkbox"
						class="checkbox rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
					/>
					<label for="remember" class="label ml-2 text-gray-700">Remember me</label>
				</div>

				{#if login.result?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-red-700">
						<p class="text-sm font-medium">{login.result.error}</p>
					</div>
				{/if}

				<button
					type="submit"
					disabled={login.pending > 0 || false}
					class="btn relative flex w-full bg-indigo-600 px-4 py-3 text-sm transition-colors duration-200 hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if login.pending}
						<svg
							class="mr-3 -ml-1 h-5 w-5 animate-spin text-white"
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
						>
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
							></circle>
							<path
								class="opacity-75"
								fill="currentColor"
								d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
							></path>
						</svg>
					{/if}
					Sign In
				</button>
			</fieldset>
		</form>

		<footer class="text-center">
			<p class="text-sm text-gray-600">
				Don't have an account?
				<a
					href="/register"
					class="font-medium text-indigo-600 transition-colors duration-200 hover:text-indigo-500"
				>
					Create one here
				</a>
			</p>
		</footer>
	</section>
</main>
