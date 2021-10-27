const isRelease = false;

const backend = isRelease ? "https://clientellapp.com:8000" : "http://localhost:8080";

import axios from "axios";
export let api = null;
export const initApi = token => { 
    api  = axios.create({
        baseURL: `${backend}/`,
        timeout: 5000,
        headers: { "Token": token }
    });
};

export default {
    isRelease,
    backend
};
