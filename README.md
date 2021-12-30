# Overview

There is no time limit to completing this exercise, however, we recommend you aim to spend in the region of 2 hours on it. Once you have finished - and are happy with your solution - give us a shout.

# Task

- Fork the repository in your private Github account
- Write a Golang script, which scrapes a value from a HTML website, multiplies it with some other value and returns the result
- Push and send us the link (optional raise a PR)

## Description

In this task, we want to calculate the `annualized earnings` for the Ardor project.
In order to do so, we need the `fees of the last 24h` and the `current price` of the Ardor token.

The fees should be scraped from this website: https://ardorportal.org/monitor (Fees (24h))

The price can be fetched from the Coingecko API: https://api.coingecko.com/api/v3/coins/ardor (market_data.current_price.usd)

The formula to calculate the annualized earnings is: `Fees of the last 24h * 365 * price in USD`

## Bonus

If the task is too simple or your're bored, you can do the following bonuses:

- Write tests
- Do actionable error handling
- Save the output to a database
- Parametrize the input URL and field
- Parametrize the calculation formula
