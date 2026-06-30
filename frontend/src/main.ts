import { createInertiaApp } from "@inertiajs/svelte";
import { installGlobalErrorHandler } from "./lib/error-handler";

installGlobalErrorHandler();

createInertiaApp();
