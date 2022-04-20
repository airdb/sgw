A Simple WAF Based on Caddy Server.


## Documentation

Documentation is hosted at https://github.com/airdb/caddywaf/wiki.
# Decision Center
This module is charge for the decision, will do action for each request.
The action is in "allow", "block", and "watch".

Allow: No effect in user side except add a repsone header.
Block: Will block the request and return http response with 5xx code.
Watch: Just for internal audit check,  for attack analysis.


# Strategy

Block By IP
Block By URI
Block By Cookie
Block By UserAgent


# ipip.net
https://github.com/ipipdotnet/ipdb-go
