import Axios from "axios";

export const axios = Axios.create({
    baseURL: import.meta.env.BASE_URL + "/api",
    headers: {
        "Content-type": "application/json",
    },
});
