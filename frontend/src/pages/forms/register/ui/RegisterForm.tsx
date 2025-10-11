import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Link } from "react-router";

export function RegisterForm() {
    const urlParams = new URLSearchParams(window.location.search);
    const redirectTo = urlParams.get("redirect_to");

    async function sendForm() {
        let usernameInput = (document.getElementById("usernameInput") as HTMLInputElement).value;
        let emailInput = (document.getElementById("emailInput") as HTMLInputElement).value;
        let passwordInput = (document.getElementById("passwordInput") as HTMLInputElement).value;
        const url = `http://localhost:9100/api/v1/register`;
        try {
            const response = await fetch(url, {
                method: "POST",
                body: JSON.stringify({ username: usernameInput, email: emailInput, password: passwordInput }),
            });
            if (!response.ok) {
                throw new Error(`Response status: ${response.status}`);
            }

            // FIXME: redirect to email verification
            if (redirectTo != null) {
                window.location.replace(redirectTo);
            } else {
                window.location.replace("http://localhost:9101/");
            }
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className="p-40">
            <div className="flex flex-col items-center gap-2">
                <h1 className="text-3xl">Sign up</h1>
                <div className="flex flex-col gap-2">
                    <Label htmlFor="usernameInput">Username</Label>
                    <Input id="usernameInput" placeholder="supercooldude1337" />
                </div>
                <div className="flex flex-col gap-2">
                    <Label htmlFor="emailInput">Login</Label>
                    <Input id="emailInput" placeholder="supercooldude@gmail.com" />
                </div>
                <div className="flex flex-col gap-2">
                    <Label htmlFor="passwordInput">Password</Label>
                    <Input id="passwordInput" placeholder="not 12345" />
                </div>
                <Button onClick={sendForm}>Submit</Button>
                <p>Already have an account? <Link to="/login" className="border-b-1 border-white">Sign in</Link></p>
            </div>
        </div>
    )
}
