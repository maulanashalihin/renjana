/**
 * Global client-side error handlers.
 *
 * Catches errors that fall outside Svelte's render tree (inertia bootstrap,
 * event handlers, async/fetch, resource load failures, etc).
 * Svelte render errors ($derived, template) are handled by <svelte:boundary>
 * in AppLayout.svelte — this is the last line of defense for everything else.
 */

const API_ENDPOINT = "/api/errors";

interface ErrorPayload {
	message: string;
	stack: string | null;
	url: string;
	userAgent: string;
	type: "runtime" | "rejection";
}

function sendToServer(payload: ErrorPayload) {
	fetch(API_ENDPOINT, {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify(payload),
	}).catch((e) => {
		// Don't let error-reporting itself become an error
		console.warn("[ErrorReporter] failed to send", e);
	});
}

/**
 * Install global error listeners. Call BEFORE createInertiaApp()
 * so no bootstrap error goes uncaught.
 */
export function installGlobalErrorHandler(): void {
	// --- Runtime errors ---
	window.addEventListener("error", (event: ErrorEvent) => {
		const { message, error, filename, lineno, colno } = event;

		console.error("[Global Error]", {
			message,
			filename,
			lineno,
			colno,
			error,
		});

		sendToServer({
			message: message ?? String(error ?? "Unknown error"),
			stack: error?.stack ?? null,
			url: location.href,
			userAgent: navigator.userAgent,
			type: "runtime",
		});
	});

	// --- Unhandled promise rejections ---
	window.addEventListener(
		"unhandledrejection",
		(event: PromiseRejectionEvent) => {
			const reason = event.reason;

			console.error("[Unhandled Rejection]", reason);

			sendToServer({
				message: reason?.message ?? String(reason ?? "Unknown rejection"),
				stack: reason?.stack ?? null,
				url: location.href,
				userAgent: navigator.userAgent,
				type: "rejection",
			});
		},
	);
}
