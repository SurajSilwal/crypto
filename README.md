# crypto


The directory contains a Go application that fetches the current prices of Bitcoin, Ethereum, and Tether using an external API.

The application will include an API endpoint named "price".

The following API is used to fetch the information from the web (Coingecko.com)

"https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad"



We designed the program in such a way that the price of cryptocurrencies will be fetched with 3 different links as below.

For fetching Bitcoin price:
http://localhost:9090/price/bitcoin/

For fetching Ethereum price:
http://localhost:9090/price/ethereum/

For fetching Tether price:
http://localhost:9090/price/tether/

-----------------------------------------------

Docker link : harshdeep14/crypto
Github link: https://github.com/SurajSilwal/crypto.git/

