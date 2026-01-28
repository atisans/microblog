<script>
import { register } from "./register.remote";

$inspect(register.result?.error);
</script>

<main
	class="flex min-h-screen items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 px-4 py-12 sm:px-6 lg:px-8"
>
	<section class="w-full max-w-md space-y-8">
		<header class="text-center">
			<h1 class="mb-2 text-3xl font-bold text-gray-900">Create your account</h1>
			<p class="text-sm text-gray-600">Join our microblog community</p>
		</header>

		<form
			{...register}
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
						class="input w-full text-gray-900 placeholder-gray-500 transition-colors duration-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Enter your username"
					/>
				</div>

				<div>
					<label for="email" class="label mb-2 text-gray-700"> Email address </label>
					<input
						id="email"
						name="email"
						type="email"
						required
						class="input w-full text-gray-900 placeholder-gray-500 transition-colors duration-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Enter your email"
					/>
				</div>

				<div>
					<label for="password" class="label mb-2 text-gray-700"> Password </label>
					<input
						id="password"
						name="password"
						type="password"
						required
						class="input w-full text-gray-900 placeholder-gray-500 transition-colors duration-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Create a password (min 6 characters)"
					/>
				</div>

				<div>
					<label for="password2" class="label mb-2 text-gray-700"> Confirm Password </label>
					<input
						id="password2"
						name="password2"
						type="password"
						required
						class="input w-full text-gray-900 placeholder-gray-500 transition-colors duration-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-500 focus:outline-none"
						placeholder="Confirm your password"
					/>
				</div>

				{#if register.result?.error}
					<div class="rounded-lg border border-red-200 bg-red-50 px-4 py-3 text-red-700">
						<p class="text-sm font-medium">{register.result.error}</p>
					</div>
				{/if}

				<button
					type="submit"
					disabled={register.pending > 0 || false}
					class="group relative flex w-full justify-center rounded-lg border border-transparent bg-indigo-600 px-4 py-3 text-sm font-medium text-white transition-colors duration-200 hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					{#if register.pending}
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
					Create Account
				</button>
			</fieldset>
		</form>

		<footer class="text-center">
			<p class="text-sm text-gray-600">
				Already have an account?
				<a
					href="/login"
					class="font-medium text-indigo-600 transition-colors duration-200 hover:text-indigo-500"
				>
					Sign in here
				</a>
			</p>
		</footer>
	</section>
</main>
