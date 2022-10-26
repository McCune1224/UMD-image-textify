import axios from 'axios'
import { useState } from 'react'

const Listify = (props) => {

    const [image, setImage] = useState([])
    let folders = props["folders"]["entries"]
    /**/
    /* const getCollections = async () => { */
    /*         .then((response) => { */
    /*     setCollections(response.data) */
    /* }) */
    /* .catch((error) => { */
    /*     console.log(error); */
    /* }) */
    /* } */
    const listItems = folders.map((number) => {
        axios.get(`https://api.box.com/2.0/files/` + number.id + `/thumbnail.png`, { 'headers': { 'Authorization': `Bearer ` + props["accessToken"], } })
            .then((rsp) => console.log(rsp.data))
        return < li key={number.id} > {number.name}</li >
    });

    return (
        <ul>{listItems}</ul>
    )
}


export default Listify
