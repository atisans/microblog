<script>
// Simple homepage without over-abstraction
</script>

<main class="min-h-screen bg-gradient-to-br from-slate-50 to-blue-100">
	<div class="container mx-auto px-4 py-16">
		<div class="text-center">
			<h1 class="mb-4 text-6xl font-bold text-gray-900">Microblog</h1>
			<p class="mx-auto mb-12 max-w-2xl text-xl text-gray-600">
				Share your thoughts, connect with others, and discover amazing content in our vibrant
				microblogging community.
			</p>

			<div class="flex flex-col items-center justify-center gap-4 sm:flex-row">
				<a
					href="/register"
					class="inline-flex items-center rounded-lg border border-transparent bg-indigo-600 px-8 py-3 text-base font-medium text-white transition-colors duration-200 hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none"
				>
					Get Started
				</a>

				<a
					href="/login"
					class="inline-flex items-center rounded-lg border border-gray-300 bg-white px-8 py-3 text-base font-medium text-gray-700 transition-colors duration-200 hover:bg-gray-50 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none"
				>
					Sign In
				</a>
			</div>
		</div>

		<section class="mt-20">
			<div class="mx-auto grid max-w-4xl grid-cols-1 gap-8 md:grid-cols-3">
				<div class="rounded-xl border border-gray-200 bg-white p-6 text-center shadow-sm">
					<div
						class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-indigo-100"
					>
						<svg
							class="h-6 w-6 text-indigo-600"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
							></path>
						</svg>
					</div>
					<h3 class="mb-2 text-lg font-semibold text-gray-900">Share Your Story</h3>
					<p class="text-gray-600">
						Express yourself with short, impactful posts that resonate with your audience.
					</p>
				</div>

				<div class="rounded-xl border border-gray-200 bg-white p-6 text-center shadow-sm">
					<div
						class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-indigo-100"
					>
						<svg
							class="h-6 w-6 text-indigo-600"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z"
							></path>
						</svg>
					</div>
					<h3 class="mb-2 text-lg font-semibold text-gray-900">Connect & Engage</h3>
					<p class="text-gray-600">
						Build meaningful connections and engage in conversations that matter.
					</p>
				</div>

				<div class="rounded-xl border border-gray-200 bg-white p-6 text-center shadow-sm">
					<div
						class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-indigo-100"
					>
						<svg
							class="h-6 w-6 text-indigo-600"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"
							></path>
						</svg>
					</div>
					<h3 class="mb-2 text-lg font-semibold text-gray-900">Discover Content</h3>
					<p class="text-gray-600">
						Explore diverse perspectives and discover content that inspires and informs.
					</p>
				</div>
			</div>
		</section>
	</div>
</main>
