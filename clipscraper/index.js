const ScraperClass = require('./scraper/scraper');
const ProducerClass = require('./producer/producer');
const adapter = require('./adapter')

// ENV vars
const producerArn =  process.env.PRODUCER_ARN;
const producerEndpoint = process.env.PRODUCER_ENDPOINT;

main = async (event) => {
    console.log('scraper started')
    const producer = new ProducerClass(producerArn, producerEndpoint);
    const scraper = await ScraperClass.build();
    console.log('chrome unzipped')
    let message = adapter.adaptEventToMessage(event);
    let srcPromises = [];
    message.videoSlugs.forEach((slug)=> {
        srcPromises.push(scraper.getVidSrcFromUrl(slug))
    });
    try{
        srcs = await Promise.all(srcPromises)
        console.log(`video links founds for ${message.channelName}`)
        await producer.publishMessage(srcs, message.videoDescription, message.channelName)
        console.log(`new message published for ${message.channelName}`)
    } catch(e) {
        console.log(e)
        process.exit(1)
    }
}

exports.handler = async(event, context) => {
   await main(event)
}

