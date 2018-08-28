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
            message: "2018-08-27T21:52:30.253Z\t7eb4835e-aa43-11e8-a224-7ffcf0d7b41f\tvalidating event\n"
        },
        {
            id: "34240714714811841357017894486247645326251945290095394817",
            timestamp: 1535406750312,
            message: "2018-08-27T21:52:30.312Z\t7eb4835e-aa43-11e8-a224-7ffcf0d7b41f\texiting due to error\n"
        },
        {
            id: "34240714715257856260988506949078359691704912520215003138",
            timestamp: 1535406750332,
            message: "2018-08-27T21:52:30.312Z\t7eb4835e-aa43-11e8-a224-7ffcf0d7b41f\tError: invalid records found in message\n    at Object.validateEvent (/var/task/validator/eventValidator.js:4:15)\n    at main (/var/task/index.js:38:24)\n    at exports.handler (/var/task/index.js:51:11)\n"
        },
        {
            id: "34240714715726171910157650035050609775430528111840591875",
            timestamp: 1535406750353,
            message: "END RequestId: 7eb4835e-aa43-11e8-a224-7ffcf0d7b41f\n"
        },
        {
            id: "34240714715726171910157650035050609775430528111840591876",
            timestamp: 1535406750353,
            message: "REPORT RequestId: 7eb4835e-aa43-11e8-a224-7ffcf0d7b41f\tDuration: 284.23 ms\tBilled Duration: 300 ms \tMemory Size: 128 MB\tMax Memory Used: 33 MB\t\n"
        },
        {
            id: "34240714715726171910157650035050609775430528111840591877",
            timestamp: 1535406750353,
            message: "RequestId: 7eb4835e-aa43-11e8-a224-7ffcf0d7b41f Process exited before completing request\n\n"
        }
    ]
}

module.exports = testLog
