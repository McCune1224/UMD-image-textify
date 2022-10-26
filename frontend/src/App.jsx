import { useState } from 'react'
import { useCookies } from 'react-cookie'
import Login from './components/Login'
import Dashboard from './components/Dashboard'

function App() {
    const [cookies, setCookies] = useCookies(["Tokens"])

    if (cookies["access_token"]) {
        return <Dashboard cookie={cookies["access_token"]} />
    }

    const accessCode = new URLSearchParams(window.location.search).get("code")


    return accessCode ? <Dashboard code={accessCode} /> : <Login />
}

export default App
