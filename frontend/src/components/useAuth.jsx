import { useState, useEffect } from 'react'
import axios from 'axios'

export default function useAuth(props) {
    const [accessToken, setAccessToken] = useState("")

    useEffect(() => {
        axios.post(import.meta.env.VITE_ACCESS_TOKEN_URI, { code: props.code, })
            .then((rsp) =>
                setAccessToken(rsp.data.access_token),
                window.history.pushState({},"", "/")
            )
            .catch((err) => console.log(err))
    }, [props]
    )

    console.log("AccessToken", accessToken)
    return accessToken
}

