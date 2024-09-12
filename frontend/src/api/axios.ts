import Axios from "axios";

export const axios = Axios.create({
    baseURL: import.meta.env.BASE_URL + "/api",
    headers: {
        Authorization: "Bearer " + localStorage.getItem("access_token"),
    },
});
