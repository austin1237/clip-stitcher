const logAdapter = require("../adapters/logAdapter.js");
const assert = require("chai").assert;

describe('logAdapter', () => {
    describe('#transformCWToLD()', () =>{
        let checkLine = (line, expectedApp, expectedLevel, expectedText) => {
            assert.equal(line.app, expectedApp)
            assert.equal(line.level, expectedLevel)
            assert.equal(line.line, expectedText)
        }
        
        it("should transform the cloudwatch log properly", ()=>{
            let testLog = {
                messageType: "DATA_MESSAGE",
                owner: "359864072703",
                logGroup: "/aws/lambda/local-test",
                logStream: "2018/08/27/[$LATEST]714a65d043464121ad3c49e3c16123e7",
                subscriptionFilters: [
                    "log-aggregation-/aws/lambda/local-test"
                ],
                logEvents: [
                    {
                        id: "34240714713919811549076669560586216595346010829856178176",
                        timestamp: 1535406750272,
                        message: "example error statment"
                    },
                    {
                        id: "34240714714811841357017894486247645326251945290095394817",
                        timestamp: 1535406750312,
                        message: "example warn statement retrying"
                    },
                    {
                        id: "34240714714811841357017894486247645326251945290095394817",
                        timestamp: 1535406750312,
                        message: "example info statement"
                    }
                ]
            }
            result = logAdapter.transformCWToLD(testLog)
            assert.isArray(result.lines)
            assert.equal(result.lines.length, 3)
            checkLine(result.lines[0], "local-test", "ERROR", "example error statment")
            checkLine(result.lines[1], "local-test", "WARN", "example warn statement retrying")
            checkLine(result.lines[2], "local-test", "INFO", "example info statement")
        });
    });

    describe('#detectLevel()', () =>{
        it("if no keywords are found it should return INFO", ()=>{
            let testText = "errantia in a statment should return info"
            result = logAdapter.dectectLevel(testText)
            assert.equal(result, "INFO")
        });

        it("retry should return WARN", ()=>{
            let testText = "will retry return correctly"
            result = logAdapter.dectectLevel(testText)
            assert.equal(result, "WARN")
        });

        it("retrying should return WARN", ()=>{
            let testText = "will retrying return correctly"
            result = logAdapter.dectectLevel(testText)
            assert.equal(result, "WARN")
        });

        it("err should return ERROR", ()=>{
            let testText = "will err return correctly"
            result = logAdapter.dectectLevel(testText)
            assert.equal(result, "ERROR")
        });

        it("error should return ERROR", ()=>{
            let testText = "will error return correctly"
            result = logAdapter.dectectLevel(testText)
            assert.equal(result, "ERROR")
        });

    });
});