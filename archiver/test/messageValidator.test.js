const messageValidator = require("../validator/messageValidator.js");
const expect = require("chai").expect;

describe('messageValidator', () => {

    describe('#validateMessage()', () =>{
        it("should not throw and error with a valid message", ()=>{
            let mockMessage = {
                "MessageId" : "b",
                "Message" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).not.to.throw();
        })

        it("should throw an error with missing MessageId", ()=>{
            let mockMessage = {
                "Message" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid messageId found in message");
        })

        it("should throw an error with an empty MessageId", ()=>{
            let mockMessage = {
                "MessageId" : "",
                "Message" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid messageId found in message");
        })

        it("should throw an error with missing Message", ()=>{
            let mockMessage = {
                "MessageId" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid messsage body found in message");
        })

        it("should throw an error with empty Message", ()=>{
            let mockMessage = {
                "Message" : "",
                "MessageId" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid messsage body found in message");
        })

        it("should throw an error with empty Message", ()=>{
            let mockMessage = {
                "Message" : "",
                "MessageId" : "d",
                "Timestamp" : "e",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid messsage body found in message");
        })

        it("should throw an error with missing Timestamp", ()=>{
            let mockMessage = {
                "Message" : "a",
                "MessageId" : "d",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid timestamp found in message");
        })

        it("should throw an error with empty Timestamp", ()=>{
            let mockMessage = {
                "Message" : "a",
                "MessageId" : "d",
                "Timestamp" : "",
            }
            expect(()=>{
                messageValidator.validateMessage(mockMessage)
            }).to.throw("invalid timestamp found in message");
        })
    })
})