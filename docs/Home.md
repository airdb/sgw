Welcome to the caddywaf wiki!

A Simple WAF Based on Caddy Server.

## Decision Center
This module is charge for the decision, will do action for each request. The action is in following:

- Allow: No effect in user side except add a repsone header. 
- Block: Will block the request and return http response with 5xx code. 
- Watch: Just for internal audit check, for attack analysis.
- Tag: Tag some requests which probably(higher than 80%) be cralwer.
- Challenge: some requests maybe(50%-80%)) crawler, need to twice login.
