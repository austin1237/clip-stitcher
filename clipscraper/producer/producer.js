const AWS = require('aws-sdk');

class producer {

    constructor(topicARN,  customEndpoint){
        let snsConfig = {apiVersion: '2010-03-31'};
        if (customEndpoint) {
            snsConfig.endpoint = new AWS.Endpoint(customEndpoint);
        }
        this._topicARN = topicARN
        this._snsClient = new AWS.SNS(snsConfig);
    }

    publishMessage(videoLinks, videoDescription, channelName){
        const message = {
            videoLinks: videoLinks,
            videoDescription: videoDescription,
            channelName: channelName
        }
        const params = {
            TopicArn: this._topicARN,
            Message: JSON.stringify(message)
        }
        return this._snsClient.publish(params).promise();
    }

}

module.exports = producer;