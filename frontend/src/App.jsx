import { useState } from 'react'
import Login from './components/Login'
import Dashboard from './components/Dashboard'

function App() {
    const accessCode = new URLSearchParams(window.location.search).get("code")

    return accessCode ? <Dashboard code={accessCode} /> : <Login />
}

export default App
