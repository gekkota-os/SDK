Relay is a solution to fetch result from many comptipix using a relay server.
Relay server load index.html in relay_web/
So you can put your file in relay_web/ and use html + javascript to
to fetch result from many comptipix : passing by a relay server bypass the CORS (Cross Origin Resource Sharing) limitation of browser !

How to :
---
1/ Start relay
2/ With internet browser connect to your machine on port 8888 (or another port if you used -port option)
3/ Play with demo
4/ Check option with "-h" flag
5/ Check usage exemple with "-help_usage" flag

Advices for developer :
---
1/ Reuse app_relay.js + app_request.js as they are and :
2/ Adapt your page in index.html + demo.js

Note to go further :
---
relay server respond to JSON request and relay your request to the target define in your JSON.
relay JSON request is documented and explained in app_relay.js (in comment)
index.html is just a demo, you can go further and fetch result from many comptipix in the same web page
