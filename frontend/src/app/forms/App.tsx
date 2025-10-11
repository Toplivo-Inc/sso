import './App.css'

function App() {
    function signIn() {
        const urlParams = new URLSearchParams(window.location.search);
        const authRequest = urlParams.get("auth_request");

        let loginBtn = document.getElementById("login_btn") as HTMLElement;
        loginBtn.addEventListener("click", sendForm);
        async function sendForm() {
            let loginInput = (document.getElementById("loginInput") as HTMLInputElement).value;
            let passwordInput = (document.getElementById("passwordInput") as HTMLInputElement).value;
            const url = `http://localhost:9100/api/v1/login?auth_request=${authRequest}`;
            try {
                const response = await fetch(url, {
                    method: "POST",
                    body: JSON.stringify({ login: loginInput, password: passwordInput }),
                });
                if (!response.ok) {
                    throw new Error(`Response status: ${response.status}`);
                }

                const newURL = `http://localhost:9100/oauth/authorize${window.location.search}&auth_request=${authRequest}`
                window.location.replace(newURL);
            } catch (error: any) {
                console.error(error.message);
            }
        }
    }

    return (
        <>
            <h1>
                Sign in
            </h1>
            <div className="login">
                <div className="login-input">
                    Login
                    <input type="text" id="loginInput" />
                </div>
                <div className="login-input">
                    Password
                    <input type="text" id="passwordInput" />
                </div>
                <button onClick={signIn} id="login_btn" >Sign in</button>
            </div>
        </>
    )
}

export default App
