import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Link } from "react-router";

export function LoginForm() {
    async function sendForm() {
        let loginInput = (document.getElementById("loginInput") as HTMLInputElement).value;
        let passwordInput = (document.getElementById("passwordInput") as HTMLInputElement).value;
        const url = `http://localhost:9100/api/v1/login${window.location.search}`;
        try {
            let response: any = {
                message: "",
                ok: true,
            };
            await fetch(url, {
                method: "POST",
                body: JSON.stringify({ login: loginInput, password: passwordInput }),
                credentials: "include",
            })
                .then(resp => {
                    response.ok = resp.ok;
                    resp.json()
                })
                .then(data => response.message = data);

            if (!response.ok) {
                throw new Error(response.message);
            }

            console.log(response);
            const newURL = `http://localhost:9100/oauth/authorize${window.location.search}`;
            window.location.replace(newURL);
        } catch (error: unknown) {
            const err = error as Error;
            alert(err.message);
            console.log(error);
        }
    }

    return (
        <div className="p-40">
            <div className="flex flex-col items-center gap-2">
                <h1 className="text-3xl">Authorization</h1>
                <div className="flex flex-col gap-2">
                    <Label htmlFor="loginInput">Login</Label>
                    <Input id="loginInput" placeholder="Email or username" />
                </div>
                <div className="flex flex-col gap-2">
                    <Label htmlFor="passwordInput">Password</Label>
                    <Input id="passwordInput" placeholder="Password" />
                </div>
                <Button onClick={sendForm}>Sign in</Button>
                <p>Don't have an account? <Link to="/register" className="border-b-1 border-white">Sign up</Link></p>
            </div>
        </div>
    )
}
