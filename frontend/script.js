const axios = require("axios");

axios.get("http://localhost:8080").then((res) => console.log(res.data));
