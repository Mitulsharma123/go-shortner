# go-shortner
Build a simple URL shortener service that will accept a URL as an argument over a REST API and
return a shortened URL as a result.
a. If you have not used or seen a URL shortener before - please visit h ps://bitly.com/ and
try to SHORTEN a URL. The goal in this assignment is not to build a fancy UI but an API
only version of that.
2. The URL and shortened URL should be stored in memory by applica on.
a. [BONUS] Instead of in memory, store these things in a text ﬁle.
3. If I again ask for the same URL, it should give me the same URL as it gave before instead of
genera ng a new one.
4. [BONUS] Put this applica on in a Docker image by wri ng a Dockerﬁle and provide the docker
image link along with the source code link

Steps to start/run the project : 
1. run main.go and curl from different termianl $ curl http://localhost:9808
2. Server should start at localhost and print message
3. Request url shortening action b post $ curl http://localhost:9808/create-short-url
curl --request POST \
--data '{
    "long_url": "http://www.geekforgeeks.com",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}' \
  http://localhost:9808/create-short-url
4. Post with long_url and user_id, it should return json response as below 
5. tesing the generated url, once you click on url it should redirect you to original url
