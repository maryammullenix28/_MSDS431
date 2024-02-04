# Wikipedia Scraper

A simple Go program to scrape information from specified Wikipedia pages using the Colly web scraping framework.

#### Wikipedia Pages Visited
- https://en.wikipedia.org/wiki/Robotics
- https://en.wikipedia.org/wiki/Robot
- https://en.wikipedia.org/wiki/Reinforcement_learning
- https://en.wikipedia.org/wiki/Robot_Operating_System
- https://en.wikipedia.org/wiki/Intelligent_agent
- https://en.wikipedia.org/wiki/Software_agent
- https://en.wikipedia.org/wiki/Robotic_process_automation
- https://en.wikipedia.org/wiki/Chatbot
- https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence
- https://en.wikipedia.org/wiki/Android_(robot)

## Implementation

This program utilizes the [Colly](https://github.com/gocolly/colly) library to scrape information from a set of Wikipedia pages and save the extracted data in JSON format. The final output is saved to 'items.JSON' in the current directory.

## Memory & Processing Benchmarks
Benchmarks obtained from the terminal's 'time' function indicates that Go has significantly higher overall CPU usage (350% vs 76%) suggesting that it's better at utilizing multiple CPU cores. 

This can be seen with the difference in real and user time whereby despite the higher 'user' time than 'real' time. This is possible due to the program's efficincy and parallel execution. The Python program on the other hand utilizes only one core (<100%) where the 'real' program runtime is longer, despite a shorter 'user' time.

Go
- real: 4.716 s
- user: 13.86 s
- sys: 2.68 s
- cpu: 350%

Python
- real: 11.410 s
- user: 7.51 s
- sys: 1.26 s
- cpu: 76%

## Testing
The test written checks that the scraper is correctly limited to the defined allowed domain (en.wikipedia.org) by taking an input of a webpage from de.wikipedia.org.



