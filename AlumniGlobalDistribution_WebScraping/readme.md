## This GoLang project aims at generating a visual representation depicting the global distribution of students belonging to a well known research group of a University.

Flow of the development:
- Scraping a web page to create a collection of alumni & faculty full names affiliated to a University Research Group (Complete)
- Find their current locations. (InProgress) 
Status: Not able to find a good REST API source that provides details like current location - city by seeing full name & some affiliation with a University name. Thought linkedIn could help but with recent changes in their data privacy policies, a user needs to authenticate in order for any application to receive even public profile information.
The other reliable source of information is Google Scholar Citations but there is no API provided by Google to simply plug in and use it. Planning to create a RESTful JSON service for the same.
- Once the current location, citations, current career & affiliation information is available, should be able to create an interactive visualization may be using maps and on-hover popups.