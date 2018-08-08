const messageAdapter = require("../adapter/messageAdapter.js");
const assert = require("chai").assert;

describe('messageAdapter', () => {

    describe('#adaptQueToDb()', () =>{
        it("should transform the message properly", ()=>{
            let mockDbTable = "fakeTable"
            let mockQueUrl = "http://localstack:4576/queue/mock-dead-letter"
            let mockMessage = {
                "Type" : "a",
                "MessageId" : "b",
                "TopicArn" : "c",
                "Message" : "d",
                "Timestamp" : "e",
                "SignatureVersion" : "f",
                "Signature" : "g",
                "SigningCertURL" : "h",
                "UnsubscribeURL" : "j"
            }
            let result = messageAdapter.adaptQueToDb(mockQueUrl, mockDbTable, mockMessage)
            assert.equal(mockMessage.MessageId, result.messageId)
            assert.equal(mockMessage.Timestamp, result.timestamp)
            assert.equal(mockMessage.Message, result.message)
            assert.equal(mockDbTable, result.tableName)
            assert.equal("mock-dead-letter", result.queName)
        })
    })
})