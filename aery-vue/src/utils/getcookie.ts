import url from '../config/tsconfig.json'
import axios from "axios";
export const getSession= async ()=> {
        const result = await axios.post(url.zurl + "/getsession", {}, { withCredentials: true });
        return result.data;
}