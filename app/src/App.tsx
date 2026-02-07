import { useEffect } from "preact/hooks";
import { signal } from "@preact/signals";
import { Link, Route, Switch } from "wouter";
import { GreetUser, SaveName } from "./wailsjs/go/backend/App";

const name = signal("");
const view = signal<"loading" | "onboarding" | "greeting">("loading");
const errorMsg = signal("");

function Onboarding() {
    const handleSubmit = async (e: Event) => {
        e.preventDefault();
        errorMsg.value = "";
        try {
            await SaveName(name.value);
            view.value = "greeting";
        } catch (err: any) {
            errorMsg.value = err.toString();
        }
    };

    return (
        <article>
            <header>Welcome</header>
            <p>Please enter your name to get started.</p>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    value={name.value}
                    onInput={(e) => (name.value = (e.target as HTMLInputElement).value)}
                    placeholder="Your Name"
                    required
                />
                <small style={{ color: "red" }}>{errorMsg.value}</small>
                <footer>
                    <button type="submit">Continue</button>
                </footer>
            </form>
        </article>
    );
}

function Greeting() {
    return (
        <article>
            <header>Hello, {name}</header>
            <p>Welcome back!</p>
            <footer>
                <button
                    className="primary"
                    onClick={() => {
                        name.value = "";
                        view.value = "onboarding";
                    }}
                >
                    Change Name
                </button>
            </footer>
        </article>
    );
}

function Home() {
    useEffect(() => {
        GreetUser().then((n) => {
            if (n) {
                name.value = n;
                view.value = "greeting";
            } else {
                view.value = "onboarding";
            }
        });
    }, []);

    return (
        <>
            {view.value === "loading" && <p aria-busy="true">Loading...</p>}
            {view.value === "onboarding" && <Onboarding />}
            {view.value === "greeting" && <Greeting />}
        </>
    );
}

function About() {
    return (
        <article>
            <header>About Codify Lite</header>
            <p>This starter kit is built with the following 15 technologies:</p>

            <details open>
                <summary>Environment & Tools</summary>
                <ul>
                    <li><strong>Mise</strong>: Environment Management</li>
                    <li><strong>Task</strong>: Task Orchestrator</li>
                    <li><strong>Bun</strong>: Package Manager</li>
                    <li><strong>Wails</strong>: Application Packaging</li>
                </ul>
            </details>

            <details open>
                <summary>Backend (Go)</summary>
                <ul>
                    <li><strong>PocketBase</strong>: Backend Framework & DB</li>
                    <li><strong>Echo</strong>: HTTP Router</li>
                    <li><strong>SQLc</strong>: Type-safe Database Access</li>
                    <li><strong>Ozzo</strong>: Data Validation</li>
                    <li><strong>Slog</strong>: Structured Logging</li>
                    <li><strong>Goblin</strong>: BDD Testing Framework</li>
                </ul>
            </details>

            <details open>
                <summary>Frontend (TypeScript)</summary>
                <ul>
                    <li><strong>Vite</strong>: Frontend Bundler</li>
                    <li><strong>Preact</strong>: UI Library</li>
                    <li><strong>Signals</strong>: State Management</li>
                    <li><strong>Wouter</strong>: Client-side Routing</li>
                    <li><strong>PicoCSS</strong>: Semantic Styling</li>
                </ul>
            </details>
        </article>
    );
}

export function App() {
    return (
        <main className="container">
            <nav>
                <ul>
                    <li><strong>Codify Lite</strong></li>
                </ul>
                <ul>
                    <li><Link href="/">Home</Link></li>
                    <li><Link href="/about">About</Link></li>
                </ul>
            </nav>

            <Switch>
                <Route path="/" component={Home} />
                <Route path="/about" component={About} />
                <Route>404: No such page!</Route>
            </Switch>
        </main>
    );
}
