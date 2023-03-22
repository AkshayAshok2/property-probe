const puppeteer = require('puppeteer');
const fs = require('fs');

(async () => {
  //set scrape url
  const startDate = new Date(); 
  const endDate = new Date(startDate.getTime());
  endDate.setDate(startDate.getDate() + 30);
  fs.writeFileSync('properties.txt', '');
  
  //launch and wait for page to load
  const browser = await puppeteer.launch();
  const page = await browser.newPage();
  //set useragent
  await page.setUserAgent('Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3');
  
  for (let d = startDate; d <= endDate; d.setDate(d.getDate() + 1)) {
      //getting the date
      const year = d.getFullYear();
      const month = d.getMonth() + 1;
      const day = d.getDate();
      //Formatting the date with leading 0s
      const formattedDate =(month < 10 ? "0" + month : month) + "/" + (day < 10 ? "0" + day : day) + "/" + year;
      //Sets the formatted date for the url loop
      const url = 'https://alachua.realforeclose.com/index.cfm?zaction=AUCTION&zmethod=PREVIEW&AUCTIONDATE=' + formattedDate;
      await page.goto(url);
      
      try {
        //wait for auctions to show in html
        await page.waitForSelector('.AUCTION_ITEM.PREVIEW', { timeout: 900 });
        const html = await page.content();
        //grab auctions with cheerio
        const cheerio = require('cheerio');
        const $ = cheerio.load(html);

        const auctionItems = $('.AUCTION_ITEM.PREVIEW');

        // Loop through each auction item and print/add to txt file with all details
        for (let i = 0; i < auctionItems.length; i++) {
          console.log(formattedDate);
          fs.appendFileSync('./properties.txt', formattedDate + '\n');
          const auctionItem = auctionItems.eq(i);
          //Gets each data row from the table
          const tableRows = auctionItem.find('table tr');
          tableRows.each((index, row) => {
            const label = $(row).find('th').text().trim();
            //Puts the city,state,zipcode on the same line as the address
            if (label === 'Property Address:') {
              const data = $(row).find('td').text().trim();
              const data2 = $(tableRows.eq(index + 1)).find('td').text().trim();
              console.log(label, data, data2);
              fs.appendFileSync('./properties.txt', label + ' ' + data + ' ' + data2 + '\n');
            //Inserts the each row into the txt
            }else if(!(label === '')){
              const data = $(row).find('td').text().trim();
              console.log(label, data);
              fs.appendFileSync('./properties.txt', label + ' ' + data + '\n');
          }
          });
          //space out each auction item
          console.log("\n");
          fs.appendFileSync('./properties.txt', '\n');
        }

      }catch{
        continue
      }
    }
  await browser.close();
})();
