import axios from 'axios'
import { useEffect, useState } from 'react'
import Listify from './Listify'
import useAuth from './useAuth'
const Dashboard = (props) => {
    let accessToken = ""
    const [user, setUser] = useState([])
    const [collections, setCollections] = useState([])
    props.cookie ? accessToken = props.cookie : accessToken = useAuth(props.code)

    const getUser = async () => {
        await axios.get("https://api.box.com/2.0/users/me", { 'headers': { 'Authorization': `Bearer ` + accessToken, } })
            .then((response) => {
                setUser(response.data)
            })
            .catch((error) => {
                console.log(error);
            })
    }

    useEffect(() => {
        getUser();
    }, [])

    const getCollections = async () => {
        await axios.get("https://api.box.com/2.0/folders/38604666009?limit=10", { 'headers': { 'Authorization': `Bearer ` + accessToken, } })
            .then((response) => {
                setCollections(response.data)
            })
            .catch((error) => {
                console.log(error);
            })
    }


    return (
        <div>
            <button type="" onClick={getCollections}>CLICK</button>
            <h1>{user.name}</h1>
            {
                collections.item_collection &&
                <Listify accessToken={accessToken} folders={collections.item_collection} />
            }

        </ div>
    )
}

export default Dashboard
