const axios = require('axios');

let sendLogsOff = async (payload, logDnaKey) => {
    const axiosClient = axios.create({
        baseURL: 'https://logs.logdna.com/logs/ingest?hostname=aws',
        timeout: 10000,
        headers: {'apikey': logDnaKey}
    });
    return axiosClient.post("", payload)
}

module.exports.sendLogsOff = sendLogsOff