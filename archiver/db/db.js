const AWS = require('aws-sdk');

class db {
    constructor(customEndpoint){
        let dyanmodbConfig = {apiVersion: '2012-11-05'};
        if (customEndpoint) {
            dyanmodbConfig.region = "us-east-1"
            dyanmodbConfig.endpoint = new AWS.Endpoint(customEndpoint);
        }
        this._dyanmodbClient = new AWS.DynamoDB(dyanmodbConfig);
        this._retryCount = 60
    }

    async saveMessage(dbItem){
        let params = {
            TableName: dbItem.tableName
        };
        let item = {
                QueName :{
                    S: dbItem.queName
                },
                MessageID:{
                    S: dbItem.messageId
                },
                TimeStamp:{
                    S: dbItem.timestamp
                },
                Message:{
                    S: dbItem.message
                }
        };
        params.Item = item
        console.log(params)
        return this._dyanmodbClient.putItem(params).promise();
    }
}

module.exports = db