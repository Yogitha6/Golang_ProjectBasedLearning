# Google Scholar API
To learn golang website creation by creating a website & REST API end point for Google scholar website.

Current functionality implemented:
* GET request displays A simple HTML web page with a form field and submit button. {User is expected to enter full name of whose details we want to retrieve from google scholar.}
* Upon filling the form field followed by a submit, client sends a POST request
* Backend code on receiving a post request, performs Google Scholar web page scraping based on the input form data (full name), the retrieved data is displayed back on the UI.
* Provides JSON data at a REST end point. {Authentication is not yet implemented}
* Added local mysql database functionality. The website stores scrapped information locally in the database. Everytime user runs a search, a db lookup is performed and if data is not stale the information is presented to the user.
 {Information is considered stale if it is older than 7 days} If data is stale a web scrapping operation is triggered to fetch the latest data.
* Tried deploying the website to Heroku, but interestingly google seems to block all the heroku IP ranges, when my application is run on my local machine, it works fine but on heroku - Google blocks the request saying an automatic script is detected, can't serve the page.

Yet to do:
* Authentication for REST API access
* Security & Encryption of credentials
* Writing test cases
