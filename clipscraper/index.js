const ScraperClass = require('./scraper/scraper');
const ProducerClass = require('./producer/producer');
const ConsumerClass = require('./consumer/consumer');

// ENV vars
const consumerUrl = process.env.CONSUMER_URL;
const consumerEndpoint = process.env.CONSUMER_ENDPOINT;
const producerArn =  process.env.PRODUCER_ARN;
const producerEndpoint = process.env.PRODUCER_ENDPOINT;

main = async () => {
    console.log('scraper started')
    const producer = new ProducerClass(producerArn, producerEndpoint)
    const consumer = new ConsumerClass(consumerUrl, consumerEndpoint) 
    const scraper = await ScraperClass.build()
    console.log('chrome unzipped')
    let message = {}
    try {
        message = await consumer.getMessage()
    } catch(e) {
        console.log(e)
        process.exit(1)
    }
    console.log(`message recevied for ${message.channelName}`)
    let srcPromises = [];
    message.videoSlugs.forEach((slug)=> {
        srcPromises.push(scraper.getVidSrcFromUrl(slug))
    });
    try{
        srcs = await Promise.all(srcPromises)
        console.log(`video links founds for ${message.channelName}`)
        await producer.publishMessage(srcs, message.videoDescription, message.channelName)
        console.log(`new message published for ${message.channelName}`)
        await consumer.deleteMessage(message)
        console.log(`current message deleted for ${message.channelName}`)
    } catch(e) {
        console.log(e)
        process.exit(1)
    }
}

exports.handler = async(event, context) => {
   await main()
   
}

