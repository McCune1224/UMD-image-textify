import { useState, useEffect } from 'react'
import { useCookies } from 'react-cookie'
import axios from 'axios'

const useAuth = (props) => {
    const [accessToken, setAccessToken] = useState("")
    const [refreshToken, setRefreshToken] = useState("")
    const [tokenExpiry, setTokenExpiry] = useState("")

    const [cookies, setCookies] = useCookies(["Tokens"])

    const getTokens = async () => {
        await axios.post(import.meta.env.VITE_ACCESS_TOKEN_URI, { code: props, })
            .then((rsp) => {
                setAccessToken(rsp.data.access_token)
                setTokenExpiry(rsp.data.expiry)
                setRefreshToken(rsp.data.refresh_token)
                window.history.pushState({}, "", "/")
            }
            )
            .catch((err) => console.log(err))
    }

    useEffect(() => {
        getTokens()
    }, [props]
    )


    if (accessToken) {
        const date = new Date(tokenExpiry)

        setCookies("access_token", accessToken, { expires: date })
        setCookies("refresh_token", refreshToken)
    }
    return accessToken
}

export default useAuth
