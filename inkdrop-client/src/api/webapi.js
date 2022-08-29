import axios from "axios";

const root = "http://127.0.0.1:9000/api/";

const webapi = {};

webapi.get = (url) => {
  return axios.get(root + url);
};

export default webapi;
